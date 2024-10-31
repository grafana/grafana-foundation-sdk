// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	// Represents the index of the selected frame
	FrameIndex float64 `json:"frameIndex"`
	// Controls whether the panel should show the header
	ShowHeader bool `json:"showHeader"`
	// Controls whether the header should show icons for the column types
	ShowTypeIcons *bool `json:"showTypeIcons,omitempty"`
	// Used to control row sorting
	SortBy []common.TableSortByFieldState `json:"sortBy,omitempty"`
	// Controls footer options
	Footer *common.TableFooterOptions `json:"footer,omitempty"`
	// Controls the height of the rows
	CellHeight *common.TableCellHeight `json:"cellHeight,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "frameIndex"
	if fields["frameIndex"] != nil {
		if string(fields["frameIndex"]) != "null" {
			if err := json.Unmarshal(fields["frameIndex"], &resource.FrameIndex); err != nil {
				errs = append(errs, cog.MakeBuildErrors("frameIndex", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("frameIndex", errors.New("required field is null"))...)

		}
		delete(fields, "frameIndex")
	} else {
		errs = append(errs, cog.MakeBuildErrors("frameIndex", errors.New("required field is missing from input"))...)
	}
	// Field "showHeader"
	if fields["showHeader"] != nil {
		if string(fields["showHeader"]) != "null" {
			if err := json.Unmarshal(fields["showHeader"], &resource.ShowHeader); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showHeader", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showHeader", errors.New("required field is null"))...)

		}
		delete(fields, "showHeader")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showHeader", errors.New("required field is missing from input"))...)
	}
	// Field "showTypeIcons"
	if fields["showTypeIcons"] != nil {
		if string(fields["showTypeIcons"]) != "null" {
			if err := json.Unmarshal(fields["showTypeIcons"], &resource.ShowTypeIcons); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showTypeIcons", err)...)
			}

		}
		delete(fields, "showTypeIcons")

	}
	// Field "sortBy"
	if fields["sortBy"] != nil {
		if string(fields["sortBy"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["sortBy"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 common.TableSortByFieldState

				result1 = common.TableSortByFieldState{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("sortBy["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.SortBy = append(resource.SortBy, result1)
			}

		}
		delete(fields, "sortBy")

	}
	// Field "footer"
	if fields["footer"] != nil {
		if string(fields["footer"]) != "null" {

			resource.Footer = &common.TableFooterOptions{}
			if err := resource.Footer.UnmarshalJSONStrict(fields["footer"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("footer", err)...)
			}

		}
		delete(fields, "footer")

	}
	// Field "cellHeight"
	if fields["cellHeight"] != nil {
		if string(fields["cellHeight"]) != "null" {
			if err := json.Unmarshal(fields["cellHeight"], &resource.CellHeight); err != nil {
				errs = append(errs, cog.MakeBuildErrors("cellHeight", err)...)
			}

		}
		delete(fields, "cellHeight")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Options` objects.
func (resource Options) Equals(other Options) bool {
	if resource.FrameIndex != other.FrameIndex {
		return false
	}
	if resource.ShowHeader != other.ShowHeader {
		return false
	}
	if resource.ShowTypeIcons == nil && other.ShowTypeIcons != nil || resource.ShowTypeIcons != nil && other.ShowTypeIcons == nil {
		return false
	}

	if resource.ShowTypeIcons != nil {
		if *resource.ShowTypeIcons != *other.ShowTypeIcons {
			return false
		}
	}

	if len(resource.SortBy) != len(other.SortBy) {
		return false
	}

	for i1 := range resource.SortBy {
		if !resource.SortBy[i1].Equals(other.SortBy[i1]) {
			return false
		}
	}
	if resource.Footer == nil && other.Footer != nil || resource.Footer != nil && other.Footer == nil {
		return false
	}

	if resource.Footer != nil {
		if !resource.Footer.Equals(*other.Footer) {
			return false
		}
	}
	if resource.CellHeight == nil && other.CellHeight != nil || resource.CellHeight != nil && other.CellHeight == nil {
		return false
	}

	if resource.CellHeight != nil {
		if *resource.CellHeight != *other.CellHeight {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.SortBy {
		if err := resource.SortBy[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("sortBy["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Footer != nil {
		if err := resource.Footer.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("footer", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type FieldConfig = common.TableFieldOptions

// VariantConfig returns the configuration related to table panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "table",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := json.Unmarshal(raw, options); err != nil {
				return nil, err
			}

			return options, nil
		},
		StrictOptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := options.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return options, nil
		},
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := json.Unmarshal(raw, fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
		StrictFieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := fieldConfig.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
