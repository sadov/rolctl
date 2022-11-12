// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sadov/rolctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for DtosDeviceTemplateNetworkDto

// register flags to command
func registerModelDtosDeviceTemplateNetworkDtoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDtosDeviceTemplateNetworkDtoManagement(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosDeviceTemplateNetworkDtoName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosDeviceTemplateNetworkDtoNetBoot(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosDeviceTemplateNetworkDtoPoein(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDtosDeviceTemplateNetworkDtoManagement(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	managementDescription := `Management only one network interface can be mark as management`

	var managementFlagName string
	if cmdPrefix == "" {
		managementFlagName = "management"
	} else {
		managementFlagName = fmt.Sprintf("%v.management", cmdPrefix)
	}

	var managementFlagDefault bool

	_ = cmd.PersistentFlags().Bool(managementFlagName, managementFlagDefault, managementDescription)

	return nil
}

func registerDtosDeviceTemplateNetworkDtoName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of network interface. This field is unique within device template network interfaces`

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

func registerDtosDeviceTemplateNetworkDtoNetBoot(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	netBootDescription := `NetBoot flags whether the interface can be loaded over the network`

	var netBootFlagName string
	if cmdPrefix == "" {
		netBootFlagName = "netBoot"
	} else {
		netBootFlagName = fmt.Sprintf("%v.netBoot", cmdPrefix)
	}

	var netBootFlagDefault bool

	_ = cmd.PersistentFlags().Bool(netBootFlagName, netBootFlagDefault, netBootDescription)

	return nil
}

func registerDtosDeviceTemplateNetworkDtoPoein(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	poeinDescription := `POEIn only one network interface can be mark as POEIn`

	var poeinFlagName string
	if cmdPrefix == "" {
		poeinFlagName = "poein"
	} else {
		poeinFlagName = fmt.Sprintf("%v.poein", cmdPrefix)
	}

	var poeinFlagDefault bool

	_ = cmd.PersistentFlags().Bool(poeinFlagName, poeinFlagDefault, poeinDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDtosDeviceTemplateNetworkDtoFlags(depth int, m *models.DtosDeviceTemplateNetworkDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, managementAdded := retrieveDtosDeviceTemplateNetworkDtoManagementFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || managementAdded

	err, nameAdded := retrieveDtosDeviceTemplateNetworkDtoNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, netBootAdded := retrieveDtosDeviceTemplateNetworkDtoNetBootFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || netBootAdded

	err, poeinAdded := retrieveDtosDeviceTemplateNetworkDtoPoeinFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || poeinAdded

	return nil, retAdded
}

func retrieveDtosDeviceTemplateNetworkDtoManagementFlags(depth int, m *models.DtosDeviceTemplateNetworkDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	managementFlagName := fmt.Sprintf("%v.management", cmdPrefix)
	if cmd.Flags().Changed(managementFlagName) {

		var managementFlagName string
		if cmdPrefix == "" {
			managementFlagName = "management"
		} else {
			managementFlagName = fmt.Sprintf("%v.management", cmdPrefix)
		}

		managementFlagValue, err := cmd.Flags().GetBool(managementFlagName)
		if err != nil {
			return err, false
		}
		m.Management = managementFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosDeviceTemplateNetworkDtoNameFlags(depth int, m *models.DtosDeviceTemplateNetworkDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDtosDeviceTemplateNetworkDtoNetBootFlags(depth int, m *models.DtosDeviceTemplateNetworkDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	netBootFlagName := fmt.Sprintf("%v.netBoot", cmdPrefix)
	if cmd.Flags().Changed(netBootFlagName) {

		var netBootFlagName string
		if cmdPrefix == "" {
			netBootFlagName = "netBoot"
		} else {
			netBootFlagName = fmt.Sprintf("%v.netBoot", cmdPrefix)
		}

		netBootFlagValue, err := cmd.Flags().GetBool(netBootFlagName)
		if err != nil {
			return err, false
		}
		m.NetBoot = netBootFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosDeviceTemplateNetworkDtoPoeinFlags(depth int, m *models.DtosDeviceTemplateNetworkDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	poeinFlagName := fmt.Sprintf("%v.poein", cmdPrefix)
	if cmd.Flags().Changed(poeinFlagName) {

		var poeinFlagName string
		if cmdPrefix == "" {
			poeinFlagName = "poein"
		} else {
			poeinFlagName = fmt.Sprintf("%v.poein", cmdPrefix)
		}

		poeinFlagValue, err := cmd.Flags().GetBool(poeinFlagName)
		if err != nil {
			return err, false
		}
		m.Poein = poeinFlagValue

		retAdded = true
	}

	return nil, retAdded
}
