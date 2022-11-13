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

// makeOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDCmd returns a cmd to handle operation getEthernetSwitchIdVlanVlanUuid
func makeOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetEthernetSwitchIDVlanVlanUUID",
		Short: ``,
		RunE:  runOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUID,
	}

	if err := registerOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUID uses cmd flags to call endpoint api
func runOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ethernet_switch.NewGetEthernetSwitchIDVlanVlanUUIDParams()
	if err, _ := retrieveOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDVlanUUIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDResult(appCli.EthernetSwitch.GetEthernetSwitchIDVlanVlanUUID(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDParamFlags registers all flags needed to fill params
func registerOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDVlanUUIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Ethernet switch ID`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}
func registerOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDVlanUUIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	vlanUuidDescription := `Required. Ethernet switch VLAN UUID`

	var vlanUuidFlagName string
	if cmdPrefix == "" {
		vlanUuidFlagName = "vlanUUID"
	} else {
		vlanUuidFlagName = fmt.Sprintf("%v.vlanUUID", cmdPrefix)
	}

	var vlanUuidFlagDefault string

	_ = cmd.PersistentFlags().String(vlanUuidFlagName, vlanUuidFlagDefault, vlanUuidDescription)

	return nil
}

func retrieveOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDIDFlag(m *ethernet_switch.GetEthernetSwitchIDVlanVlanUUIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}
func retrieveOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDVlanUUIDFlag(m *ethernet_switch.GetEthernetSwitchIDVlanVlanUUIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("vlanUUID") {

		var vlanUuidFlagName string
		if cmdPrefix == "" {
			vlanUuidFlagName = "vlanUUID"
		} else {
			vlanUuidFlagName = fmt.Sprintf("%v.vlanUUID", cmdPrefix)
		}

		vlanUuidFlagValue, err := cmd.Flags().GetString(vlanUuidFlagName)
		if err != nil {
			return err, false
		}
		m.VlanUUID = vlanUuidFlagValue

	}
	return nil, retAdded
}

// parseOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDResult parses request result and return the string content
func parseOperationEthernetSwitchGetEthernetSwitchIDVlanVlanUUIDResult(resp0 *ethernet_switch.GetEthernetSwitchIDVlanVlanUUIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ethernet_switch.GetEthernetSwitchIDVlanVlanUUIDOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getEthernetSwitchIdVlanVlanUuidNotFound is not supported

		// Non schema case: warning getEthernetSwitchIdVlanVlanUuidInternalServerError is not supported

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
