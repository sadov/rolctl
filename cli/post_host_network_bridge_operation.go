// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/sadov/rolctl/client/host"
	"github.com/sadov/rolctl/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationHostPostHostNetworkBridgeCmd returns a cmd to handle operation postHostNetworkBridge
func makeOperationHostPostHostNetworkBridgeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PostHostNetworkBridge",
		Short: ``,
		RunE:  runOperationHostPostHostNetworkBridge,
	}

	if err := registerOperationHostPostHostNetworkBridgeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHostPostHostNetworkBridge uses cmd flags to call endpoint api
func runOperationHostPostHostNetworkBridge(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := host.NewPostHostNetworkBridgeParams()
	if err, _ := retrieveOperationHostPostHostNetworkBridgeRequestFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHostPostHostNetworkBridgeResult(appCli.Host.PostHostNetworkBridge(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHostPostHostNetworkBridgeParamFlags registers all flags needed to fill params
func registerOperationHostPostHostNetworkBridgeParamFlags(cmd *cobra.Command) error {
	if err := registerOperationHostPostHostNetworkBridgeRequestParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationHostPostHostNetworkBridgeRequestParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var requestFlagName string
	if cmdPrefix == "" {
		requestFlagName = "request"
	} else {
		requestFlagName = fmt.Sprintf("%v.request", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(requestFlagName, "", "Optional json string for [request]. Host bridge fields")

	// add flags for body
	if err := registerModelDtosHostNetworkBridgeCreateDtoFlags(0, "dtosHostNetworkBridgeCreateDto", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationHostPostHostNetworkBridgeRequestFlag(m *host.PostHostNetworkBridgeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("request") {
		// Read request string from cmd and unmarshal
		requestValueStr, err := cmd.Flags().GetString("request")
		if err != nil {
			return err, false
		}

		requestValue := models.DtosHostNetworkBridgeCreateDto{}
		if err := json.Unmarshal([]byte(requestValueStr), &requestValue); err != nil {
			return fmt.Errorf("cannot unmarshal request string in models.DtosHostNetworkBridgeCreateDto: %v", err), false
		}
		m.Request = &requestValue
	}
	requestValueModel := m.Request
	if swag.IsZero(requestValueModel) {
		requestValueModel = &models.DtosHostNetworkBridgeCreateDto{}
	}
	err, added := retrieveModelDtosHostNetworkBridgeCreateDtoFlags(0, requestValueModel, "dtosHostNetworkBridgeCreateDto", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Request = requestValueModel
	}
	if dryRun && debug {

		requestValueDebugBytes, err := json.Marshal(m.Request)
		if err != nil {
			return err, false
		}
		logDebugf("Request dry-run payload: %v", string(requestValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationHostPostHostNetworkBridgeResult parses request result and return the string content
func parseOperationHostPostHostNetworkBridgeResult(resp0 *host.PostHostNetworkBridgeOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*host.PostHostNetworkBridgeOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*host.PostHostNetworkBridgeBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning postHostNetworkBridgeInternalServerError is not supported

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
