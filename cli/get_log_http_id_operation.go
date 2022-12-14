// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/sadov/rolctl/client/log"
	logops "github.com/sadov/rolctl/client/log"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationLogGetLogHTTPIDCmd returns a cmd to handle operation getLogHttpId
func makeOperationLogGetLogHTTPIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetLogHTTPID",
		Short: ``,
		RunE:  runOperationLogGetLogHTTPID,
	}

	if err := registerOperationLogGetLogHTTPIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLogGetLogHTTPID uses cmd flags to call endpoint api
func runOperationLogGetLogHTTPID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := logops.NewGetLogHTTPIDParams()
	if err, _ := retrieveOperationLogGetLogHTTPIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLogGetLogHTTPIDResult(appCli.Log.GetLogHTTPID(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLogGetLogHTTPIDParamFlags registers all flags needed to fill params
func registerOperationLogGetLogHTTPIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLogGetLogHTTPIDIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLogGetLogHTTPIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. log id`

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

func retrieveOperationLogGetLogHTTPIDIDFlag(m *logops.GetLogHTTPIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationLogGetLogHTTPIDResult parses request result and return the string content
func parseOperationLogGetLogHTTPIDResult(resp0 *log.GetLogHTTPIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*log.GetLogHTTPIDOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getLogHttpIdNotFound is not supported

		// Non schema case: warning getLogHttpIdInternalServerError is not supported

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
