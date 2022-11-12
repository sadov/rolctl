// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/sadov/rolctl/client/host"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationHostGetHostNetworkBridgeCmd returns a cmd to handle operation getHostNetworkBridge
func makeOperationHostGetHostNetworkBridgeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetHostNetworkBridge",
		Short: ``,
		RunE:  runOperationHostGetHostNetworkBridge,
	}

	if err := registerOperationHostGetHostNetworkBridgeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHostGetHostNetworkBridge uses cmd flags to call endpoint api
func runOperationHostGetHostNetworkBridge(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := host.NewGetHostNetworkBridgeParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHostGetHostNetworkBridgeResult(appCli.Host.GetHostNetworkBridge(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHostGetHostNetworkBridgeParamFlags registers all flags needed to fill params
func registerOperationHostGetHostNetworkBridgeParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationHostGetHostNetworkBridgeResult parses request result and return the string content
func parseOperationHostGetHostNetworkBridgeResult(resp0 *host.GetHostNetworkBridgeOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*host.GetHostNetworkBridgeOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getHostNetworkBridgeInternalServerError is not supported

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
