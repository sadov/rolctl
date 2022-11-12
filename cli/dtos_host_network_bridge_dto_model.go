// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sadov/rolctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for DtosHostNetworkBridgeDto

// register flags to command
func registerModelDtosHostNetworkBridgeDtoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDtosHostNetworkBridgeDtoAddresses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosHostNetworkBridgeDtoName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosHostNetworkBridgeDtoSlaves(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDtosHostNetworkBridgeDtoAddresses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: addresses []string array type is not supported by go-swagger cli yet

	return nil
}

func registerDtosHostNetworkBridgeDtoName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name interface full name`

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

func registerDtosHostNetworkBridgeDtoSlaves(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: slaves []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDtosHostNetworkBridgeDtoFlags(depth int, m *models.DtosHostNetworkBridgeDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addressesAdded := retrieveDtosHostNetworkBridgeDtoAddressesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addressesAdded

	err, nameAdded := retrieveDtosHostNetworkBridgeDtoNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, slavesAdded := retrieveDtosHostNetworkBridgeDtoSlavesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slavesAdded

	return nil, retAdded
}

func retrieveDtosHostNetworkBridgeDtoAddressesFlags(depth int, m *models.DtosHostNetworkBridgeDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDtosHostNetworkBridgeDtoNameFlags(depth int, m *models.DtosHostNetworkBridgeDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

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

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosHostNetworkBridgeDtoSlavesFlags(depth int, m *models.DtosHostNetworkBridgeDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	slavesFlagName := fmt.Sprintf("%v.slaves", cmdPrefix)
	if cmd.Flags().Changed(slavesFlagName) {
		// warning: slaves array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}