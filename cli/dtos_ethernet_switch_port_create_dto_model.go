// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sadov/rolctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for DtosEthernetSwitchPortCreateDto

// register flags to command
func registerModelDtosEthernetSwitchPortCreateDtoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDtosEthernetSwitchPortCreateDtoName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosEthernetSwitchPortCreateDtoPoeenabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosEthernetSwitchPortCreateDtoPoetype(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosEthernetSwitchPortCreateDtoPvid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDtosEthernetSwitchPortCreateDtoName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name for this port`

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

func registerDtosEthernetSwitchPortCreateDtoPoeenabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	poeenabledDescription := `POEEnabled POE status on a port`

	var poeenabledFlagName string
	if cmdPrefix == "" {
		poeenabledFlagName = "poeenabled"
	} else {
		poeenabledFlagName = fmt.Sprintf("%v.poeenabled", cmdPrefix)
	}

	var poeenabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(poeenabledFlagName, poeenabledFlagDefault, poeenabledDescription)

	return nil
}

func registerDtosEthernetSwitchPortCreateDtoPoetype(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	poetypeDescription := `POEType type of PoE for this port
can be: "poe", "poe+", "passive24", "none"`

	var poetypeFlagName string
	if cmdPrefix == "" {
		poetypeFlagName = "poetype"
	} else {
		poetypeFlagName = fmt.Sprintf("%v.poetype", cmdPrefix)
	}

	var poetypeFlagDefault string

	_ = cmd.PersistentFlags().String(poetypeFlagName, poetypeFlagDefault, poetypeDescription)

	return nil
}

func registerDtosEthernetSwitchPortCreateDtoPvid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pvidDescription := `PVID default VLAN id`

	var pvidFlagName string
	if cmdPrefix == "" {
		pvidFlagName = "pvid"
	} else {
		pvidFlagName = fmt.Sprintf("%v.pvid", cmdPrefix)
	}

	var pvidFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pvidFlagName, pvidFlagDefault, pvidDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDtosEthernetSwitchPortCreateDtoFlags(depth int, m *models.DtosEthernetSwitchPortCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrieveDtosEthernetSwitchPortCreateDtoNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, poeenabledAdded := retrieveDtosEthernetSwitchPortCreateDtoPoeenabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || poeenabledAdded

	err, poetypeAdded := retrieveDtosEthernetSwitchPortCreateDtoPoetypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || poetypeAdded

	err, pvidAdded := retrieveDtosEthernetSwitchPortCreateDtoPvidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pvidAdded

	return nil, retAdded
}

func retrieveDtosEthernetSwitchPortCreateDtoNameFlags(depth int, m *models.DtosEthernetSwitchPortCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveDtosEthernetSwitchPortCreateDtoPoeenabledFlags(depth int, m *models.DtosEthernetSwitchPortCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	poeenabledFlagName := fmt.Sprintf("%v.poeenabled", cmdPrefix)
	if cmd.Flags().Changed(poeenabledFlagName) {

		var poeenabledFlagName string
		if cmdPrefix == "" {
			poeenabledFlagName = "poeenabled"
		} else {
			poeenabledFlagName = fmt.Sprintf("%v.poeenabled", cmdPrefix)
		}

		poeenabledFlagValue, err := cmd.Flags().GetBool(poeenabledFlagName)
		if err != nil {
			return err, false
		}
		m.Poeenabled = poeenabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosEthernetSwitchPortCreateDtoPoetypeFlags(depth int, m *models.DtosEthernetSwitchPortCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	poetypeFlagName := fmt.Sprintf("%v.poetype", cmdPrefix)
	if cmd.Flags().Changed(poetypeFlagName) {

		var poetypeFlagName string
		if cmdPrefix == "" {
			poetypeFlagName = "poetype"
		} else {
			poetypeFlagName = fmt.Sprintf("%v.poetype", cmdPrefix)
		}

		poetypeFlagValue, err := cmd.Flags().GetString(poetypeFlagName)
		if err != nil {
			return err, false
		}
		m.Poetype = poetypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosEthernetSwitchPortCreateDtoPvidFlags(depth int, m *models.DtosEthernetSwitchPortCreateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pvidFlagName := fmt.Sprintf("%v.pvid", cmdPrefix)
	if cmd.Flags().Changed(pvidFlagName) {

		var pvidFlagName string
		if cmdPrefix == "" {
			pvidFlagName = "pvid"
		} else {
			pvidFlagName = fmt.Sprintf("%v.pvid", cmdPrefix)
		}

		pvidFlagValue, err := cmd.Flags().GetInt64(pvidFlagName)
		if err != nil {
			return err, false
		}
		m.Pvid = pvidFlagValue

		retAdded = true
	}

	return nil, retAdded
}
