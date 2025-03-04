package host

import (
	"encoding/json"
	"time"

	"github.com/filanov/bm-inventory/internal/validators"

	"github.com/filanov/bm-inventory/internal/connectivity"

	"github.com/filanov/bm-inventory/internal/common"
	"github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/internal/hardware"
	"github.com/filanov/bm-inventory/models"
	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	statusInfoDisconnected               = "Host keepalive timeout"
	statusInfoDisabled                   = "Host is disabled"
	statusInfoDiscovering                = "Waiting for host hardware info"
	statusInfoInstalling                 = "Installation in progress"
	statusInfoResettingPendingUserAction = "Reboot the host into the installation image to complete resetting the installation"
)

type UpdateReply struct {
	State     string
	IsChanged bool
}

type baseState struct {
	log logrus.FieldLogger
	db  *gorm.DB
}

func updateState(log logrus.FieldLogger, state, stateInfo string, h *models.Host, db *gorm.DB) (*UpdateReply, error) {
	return updateStateWithParams(log, state, stateInfo, h, db)
}

func defaultReply(h *models.Host) (*UpdateReply, error) {
	return &UpdateReply{
		State:     swag.StringValue(h.Status),
		IsChanged: false,
	}, nil
}

func updateByKeepAlive(log logrus.FieldLogger, h *models.Host, db *gorm.DB) (*UpdateReply, error) {
	if h.CheckedInAt.String() != "" && time.Since(time.Time(h.CheckedInAt)) > 3*time.Minute {
		return updateState(log, HostStatusDisconnected, statusInfoDisconnected, h, db)
	}
	return defaultReply(h)
}

func updateStateWithParams(log logrus.FieldLogger, status, statusInfo string, h *models.Host, db *gorm.DB, extra ...interface{}) (*UpdateReply, error) {
	updates := map[string]interface{}{"status": status, "status_info": statusInfo, "status_updated_at": strfmt.DateTime(time.Now())}

	if len(extra)%2 != 0 {
		return nil, errors.Errorf("invalid update extra parameters %+v", extra)
	}
	for i := 0; i < len(extra); i += 2 {
		updates[extra[i].(string)] = extra[i+1]
	}

	// Query by <host-id, cluster-id, status>
	// Status is queried as well to avoid races between different components.
	dbReply := db.Model(&models.Host{}).Where("id = ? and cluster_id = ? and status = ?",
		h.ID.String(), h.ClusterID.String(), swag.StringValue(h.Status)).
		Updates(updates)
	if dbReply.Error != nil {
		return nil, errors.Wrapf(dbReply.Error, "failed to update host %s from cluster %s state from %s to %s",
			h.ID.String(), h.ClusterID, swag.StringValue(h.Status), status)
	}
	if dbReply.RowsAffected == 0 && swag.StringValue(h.Status) != status {
		return nil, errors.Errorf("failed to update host %s from cluster %s state from %s to %s, nothing have changed",
			h.ID.String(), h.ClusterID, swag.StringValue(h.Status), status)
	}
	log.Infof("Updated host <%s> status from <%s> to <%s> with fields: %s",
		h.ID.String(), swag.StringValue(h.Status), status, updates)
	isChanged := status != swag.StringValue(h.Status)
	h.Status = &status
	h.StatusInfo = &statusInfo
	return &UpdateReply{
		State:     status,
		IsChanged: isChanged,
	}, nil
}

func updateHostStateWithParams(log logrus.FieldLogger, srcStatus, statusInfo string, h *models.Host, db *gorm.DB,
	extra ...interface{}) error {

	updates := map[string]interface{}{"status": swag.StringValue(h.Status), "status_info": statusInfo}
	if len(extra)%2 != 0 {
		return errors.Errorf("invalid update extra parameters %+v", extra)
	}
	for i := 0; i < len(extra); i += 2 {
		updates[extra[i].(string)] = extra[i+1]
	}
	dbReply := db.Model(&models.Host{}).Where("id = ? and cluster_id = ? and status = ?",
		h.ID.String(), h.ClusterID.String(), srcStatus).
		Updates(updates)
	if dbReply.Error != nil {
		return errors.Wrapf(dbReply.Error, "failed to update host %s from cluster %s state from %s to %s",
			h.ID.String(), h.ClusterID, srcStatus, swag.StringValue(h.Status))
	}
	if dbReply.RowsAffected == 0 && swag.StringValue(h.Status) != srcStatus {
		return errors.Errorf("failed to update host %s from cluster %s state from %s to %s, nothing have changed",
			h.ID.String(), h.ClusterID, srcStatus, swag.StringValue(h.Status))
	}
	log.Infof("Updated host <%s> status from <%s> to <%s> with fields: %s",
		h.ID.String(), srcStatus, swag.StringValue(h.Status), updates)
	return nil
}

func getCluster(clusterID strfmt.UUID, db *gorm.DB) (*common.Cluster, error) {
	var cluster common.Cluster
	if err := db.Preload("Hosts", "status <> ?", HostStatusDisabled).First(&cluster, "id = ?", clusterID).Error; err != nil {
		return nil, err
	}
	return &cluster, nil
}

func isSufficientRole(h *models.Host) *validators.IsSufficientReply {
	var reason string
	isSufficient := true

	if h.Role == "" {
		isSufficient = false
		reason = "No role selected"
	}

	return &validators.IsSufficientReply{
		Type:         "role",
		IsSufficient: isSufficient,
		Reason:       reason,
	}
}

func checkAndUpdateSufficientHost(log logrus.FieldLogger, h *models.Host, db *gorm.DB, hwValidator hardware.Validator, connectivityValidator connectivity.Validator) (*UpdateReply, error) {
	//checking if need to change state to disconnect
	stateReply, err := updateByKeepAlive(log, h, db)
	if err != nil || stateReply.IsChanged {
		return stateReply, err
	}
	var statusInfoDetails = make(map[string]string)
	//checking inventory isInsufficient
	cluster, err := getCluster(h.ClusterID, db)
	if err != nil {
		return nil, err
	}
	inventoryReply, err := hwValidator.IsSufficient(h, cluster)
	if err != nil {
		statusInfoDetails["hardware"] = "parsing error"
	} else {
		statusInfoDetails[inventoryReply.Type] = inventoryReply.Reason
	}

	//checking connectivity isSufficient
	connectivityReply, err := connectivityValidator.IsSufficient(h, cluster)
	if err != nil {
		statusInfoDetails["connectivity"] = "unknown error"
	} else {
		statusInfoDetails[connectivityReply.Type] = connectivityReply.Reason
	}
	//checking role isSufficient
	roleReply := isSufficientRole(h)
	statusInfoDetails[roleReply.Type] = roleReply.Reason

	var newStatus, newStatusInfo string
	if inventoryReply != nil && inventoryReply.IsSufficient && roleReply.IsSufficient && connectivityReply.IsSufficient {
		newStatus = HostStatusKnown
		newStatusInfo = ""
	} else {
		statusInfo, err := json.Marshal(statusInfoDetails)
		if err != nil {
			return nil, err
		}
		newStatus = HostStatusInsufficient
		newStatusInfo = string(statusInfo)
	}

	//update status & status info in DB only if there is a change
	if swag.StringValue(h.Status) != newStatus || swag.StringValue(h.StatusInfo) != newStatusInfo {
		log.Infof("is sufficient host: %s role reply %+v inventory reply %+v connectivity reply %+v", h.ID, roleReply, inventoryReply, connectivityReply)
		return updateState(log, newStatus, newStatusInfo, h, db)
	}
	return &UpdateReply{
		State:     swag.StringValue(h.Status),
		IsChanged: false,
	}, nil
}
