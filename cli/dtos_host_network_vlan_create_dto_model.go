// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sadov/rolctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for DtosHostNetworkVlanCreateDto

// register flags to command
func registerModelDtosHostNetworkVlanCreateDtoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDtosHostNetworkVlanCreateDtoAddresses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosHostNetworkVlanCreateDtoParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosHostNetworkVlanCreateDtoVlanID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDtosHostNetworkVlanCreateDtoAddresses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: addresses []string array type is not supported by go-swagger cli yet

	return nil
}

func registerDtosHostNetworkVlanCreateDtoParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	parentDescription := `Parent interface name`

	var parentFlagName string
	if cmdPrefix == "" {
		parentFlagName = "parent"
	} else {
		parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
	}

	var parentFlagDefault string

	_ = cmd.PersistentFlags().String(parentFlagName, parentFlagDefault, parentDescription)

	return nil
}

func registerDtosHostNetworkVlanCreateDtoVlanID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vlanIdDescription := `VlanID vlan id`

	var vlanIdFlagName string
	if cmdPrefix == "" {
		vlanIdFlagName = "vlanID"
	} else {
		vlanIdFlagName = fmt.Sprintf("%v.vlanID", cmdPrefix)
	}

	var vlanIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(vlanIdFlagName, vlanIdFlagDefault, vlanIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDtosHostNetworkVlanCreateDtoFlags(depth int, m *models.DtosHostNetworkVlanCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressesAdded := retrieveDtosHostNetworkVlanCreateDtoAddressesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressesAdded

	err, parentAdded := retrieveDtosHostNetworkVlanCreateDtoParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	err, vlanIdAdded := retrieveDtosHostNetworkVlanCreateDtoVlanIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vlanIdAdded

	return nil, retAdded
}

func retrieveDtosHostNetworkVlanCreateDtoAddressesFlags(depth int, m *models.DtosHostNetworkVlanCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addressesFlagName := fmt.Sprintf("%v.addresses", cmdPrefix)
	if cmd.Flags().Changed(addressesFlagName) {
		// warning: addresses array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDtosHostNetworkVlanCreateDtoParentFlags(depth int, m *models.DtosHostNetworkVlanCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parentFlagName := fmt.Sprintf("%v.parent", cmdPrefix)
	if cmd.Flags().Changed(parentFlagName) {

		var parentFlagName string
		if cmdPrefix == "" {
			parentFlagName = "parent"
		} else {
			parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
		}

		parentFlagValue, err := cmd.Flags().GetString(parentFlagName)
		if err != nil {
			return err, false
		}
		m.Parent = parentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosHostNetworkVlanCreateDtoVlanIDFlags(depth int, m *models.DtosHostNetworkVlanCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vlanIdFlagName := fmt.Sprintf("%v.vlanID", cmdPrefix)
	if cmd.Flags().Changed(vlanIdFlagName) {

		var vlanIdFlagName string
		if cmdPrefix == "" {
			vlanIdFlagName = "vlanID"
		} else {
			vlanIdFlagName = fmt.Sprintf("%v.vlanID", cmdPrefix)
		}

		vlanIdFlagValue, err := cmd.Flags().GetInt64(vlanIdFlagName)
		if err != nil {
			return err, false
		}
		m.VlanID = vlanIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}