// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new host API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for host API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteHostNetworkBridgeName(params *DeleteHostNetworkBridgeNameParams, opts ...ClientOption) (*DeleteHostNetworkBridgeNameNoContent, error)

	DeleteHostNetworkVlanName(params *DeleteHostNetworkVlanNameParams, opts ...ClientOption) (*DeleteHostNetworkVlanNameNoContent, error)

	GetHostNetworkBridge(params *GetHostNetworkBridgeParams, opts ...ClientOption) (*GetHostNetworkBridgeOK, error)

	GetHostNetworkBridgeName(params *GetHostNetworkBridgeNameParams, opts ...ClientOption) (*GetHostNetworkBridgeNameOK, error)

	GetHostNetworkPing(params *GetHostNetworkPingParams, opts ...ClientOption) (*GetHostNetworkPingNoContent, error)

	GetHostNetworkVlan(params *GetHostNetworkVlanParams, opts ...ClientOption) (*GetHostNetworkVlanOK, error)

	GetHostNetworkVlanName(params *GetHostNetworkVlanNameParams, opts ...ClientOption) (*GetHostNetworkVlanNameOK, error)

	PostHostNetworkBridge(params *PostHostNetworkBridgeParams, opts ...ClientOption) (*PostHostNetworkBridgeOK, error)

	PostHostNetworkVlan(params *PostHostNetworkVlanParams, opts ...ClientOption) (*PostHostNetworkVlanOK, error)

	PutHostNetworkBridgeName(params *PutHostNetworkBridgeNameParams, opts ...ClientOption) (*PutHostNetworkBridgeNameOK, error)

	PutHostNetworkVlanName(params *PutHostNetworkVlanNameParams, opts ...ClientOption) (*PutHostNetworkVlanNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteHostNetworkBridgeName deletes host network bridge by name
*/
func (a *Client) DeleteHostNetworkBridgeName(params *DeleteHostNetworkBridgeNameParams, opts ...ClientOption) (*DeleteHostNetworkBridgeNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHostNetworkBridgeNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteHostNetworkBridgeName",
		Method:             "DELETE",
		PathPattern:        "/host/network/bridge/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteHostNetworkBridgeNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteHostNetworkBridgeNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteHostNetworkBridgeName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteHostNetworkVlanName deletes host network vlan by name
*/
func (a *Client) DeleteHostNetworkVlanName(params *DeleteHostNetworkVlanNameParams, opts ...ClientOption) (*DeleteHostNetworkVlanNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteHostNetworkVlanNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteHostNetworkVlanName",
		Method:             "DELETE",
		PathPattern:        "/host/network/vlan/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteHostNetworkVlanNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteHostNetworkVlanNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteHostNetworkVlanName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostNetworkBridge gets list of host network bridges
*/
func (a *Client) GetHostNetworkBridge(params *GetHostNetworkBridgeParams, opts ...ClientOption) (*GetHostNetworkBridgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostNetworkBridgeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHostNetworkBridge",
		Method:             "GET",
		PathPattern:        "/host/network/bridge/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostNetworkBridgeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostNetworkBridgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostNetworkBridge: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostNetworkBridgeName gets bridge port by name
*/
func (a *Client) GetHostNetworkBridgeName(params *GetHostNetworkBridgeNameParams, opts ...ClientOption) (*GetHostNetworkBridgeNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostNetworkBridgeNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHostNetworkBridgeName",
		Method:             "GET",
		PathPattern:        "/host/network/bridge/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostNetworkBridgeNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostNetworkBridgeNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostNetworkBridgeName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostNetworkPing calls the backend to notify that the current setting does not break the connection
*/
func (a *Client) GetHostNetworkPing(params *GetHostNetworkPingParams, opts ...ClientOption) (*GetHostNetworkPingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostNetworkPingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHostNetworkPing",
		Method:             "GET",
		PathPattern:        "/host/network/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostNetworkPingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostNetworkPingNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostNetworkPing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostNetworkVlan gets list of host network v l a n s
*/
func (a *Client) GetHostNetworkVlan(params *GetHostNetworkVlanParams, opts ...ClientOption) (*GetHostNetworkVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostNetworkVlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHostNetworkVlan",
		Method:             "GET",
		PathPattern:        "/host/network/vlan/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostNetworkVlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostNetworkVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostNetworkVlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostNetworkVlanName gets vlan port by name
*/
func (a *Client) GetHostNetworkVlanName(params *GetHostNetworkVlanNameParams, opts ...ClientOption) (*GetHostNetworkVlanNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostNetworkVlanNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetHostNetworkVlanName",
		Method:             "GET",
		PathPattern:        "/host/network/vlan/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostNetworkVlanNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostNetworkVlanNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHostNetworkVlanName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostHostNetworkBridge creates new host bridge
*/
func (a *Client) PostHostNetworkBridge(params *PostHostNetworkBridgeParams, opts ...ClientOption) (*PostHostNetworkBridgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostNetworkBridgeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostHostNetworkBridge",
		Method:             "POST",
		PathPattern:        "/host/network/bridge/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHostNetworkBridgeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHostNetworkBridgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostHostNetworkBridge: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostHostNetworkVlan creates new host vlan
*/
func (a *Client) PostHostNetworkVlan(params *PostHostNetworkVlanParams, opts ...ClientOption) (*PostHostNetworkVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostNetworkVlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostHostNetworkVlan",
		Method:             "POST",
		PathPattern:        "/host/network/vlan/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHostNetworkVlanReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHostNetworkVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostHostNetworkVlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutHostNetworkBridgeName updates host network bridge
*/
func (a *Client) PutHostNetworkBridgeName(params *PutHostNetworkBridgeNameParams, opts ...ClientOption) (*PutHostNetworkBridgeNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutHostNetworkBridgeNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutHostNetworkBridgeName",
		Method:             "PUT",
		PathPattern:        "/host/network/bridge/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutHostNetworkBridgeNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutHostNetworkBridgeNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutHostNetworkBridgeName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutHostNetworkVlanName updates host network vlan
*/
func (a *Client) PutHostNetworkVlanName(params *PutHostNetworkVlanNameParams, opts ...ClientOption) (*PutHostNetworkVlanNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutHostNetworkVlanNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutHostNetworkVlanName",
		Method:             "PUT",
		PathPattern:        "/host/network/vlan/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutHostNetworkVlanNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutHostNetworkVlanNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutHostNetworkVlanName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
