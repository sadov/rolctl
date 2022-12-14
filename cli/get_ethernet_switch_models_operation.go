// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/sadov/rolctl/client/ethernet_switch"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationEthernetSwitchGetEthernetSwitchModelsCmd returns a cmd to handle operation getEthernetSwitchModels
func makeOperationEthernetSwitchGetEthernetSwitchModelsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetEthernetSwitchModels",
		Short: ``,
		RunE:  runOperationEthernetSwitchGetEthernetSwitchModels,
	}

	if err := registerOperationEthernetSwitchGetEthernetSwitchModelsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEthernetSwitchGetEthernetSwitchModels uses cmd flags to call endpoint api
func runOperationEthernetSwitchGetEthernetSwitchModels(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ethernet_switch.NewGetEthernetSwitchModelsParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEthernetSwitchGetEthernetSwitchModelsResult(appCli.EthernetSwitch.GetEthernetSwitchModels(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEthernetSwitchGetEthernetSwitchModelsParamFlags registers all flags needed to fill params
func registerOperationEthernetSwitchGetEthernetSwitchModelsParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationEthernetSwitchGetEthernetSwitchModelsResult parses request result and return the string content
func parseOperationEthernetSwitchGetEthernetSwitchModelsResult(resp0 *ethernet_switch.GetEthernetSwitchModelsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ethernet_switch.GetEthernetSwitchModelsOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

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
