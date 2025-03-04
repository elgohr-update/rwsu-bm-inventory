// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the installer client
type API interface {
	/*
	   CancelInstallation cancels an ongoing installation*/
	CancelInstallation(ctx context.Context, params *CancelInstallationParams) (*CancelInstallationAccepted, error)
	/*
	   DeregisterCluster deletes an open shift bare metal cluster definition*/
	DeregisterCluster(ctx context.Context, params *DeregisterClusterParams) (*DeregisterClusterNoContent, error)
	/*
	   DeregisterHost deregisters an open shift bare metal host*/
	DeregisterHost(ctx context.Context, params *DeregisterHostParams) (*DeregisterHostNoContent, error)
	/*
	   DisableHost disables a host for inclusion in the cluster*/
	DisableHost(ctx context.Context, params *DisableHostParams) (*DisableHostOK, error)
	/*
	   DownloadClusterFiles downloads files relating to the installed installing cluster*/
	DownloadClusterFiles(ctx context.Context, params *DownloadClusterFilesParams, writer io.Writer) (*DownloadClusterFilesOK, error)
	/*
	   DownloadClusterISO downloads the open shift per cluster discovery i s o*/
	DownloadClusterISO(ctx context.Context, params *DownloadClusterISOParams, writer io.Writer) (*DownloadClusterISOOK, error)
	/*
	   DownloadClusterKubeconfig downloads the kubeconfig file for this cluster*/
	DownloadClusterKubeconfig(ctx context.Context, params *DownloadClusterKubeconfigParams, writer io.Writer) (*DownloadClusterKubeconfigOK, error)
	/*
	   EnableHost enables a host for inclusion in the cluster*/
	EnableHost(ctx context.Context, params *EnableHostParams) (*EnableHostOK, error)
	/*
	   GenerateClusterISO creates a new open shift per cluster discovery i s o*/
	GenerateClusterISO(ctx context.Context, params *GenerateClusterISOParams) (*GenerateClusterISOCreated, error)
	/*
	   GetCluster retrieves the details of the open shift bare metal cluster*/
	GetCluster(ctx context.Context, params *GetClusterParams) (*GetClusterOK, error)
	/*
	   GetCredentials gets the the cluster admin credentials*/
	GetCredentials(ctx context.Context, params *GetCredentialsParams) (*GetCredentialsOK, error)
	/*
	   GetFreeAddresses retrieves the free address list for a network*/
	GetFreeAddresses(ctx context.Context, params *GetFreeAddressesParams) (*GetFreeAddressesOK, error)
	/*
	   GetHost retrieves the details of the open shift bare metal host*/
	GetHost(ctx context.Context, params *GetHostParams) (*GetHostOK, error)
	/*
	   GetNextSteps retrieves the next operations that the host agent needs to perform*/
	GetNextSteps(ctx context.Context, params *GetNextStepsParams) (*GetNextStepsOK, error)
	/*
	   InstallCluster installs the open shift bare metal cluster*/
	InstallCluster(ctx context.Context, params *InstallClusterParams) (*InstallClusterAccepted, error)
	/*
	   ListClusters retrieves the list of open shift bare metal clusters*/
	ListClusters(ctx context.Context, params *ListClustersParams) (*ListClustersOK, error)
	/*
	   ListHosts retrieves the list of open shift bare metal hosts*/
	ListHosts(ctx context.Context, params *ListHostsParams) (*ListHostsOK, error)
	/*
	   PostStepReply posts the result of the operations from the host agent*/
	PostStepReply(ctx context.Context, params *PostStepReplyParams) (*PostStepReplyNoContent, error)
	/*
	   RegisterCluster creates a new open shift bare metal cluster definition*/
	RegisterCluster(ctx context.Context, params *RegisterClusterParams) (*RegisterClusterCreated, error)
	/*
	   RegisterHost registers a new open shift bare metal host*/
	RegisterHost(ctx context.Context, params *RegisterHostParams) (*RegisterHostCreated, error)
	/*
	   ResetCluster resets a failed installation*/
	ResetCluster(ctx context.Context, params *ResetClusterParams) (*ResetClusterAccepted, error)
	/*
	   SetDebugStep sets a single shot debug step that will be sent next time the host agent will ask for a command*/
	SetDebugStep(ctx context.Context, params *SetDebugStepParams) (*SetDebugStepNoContent, error)
	/*
	   UpdateCluster updates an open shift bare metal cluster definition*/
	UpdateCluster(ctx context.Context, params *UpdateClusterParams) (*UpdateClusterCreated, error)
	/*
	   UpdateHostInstallProgress updates installation progress*/
	UpdateHostInstallProgress(ctx context.Context, params *UpdateHostInstallProgressParams) (*UpdateHostInstallProgressOK, error)
	/*
	   UploadClusterIngressCert transfers the ingress certificate for the cluster*/
	UploadClusterIngressCert(ctx context.Context, params *UploadClusterIngressCertParams) (*UploadClusterIngressCertCreated, error)
}

// New creates a new installer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for installer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CancelInstallation cancels an ongoing installation
*/
func (a *Client) CancelInstallation(ctx context.Context, params *CancelInstallationParams) (*CancelInstallationAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CancelInstallation",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/actions/cancel",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CancelInstallationReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CancelInstallationAccepted), nil

}

/*
DeregisterCluster deletes an open shift bare metal cluster definition
*/
func (a *Client) DeregisterCluster(ctx context.Context, params *DeregisterClusterParams) (*DeregisterClusterNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeregisterCluster",
		Method:             "DELETE",
		PathPattern:        "/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeregisterClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeregisterClusterNoContent), nil

}

/*
DeregisterHost deregisters an open shift bare metal host
*/
func (a *Client) DeregisterHost(ctx context.Context, params *DeregisterHostParams) (*DeregisterHostNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeregisterHost",
		Method:             "DELETE",
		PathPattern:        "/clusters/{cluster_id}/hosts/{host_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeregisterHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeregisterHostNoContent), nil

}

/*
DisableHost disables a host for inclusion in the cluster
*/
func (a *Client) DisableHost(ctx context.Context, params *DisableHostParams) (*DisableHostOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DisableHost",
		Method:             "DELETE",
		PathPattern:        "/clusters/{cluster_id}/hosts/{host_id}/actions/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DisableHostOK), nil

}

/*
DownloadClusterFiles downloads files relating to the installed installing cluster
*/
func (a *Client) DownloadClusterFiles(ctx context.Context, params *DownloadClusterFilesParams, writer io.Writer) (*DownloadClusterFilesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DownloadClusterFiles",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/downloads/files",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadClusterFilesReader{formats: a.formats, writer: writer},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadClusterFilesOK), nil

}

/*
DownloadClusterISO downloads the open shift per cluster discovery i s o
*/
func (a *Client) DownloadClusterISO(ctx context.Context, params *DownloadClusterISOParams, writer io.Writer) (*DownloadClusterISOOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DownloadClusterISO",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/downloads/image",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadClusterISOReader{formats: a.formats, writer: writer},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadClusterISOOK), nil

}

/*
DownloadClusterKubeconfig downloads the kubeconfig file for this cluster
*/
func (a *Client) DownloadClusterKubeconfig(ctx context.Context, params *DownloadClusterKubeconfigParams, writer io.Writer) (*DownloadClusterKubeconfigOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DownloadClusterKubeconfig",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/downloads/kubeconfig",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadClusterKubeconfigReader{formats: a.formats, writer: writer},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadClusterKubeconfigOK), nil

}

/*
EnableHost enables a host for inclusion in the cluster
*/
func (a *Client) EnableHost(ctx context.Context, params *EnableHostParams) (*EnableHostOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EnableHost",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/hosts/{host_id}/actions/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnableHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EnableHostOK), nil

}

/*
GenerateClusterISO creates a new open shift per cluster discovery i s o
*/
func (a *Client) GenerateClusterISO(ctx context.Context, params *GenerateClusterISOParams) (*GenerateClusterISOCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GenerateClusterISO",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/downloads/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GenerateClusterISOReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GenerateClusterISOCreated), nil

}

/*
GetCluster retrieves the details of the open shift bare metal cluster
*/
func (a *Client) GetCluster(ctx context.Context, params *GetClusterParams) (*GetClusterOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCluster",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOK), nil

}

/*
GetCredentials gets the the cluster admin credentials
*/
func (a *Client) GetCredentials(ctx context.Context, params *GetCredentialsParams) (*GetCredentialsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCredentials",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCredentialsReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCredentialsOK), nil

}

/*
GetFreeAddresses retrieves the free address list for a network
*/
func (a *Client) GetFreeAddresses(ctx context.Context, params *GetFreeAddressesParams) (*GetFreeAddressesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFreeAddresses",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/free_addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFreeAddressesReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFreeAddressesOK), nil

}

/*
GetHost retrieves the details of the open shift bare metal host
*/
func (a *Client) GetHost(ctx context.Context, params *GetHostParams) (*GetHostOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHost",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/hosts/{host_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHostOK), nil

}

/*
GetNextSteps retrieves the next operations that the host agent needs to perform
*/
func (a *Client) GetNextSteps(ctx context.Context, params *GetNextStepsParams) (*GetNextStepsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNextSteps",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/hosts/{host_id}/instructions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNextStepsReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNextStepsOK), nil

}

/*
InstallCluster installs the open shift bare metal cluster
*/
func (a *Client) InstallCluster(ctx context.Context, params *InstallClusterParams) (*InstallClusterAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InstallCluster",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/actions/install",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InstallClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InstallClusterAccepted), nil

}

/*
ListClusters retrieves the list of open shift bare metal clusters
*/
func (a *Client) ListClusters(ctx context.Context, params *ListClustersParams) (*ListClustersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListClusters",
		Method:             "GET",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClustersOK), nil

}

/*
ListHosts retrieves the list of open shift bare metal hosts
*/
func (a *Client) ListHosts(ctx context.Context, params *ListHostsParams) (*ListHostsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListHosts",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListHostsReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListHostsOK), nil

}

/*
PostStepReply posts the result of the operations from the host agent
*/
func (a *Client) PostStepReply(ctx context.Context, params *PostStepReplyParams) (*PostStepReplyNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStepReply",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/hosts/{host_id}/instructions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostStepReplyReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStepReplyNoContent), nil

}

/*
RegisterCluster creates a new open shift bare metal cluster definition
*/
func (a *Client) RegisterCluster(ctx context.Context, params *RegisterClusterParams) (*RegisterClusterCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegisterCluster",
		Method:             "POST",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterClusterCreated), nil

}

/*
RegisterHost registers a new open shift bare metal host
*/
func (a *Client) RegisterHost(ctx context.Context, params *RegisterHostParams) (*RegisterHostCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegisterHost",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/hosts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterHostReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterHostCreated), nil

}

/*
ResetCluster resets a failed installation
*/
func (a *Client) ResetCluster(ctx context.Context, params *ResetClusterParams) (*ResetClusterAccepted, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ResetCluster",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/actions/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResetClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResetClusterAccepted), nil

}

/*
SetDebugStep sets a single shot debug step that will be sent next time the host agent will ask for a command
*/
func (a *Client) SetDebugStep(ctx context.Context, params *SetDebugStepParams) (*SetDebugStepNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SetDebugStep",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/hosts/{host_id}/actions/debug",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetDebugStepReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetDebugStepNoContent), nil

}

/*
UpdateCluster updates an open shift bare metal cluster definition
*/
func (a *Client) UpdateCluster(ctx context.Context, params *UpdateClusterParams) (*UpdateClusterCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateCluster",
		Method:             "PATCH",
		PathPattern:        "/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterCreated), nil

}

/*
UpdateHostInstallProgress updates installation progress
*/
func (a *Client) UpdateHostInstallProgress(ctx context.Context, params *UpdateHostInstallProgressParams) (*UpdateHostInstallProgressOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateHostInstallProgress",
		Method:             "PUT",
		PathPattern:        "/clusters/{cluster_id}/hosts/{host_id}/progress",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateHostInstallProgressReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateHostInstallProgressOK), nil

}

/*
UploadClusterIngressCert transfers the ingress certificate for the cluster
*/
func (a *Client) UploadClusterIngressCert(ctx context.Context, params *UploadClusterIngressCertParams) (*UploadClusterIngressCertCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UploadClusterIngressCert",
		Method:             "POST",
		PathPattern:        "/clusters/{cluster_id}/uploads/ingress-cert",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadClusterIngressCertReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadClusterIngressCertCreated), nil

}
