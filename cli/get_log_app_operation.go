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

// makeOperationLogGetLogAppCmd returns a cmd to handle operation getLogApp
func makeOperationLogGetLogAppCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "GetLogApp",
		Short: ``,
		RunE:  runOperationLogGetLogApp,
	}

	if err := registerOperationLogGetLogAppParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLogGetLogApp uses cmd flags to call endpoint api
func runOperationLogGetLogApp(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := logops.NewGetLogAppParams()
	if err, _ := retrieveOperationLogGetLogAppOrderByFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogGetLogAppOrderDirectionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogGetLogAppPageFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogGetLogAppPageSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLogGetLogAppSearchFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLogGetLogAppResult(appCli.Log.GetLogApp(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLogGetLogAppParamFlags registers all flags needed to fill params
func registerOperationLogGetLogAppParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLogGetLogAppOrderByParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogGetLogAppOrderDirectionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogGetLogAppPageParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogGetLogAppPageSizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLogGetLogAppSearchParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLogGetLogAppOrderByParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	orderByDescription := `Order by field`

	var orderByFlagName string
	if cmdPrefix == "" {
		orderByFlagName = "orderBy"
	} else {
		orderByFlagName = fmt.Sprintf("%v.orderBy", cmdPrefix)
	}

	var orderByFlagDefault string

	_ = cmd.PersistentFlags().String(orderByFlagName, orderByFlagDefault, orderByDescription)

	return nil
}
func registerOperationLogGetLogAppOrderDirectionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	orderDirectionDescription := `'asc' or 'desc' for ascending or descending order`

	var orderDirectionFlagName string
	if cmdPrefix == "" {
		orderDirectionFlagName = "orderDirection"
	} else {
		orderDirectionFlagName = fmt.Sprintf("%v.orderDirection", cmdPrefix)
	}

	var orderDirectionFlagDefault string

	_ = cmd.PersistentFlags().String(orderDirectionFlagName, orderDirectionFlagDefault, orderDirectionDescription)

	return nil
}
func registerOperationLogGetLogAppPageParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageDescription := `page number`

	var pageFlagName string
	if cmdPrefix == "" {
		pageFlagName = "page"
	} else {
		pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
	}

	var pageFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pageFlagName, pageFlagDefault, pageDescription)

	return nil
}
func registerOperationLogGetLogAppPageSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pageSizeDescription := `number of entities per page`

	var pageSizeFlagName string
	if cmdPrefix == "" {
		pageSizeFlagName = "pageSize"
	} else {
		pageSizeFlagName = fmt.Sprintf("%v.pageSize", cmdPrefix)
	}

	var pageSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pageSizeFlagName, pageSizeFlagDefault, pageSizeDescription)

	return nil
}
func registerOperationLogGetLogAppSearchParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	searchDescription := `searchable value in entity`

	var searchFlagName string
	if cmdPrefix == "" {
		searchFlagName = "search"
	} else {
		searchFlagName = fmt.Sprintf("%v.search", cmdPrefix)
	}

	var searchFlagDefault string

	_ = cmd.PersistentFlags().String(searchFlagName, searchFlagDefault, searchDescription)

	return nil
}

func retrieveOperationLogGetLogAppOrderByFlag(m *logops.GetLogAppParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("orderBy") {

		var orderByFlagName string
		if cmdPrefix == "" {
			orderByFlagName = "orderBy"
		} else {
			orderByFlagName = fmt.Sprintf("%v.orderBy", cmdPrefix)
		}

		orderByFlagValue, err := cmd.Flags().GetString(orderByFlagName)
		if err != nil {
			return err, false
		}
		m.OrderBy = &orderByFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogGetLogAppOrderDirectionFlag(m *logops.GetLogAppParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("orderDirection") {

		var orderDirectionFlagName string
		if cmdPrefix == "" {
			orderDirectionFlagName = "orderDirection"
		} else {
			orderDirectionFlagName = fmt.Sprintf("%v.orderDirection", cmdPrefix)
		}

		orderDirectionFlagValue, err := cmd.Flags().GetString(orderDirectionFlagName)
		if err != nil {
			return err, false
		}
		m.OrderDirection = &orderDirectionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogGetLogAppPageFlag(m *logops.GetLogAppParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("page") {

		var pageFlagName string
		if cmdPrefix == "" {
			pageFlagName = "page"
		} else {
			pageFlagName = fmt.Sprintf("%v.page", cmdPrefix)
		}

		pageFlagValue, err := cmd.Flags().GetInt64(pageFlagName)
		if err != nil {
			return err, false
		}
		m.Page = &pageFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogGetLogAppPageSizeFlag(m *logops.GetLogAppParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("pageSize") {

		var pageSizeFlagName string
		if cmdPrefix == "" {
			pageSizeFlagName = "pageSize"
		} else {
			pageSizeFlagName = fmt.Sprintf("%v.pageSize", cmdPrefix)
		}

		pageSizeFlagValue, err := cmd.Flags().GetInt64(pageSizeFlagName)
		if err != nil {
			return err, false
		}
		m.PageSize = &pageSizeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationLogGetLogAppSearchFlag(m *logops.GetLogAppParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("search") {

		var searchFlagName string
		if cmdPrefix == "" {
			searchFlagName = "search"
		} else {
			searchFlagName = fmt.Sprintf("%v.search", cmdPrefix)
		}

		searchFlagValue, err := cmd.Flags().GetString(searchFlagName)
		if err != nil {
			return err, false
		}
		m.Search = &searchFlagValue

	}
	return nil, retAdded
}

// parseOperationLogGetLogAppResult parses request result and return the string content
func parseOperationLogGetLogAppResult(resp0 *log.GetLogAppOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*log.GetLogAppOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning getLogAppInternalServerError is not supported

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
