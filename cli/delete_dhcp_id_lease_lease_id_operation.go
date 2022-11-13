// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sadov/rolctl/client/dhcp"

	"github.com/spf13/cobra"
)

// makeOperationDhcpDeleteDhcpIDLeaseLeaseIDCmd returns a cmd to handle operation deleteDhcpIdLeaseLeaseId
func makeOperationDhcpDeleteDhcpIDLeaseLeaseIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "DeleteDhcpIDLeaseLeaseID",
		Short: ``,
		RunE:  runOperationDhcpDeleteDhcpIDLeaseLeaseID,
	}

	if err := registerOperationDhcpDeleteDhcpIDLeaseLeaseIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationDhcpDeleteDhcpIDLeaseLeaseID uses cmd flags to call endpoint api
func runOperationDhcpDeleteDhcpIDLeaseLeaseID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := dhcp.NewDeleteDhcpIDLeaseLeaseIDParams()
	if err, _ := retrieveOperationDhcpDeleteDhcpIDLeaseLeaseIDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationDhcpDeleteDhcpIDLeaseLeaseIDLeaseIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationDhcpDeleteDhcpIDLeaseLeaseIDResult(appCli.Dhcp.DeleteDhcpIDLeaseLeaseID(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationDhcpDeleteDhcpIDLeaseLeaseIDParamFlags registers all flags needed to fill params
func registerOperationDhcpDeleteDhcpIDLeaseLeaseIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationDhcpDeleteDhcpIDLeaseLeaseIDIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationDhcpDeleteDhcpIDLeaseLeaseIDLeaseIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationDhcpDeleteDhcpIDLeaseLeaseIDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. DHCP v4 server ID`

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
func registerOperationDhcpDeleteDhcpIDLeaseLeaseIDLeaseIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	leaseIdDescription := `Required. DHCP v4 lease ID`

	var leaseIdFlagName string
	if cmdPrefix == "" {
		leaseIdFlagName = "leaseID"
	} else {
		leaseIdFlagName = fmt.Sprintf("%v.leaseID", cmdPrefix)
	}

	var leaseIdFlagDefault string

	_ = cmd.PersistentFlags().String(leaseIdFlagName, leaseIdFlagDefault, leaseIdDescription)

	return nil
}

func retrieveOperationDhcpDeleteDhcpIDLeaseLeaseIDIDFlag(m *dhcp.DeleteDhcpIDLeaseLeaseIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationDhcpDeleteDhcpIDLeaseLeaseIDLeaseIDFlag(m *dhcp.DeleteDhcpIDLeaseLeaseIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("leaseID") {

		var leaseIdFlagName string
		if cmdPrefix == "" {
			leaseIdFlagName = "leaseID"
		} else {
			leaseIdFlagName = fmt.Sprintf("%v.leaseID", cmdPrefix)
		}

		leaseIdFlagValue, err := cmd.Flags().GetString(leaseIdFlagName)
		if err != nil {
			return err, false
		}
		m.LeaseID = leaseIdFlagValue

	}
	return nil, retAdded
}

// parseOperationDhcpDeleteDhcpIDLeaseLeaseIDResult parses request result and return the string content
func parseOperationDhcpDeleteDhcpIDLeaseLeaseIDResult(resp0 *dhcp.DeleteDhcpIDLeaseLeaseIDNoContent, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteDhcpIdLeaseLeaseIdNoContent is not supported

		// Non schema case: warning deleteDhcpIdLeaseLeaseIdNotFound is not supported

		// Non schema case: warning deleteDhcpIdLeaseLeaseIdInternalServerError is not supported

		return "", respErr
	}

	// warning: non schema response deleteDhcpIdLeaseLeaseIdNoContent is not supported by go-swagger cli yet.

	return "", nil
}
