// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sadov/rolctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for DtosDHCP4LeaseUpdateDto

// register flags to command
func registerModelDtosDHCP4LeaseUpdateDtoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDtosDHCP4LeaseUpdateDtoExpires(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosDHCP4LeaseUpdateDtoIP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosDHCP4LeaseUpdateDtoMac(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDtosDHCP4LeaseUpdateDtoExpires(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	expiresDescription := `Expires datetime`

	var expiresFlagName string
	if cmdPrefix == "" {
		expiresFlagName = "expires"
	} else {
		expiresFlagName = fmt.Sprintf("%v.expires", cmdPrefix)
	}

	var expiresFlagDefault string

	_ = cmd.PersistentFlags().String(expiresFlagName, expiresFlagDefault, expiresDescription)

	return nil
}

func registerDtosDHCP4LeaseUpdateDtoIP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipDescription := `IP address in ipv4 format`

	var ipFlagName string
	if cmdPrefix == "" {
		ipFlagName = "ip"
	} else {
		ipFlagName = fmt.Sprintf("%v.ip", cmdPrefix)
	}

	var ipFlagDefault string

	_ = cmd.PersistentFlags().String(ipFlagName, ipFlagDefault, ipDescription)

	return nil
}

func registerDtosDHCP4LeaseUpdateDtoMac(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macDescription := `MAC address in format like this 00-00-00-00-00`

	var macFlagName string
	if cmdPrefix == "" {
		macFlagName = "mac"
	} else {
		macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
	}

	var macFlagDefault string

	_ = cmd.PersistentFlags().String(macFlagName, macFlagDefault, macDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDtosDHCP4LeaseUpdateDtoFlags(depth int, m *models.DtosDHCP4LeaseUpdateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, expiresAdded := retrieveDtosDHCP4LeaseUpdateDtoExpiresFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || expiresAdded

	err, ipAdded := retrieveDtosDHCP4LeaseUpdateDtoIPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAdded

	err, macAdded := retrieveDtosDHCP4LeaseUpdateDtoMacFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAdded

	return nil, retAdded
}

func retrieveDtosDHCP4LeaseUpdateDtoExpiresFlags(depth int, m *models.DtosDHCP4LeaseUpdateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	expiresFlagName := fmt.Sprintf("%v.expires", cmdPrefix)
	if cmd.Flags().Changed(expiresFlagName) {

		var expiresFlagName string
		if cmdPrefix == "" {
			expiresFlagName = "expires"
		} else {
			expiresFlagName = fmt.Sprintf("%v.expires", cmdPrefix)
		}

		expiresFlagValue, err := cmd.Flags().GetString(expiresFlagName)
		if err != nil {
			return err, false
		}
		m.Expires = expiresFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosDHCP4LeaseUpdateDtoIPFlags(depth int, m *models.DtosDHCP4LeaseUpdateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipFlagName := fmt.Sprintf("%v.ip", cmdPrefix)
	if cmd.Flags().Changed(ipFlagName) {

		var ipFlagName string
		if cmdPrefix == "" {
			ipFlagName = "ip"
		} else {
			ipFlagName = fmt.Sprintf("%v.ip", cmdPrefix)
		}

		ipFlagValue, err := cmd.Flags().GetString(ipFlagName)
		if err != nil {
			return err, false
		}
		m.IP = ipFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosDHCP4LeaseUpdateDtoMacFlags(depth int, m *models.DtosDHCP4LeaseUpdateDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	macFlagName := fmt.Sprintf("%v.mac", cmdPrefix)
	if cmd.Flags().Changed(macFlagName) {

		var macFlagName string
		if cmdPrefix == "" {
			macFlagName = "mac"
		} else {
			macFlagName = fmt.Sprintf("%v.mac", cmdPrefix)
		}

		macFlagValue, err := cmd.Flags().GetString(macFlagName)
		if err != nil {
			return err, false
		}
		m.Mac = macFlagValue

		retAdded = true
	}

	return nil, retAdded
}