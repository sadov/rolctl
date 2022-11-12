// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/sadov/rolctl/client/ethernet_switch"
	"github.com/sadov/rolctl/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationEthernetSwitchPostEthernetSwitchCmd returns a cmd to handle operation postEthernetSwitch
func makeOperationEthernetSwitchPostEthernetSwitchCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PostEthernetSwitch",
		Short: ``,
		RunE:  runOperationEthernetSwitchPostEthernetSwitch,
	}

	if err := registerOperationEthernetSwitchPostEthernetSwitchParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEthernetSwitchPostEthernetSwitch uses cmd flags to call endpoint api
func runOperationEthernetSwitchPostEthernetSwitch(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ethernet_switch.NewPostEthernetSwitchParams()
	if err, _ := retrieveOperationEthernetSwitchPostEthernetSwitchRequestFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEthernetSwitchPostEthernetSwitchResult(appCli.EthernetSwitch.PostEthernetSwitch(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEthernetSwitchPostEthernetSwitchParamFlags registers all flags needed to fill params
func registerOperationEthernetSwitchPostEthernetSwitchParamFlags(cmd *cobra.Command) error {
	if err := registerOperationEthernetSwitchPostEthernetSwitchRequestParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationEthernetSwitchPostEthernetSwitchRequestParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var requestFlagName string
	if cmdPrefix == "" {
		requestFlagName = "request"
	} else {
		requestFlagName = fmt.Sprintf("%v.request", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(requestFlagName, "", "Optional json string for [request]. Ethernet switch fields")

	// add flags for body
	if err := registerModelDtosEthernetSwitchCreateDtoFlags(0, "dtosEthernetSwitchCreateDto", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationEthernetSwitchPostEthernetSwitchRequestFlag(m *ethernet_switch.PostEthernetSwitchParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("request") {
		// Read request string from cmd and unmarshal
		requestValueStr, err := cmd.Flags().GetString("request")
		if err != nil {
			return err, false
		}

		requestValue := models.DtosEthernetSwitchCreateDto{}
		if err := json.Unmarshal([]byte(requestValueStr), &requestValue); err != nil {
			return fmt.Errorf("cannot unmarshal request string in models.DtosEthernetSwitchCreateDto: %v", err), false
		}
		m.Request = &requestValue
	}
	requestValueModel := m.Request
	if swag.IsZero(requestValueModel) {
		requestValueModel = &models.DtosEthernetSwitchCreateDto{}
	}
	err, added := retrieveModelDtosEthernetSwitchCreateDtoFlags(0, requestValueModel, "dtosEthernetSwitchCreateDto", cmd)
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

// parseOperationEthernetSwitchPostEthernetSwitchResult parses request result and return the string content
func parseOperationEthernetSwitchPostEthernetSwitchResult(resp0 *ethernet_switch.PostEthernetSwitchOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ethernet_switch.PostEthernetSwitchOK)
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
		resp1, ok := iResp1.(*ethernet_switch.PostEthernetSwitchBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning postEthernetSwitchInternalServerError is not supported

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
