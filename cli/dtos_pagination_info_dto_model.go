// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/sadov/rolctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for DtosPaginationInfoDto

// register flags to command
func registerModelDtosPaginationInfoDtoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDtosPaginationInfoDtoPage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosPaginationInfoDtoSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosPaginationInfoDtoTotalCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDtosPaginationInfoDtoTotalPages(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDtosPaginationInfoDtoPage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pageDescription := `Page - page number`

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

func registerDtosPaginationInfoDtoSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := `Size - page size`

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
	}

	var sizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

func registerDtosPaginationInfoDtoTotalCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalCountDescription := `TotalCount - total number of items`

	var totalCountFlagName string
	if cmdPrefix == "" {
		totalCountFlagName = "totalCount"
	} else {
		totalCountFlagName = fmt.Sprintf("%v.totalCount", cmdPrefix)
	}

	var totalCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalCountFlagName, totalCountFlagDefault, totalCountDescription)

	return nil
}

func registerDtosPaginationInfoDtoTotalPages(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalPagesDescription := `TotalPages - total number of pages`

	var totalPagesFlagName string
	if cmdPrefix == "" {
		totalPagesFlagName = "totalPages"
	} else {
		totalPagesFlagName = fmt.Sprintf("%v.totalPages", cmdPrefix)
	}

	var totalPagesFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalPagesFlagName, totalPagesFlagDefault, totalPagesDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDtosPaginationInfoDtoFlags(depth int, m *models.DtosPaginationInfoDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, pageAdded := retrieveDtosPaginationInfoDtoPageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pageAdded

	err, sizeAdded := retrieveDtosPaginationInfoDtoSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, totalCountAdded := retrieveDtosPaginationInfoDtoTotalCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalCountAdded

	err, totalPagesAdded := retrieveDtosPaginationInfoDtoTotalPagesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalPagesAdded

	return nil, retAdded
}

func retrieveDtosPaginationInfoDtoPageFlags(depth int, m *models.DtosPaginationInfoDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pageFlagName := fmt.Sprintf("%v.page", cmdPrefix)
	if cmd.Flags().Changed(pageFlagName) {

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
		m.Page = pageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosPaginationInfoDtoSizeFlags(depth int, m *models.DtosPaginationInfoDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeFlagName := fmt.Sprintf("%v.size", cmdPrefix)
	if cmd.Flags().Changed(sizeFlagName) {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetInt64(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = sizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosPaginationInfoDtoTotalCountFlags(depth int, m *models.DtosPaginationInfoDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalCountFlagName := fmt.Sprintf("%v.totalCount", cmdPrefix)
	if cmd.Flags().Changed(totalCountFlagName) {

		var totalCountFlagName string
		if cmdPrefix == "" {
			totalCountFlagName = "totalCount"
		} else {
			totalCountFlagName = fmt.Sprintf("%v.totalCount", cmdPrefix)
		}

		totalCountFlagValue, err := cmd.Flags().GetInt64(totalCountFlagName)
		if err != nil {
			return err, false
		}
		m.TotalCount = totalCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDtosPaginationInfoDtoTotalPagesFlags(depth int, m *models.DtosPaginationInfoDto, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalPagesFlagName := fmt.Sprintf("%v.totalPages", cmdPrefix)
	if cmd.Flags().Changed(totalPagesFlagName) {

		var totalPagesFlagName string
		if cmdPrefix == "" {
			totalPagesFlagName = "totalPages"
		} else {
			totalPagesFlagName = fmt.Sprintf("%v.totalPages", cmdPrefix)
		}

		totalPagesFlagValue, err := cmd.Flags().GetInt64(totalPagesFlagName)
		if err != nil {
			return err, false
		}
		m.TotalPages = totalPagesFlagValue

		retAdded = true
	}

	return nil, retAdded
}