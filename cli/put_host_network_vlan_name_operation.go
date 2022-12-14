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

// makeOperationHostPutHostNetworkVlanNameCmd returns a cmd to handle operation putHostNetworkVlanName
func makeOperationHostPutHostNetworkVlanNameCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PutHostNetworkVlanName",
		Short: ``,
		RunE:  runOperationHostPutHostNetworkVlanName,
	}

	if err := registerOperationHostPutHostNetworkVlanNameParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationHostPutHostNetworkVlanName uses cmd flags to call endpoint api
func runOperationHostPutHostNetworkVlanName(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := host.NewPutHostNetworkVlanNameParams()
	if err, _ := retrieveOperationHostPutHostNetworkVlanNameNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationHostPutHostNetworkVlanNameRequestFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationHostPutHostNetworkVlanNameResult(appCli.Host.PutHostNetworkVlanName(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationHostPutHostNetworkVlanNameParamFlags registers all flags needed to fill params
func registerOperationHostPutHostNetworkVlanNameParamFlags(cmd *cobra.Command) error {
	if err := registerOperationHostPutHostNetworkVlanNameNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationHostPutHostNetworkVlanNameRequestParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationHostPutHostNetworkVlanNameNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Vlan name`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}
func registerOperationHostPutHostNetworkVlanNameRequestParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var requestFlagName string
	if cmdPrefix == "" {
		requestFlagName = "request"
	} else {
		requestFlagName = fmt.Sprintf("%v.request", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(requestFlagName, "", "Optional json string for [request]. Host vlan fields")

	// add flags for body
	if err := registerModelDtosHostNetworkVlanUpdateDtoFlags(0, "dtosHostNetworkVlanUpdateDto", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationHostPutHostNetworkVlanNameNameFlag(m *host.PutHostNetworkVlanNameParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationHostPutHostNetworkVlanNameRequestFlag(m *host.PutHostNetworkVlanNameParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("request") {
		// Read request string from cmd and unmarshal
		requestValueStr, err := cmd.Flags().GetString("request")
		if err != nil {
			return err, false
		}

		requestValue := models.DtosHostNetworkVlanUpdateDto{}
		if err := json.Unmarshal([]byte(requestValueStr), &requestValue); err != nil {
			return fmt.Errorf("cannot unmarshal request string in models.DtosHostNetworkVlanUpdateDto: %v", err), false
		}
		m.Request = &requestValue
	}
	requestValueModel := m.Request
	if swag.IsZero(requestValueModel) {
		requestValueModel = &models.DtosHostNetworkVlanUpdateDto{}
	}
	err, added := retrieveModelDtosHostNetworkVlanUpdateDtoFlags(0, requestValueModel, "dtosHostNetworkVlanUpdateDto", cmd)
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

// parseOperationHostPutHostNetworkVlanNameResult parses request result and return the string content
func parseOperationHostPutHostNetworkVlanNameResult(resp0 *host.PutHostNetworkVlanNameOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*host.PutHostNetworkVlanNameOK)
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
		resp1, ok := iResp1.(*host.PutHostNetworkVlanNameBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning putHostNetworkVlanNameNotFound is not supported

		// Non schema case: warning putHostNetworkVlanNameInternalServerError is not supported

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
