// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for DtosPaginatedItemsDtoDtosDHCP4LeaseDto

// register flags to command
func registerModelDtosPaginatedItemsDtoDtosDHCP4LeaseDtoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDtosPaginatedItemsDtoDtosDHCP4LeaseDtoItems(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosPaginatedItemsDtoDtosDHCP4LeaseDtoPagination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDtosPaginatedItemsDtoDtosDHCP4LeaseDtoItems(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: items []*DtosDHCP4LeaseDto array type is not supported by go-swagger cli yet

	return nil
}

func registerDtosPaginatedItemsDtoDtosDHCP4LeaseDtoPagination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var paginationFlagName string
	if cmdPrefix == "" {
		paginationFlagName = "pagination"
	} else {
		paginationFlagName = fmt.Sprintf("%v.pagination", cmdPrefix)
	}

	if err := registerModelDtosPaginationInfoDtoFlags(depth+1, paginationFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDtosPaginatedItemsDtoDtosDHCP4LeaseDtoFlags(depth int, m *models.DtosPaginatedItemsDtoDtosDHCP4LeaseDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, itemsAdded := retrieveDtosPaginatedItemsDtoDtosDHCP4LeaseDtoItemsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || itemsAdded

	err, paginationAdded := retrieveDtosPaginatedItemsDtoDtosDHCP4LeaseDtoPaginationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || paginationAdded

	return nil, retAdded
}

func retrieveDtosPaginatedItemsDtoDtosDHCP4LeaseDtoItemsFlags(depth int, m *models.DtosPaginatedItemsDtoDtosDHCP4LeaseDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	itemsFlagName := fmt.Sprintf("%v.items", cmdPrefix)
	if cmd.Flags().Changed(itemsFlagName) {
		// warning: items array type []*DtosDHCP4LeaseDto is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDtosPaginatedItemsDtoDtosDHCP4LeaseDtoPaginationFlags(depth int, m *models.DtosPaginatedItemsDtoDtosDHCP4LeaseDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	paginationFlagName := fmt.Sprintf("%v.pagination", cmdPrefix)
	if cmd.Flags().Changed(paginationFlagName) {
		// info: complex object pagination DtosPaginationInfoDto is retrieved outside this Changed() block
	}
	paginationFlagValue := m.Pagination
	if swag.IsZero(paginationFlagValue) {
		paginationFlagValue = &models.DtosPaginationInfoDto{}
	}

	err, paginationAdded := retrieveModelDtosPaginationInfoDtoFlags(depth+1, paginationFlagValue, paginationFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || paginationAdded
	if paginationAdded {
		m.Pagination = paginationFlagValue
	}

	return nil, retAdded
}
