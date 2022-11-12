// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sadov/rolctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for DtosHostNetworkVlanUpdateDto

// register flags to command
func registerModelDtosHostNetworkVlanUpdateDtoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDtosHostNetworkVlanUpdateDtoAddresses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDtosHostNetworkVlanUpdateDtoAddresses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: addresses []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDtosHostNetworkVlanUpdateDtoFlags(depth int, m *models.DtosHostNetworkVlanUpdateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressesAdded := retrieveDtosHostNetworkVlanUpdateDtoAddressesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressesAdded

	return nil, retAdded
}

func retrieveDtosHostNetworkVlanUpdateDtoAddressesFlags(depth int, m *models.DtosHostNetworkVlanUpdateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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