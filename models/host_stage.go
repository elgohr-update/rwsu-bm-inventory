// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HostStage host stage
//
// swagger:model host-stage
type HostStage string

const (

	// HostStageStartingInstallation captures enum value "Starting installation"
	HostStageStartingInstallation HostStage = "Starting installation"

	// HostStageStartWaitingForControlPlane captures enum value "Start Waiting for control plane"
	HostStageStartWaitingForControlPlane HostStage = "Start Waiting for control plane"

	// HostStageInstalling captures enum value "Installing"
	HostStageInstalling HostStage = "Installing"

	// HostStageWritingImageToDisk captures enum value "Writing image to disk"
	HostStageWritingImageToDisk HostStage = "Writing image to disk"

	// HostStageFinishWaitingForControlPlane captures enum value "Finish Waiting for control plane"
	HostStageFinishWaitingForControlPlane HostStage = "Finish Waiting for control plane"

	// HostStageRebooting captures enum value "Rebooting"
	HostStageRebooting HostStage = "Rebooting"

	// HostStageWaitingForIgnition captures enum value "Waiting for ignition"
	HostStageWaitingForIgnition HostStage = "Waiting for ignition"

	// HostStageConfiguring captures enum value "Configuring"
	HostStageConfiguring HostStage = "Configuring"

	// HostStageJoined captures enum value "Joined"
	HostStageJoined HostStage = "Joined"

	// HostStageDone captures enum value "Done"
	HostStageDone HostStage = "Done"

	// HostStageFailed captures enum value "Failed"
	HostStageFailed HostStage = "Failed"
)

// for schema
var hostStageEnum []interface{}

func init() {
	var res []HostStage
	if err := json.Unmarshal([]byte(`["Starting installation","Start Waiting for control plane","Installing","Writing image to disk","Finish Waiting for control plane","Rebooting","Waiting for ignition","Configuring","Joined","Done","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostStageEnum = append(hostStageEnum, v)
	}
}

func (m HostStage) validateHostStageEnum(path, location string, value HostStage) error {
	if err := validate.EnumCase(path, location, value, hostStageEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this host stage
func (m HostStage) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHostStageEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
