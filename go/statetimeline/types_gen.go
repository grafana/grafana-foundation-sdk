// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statetimeline

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	// Show timeline values on chart
	ShowValue common.VisibilityMode `json:"showValue"`
	// Controls the row height
	RowHeight float64 `json:"rowHeight"`
	// Merge equal consecutive values
	MergeValues *bool                    `json:"mergeValues,omitempty"`
	Legend      common.VizLegendOptions  `json:"legend"`
	Tooltip     common.VizTooltipOptions `json:"tooltip"`
	Timezone    []common.TimeZone        `json:"timezone,omitempty"`
	// Controls value alignment on the timelines
	AlignValue *common.TimelineValueAlignment `json:"alignValue,omitempty"`
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
	// Field "showValue"
	if fields["showValue"] != nil {
		if string(fields["showValue"]) != "null" {
			if err := json.Unmarshal(fields["showValue"], &resource.ShowValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showValue", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showValue", errors.New("required field is null"))...)

		}
		delete(fields, "showValue")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showValue", errors.New("required field is missing from input"))...)
	}
	// Field "rowHeight"
	if fields["rowHeight"] != nil {
		if string(fields["rowHeight"]) != "null" {
			if err := json.Unmarshal(fields["rowHeight"], &resource.RowHeight); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rowHeight", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("rowHeight", errors.New("required field is null"))...)

		}
		delete(fields, "rowHeight")
	} else {
		errs = append(errs, cog.MakeBuildErrors("rowHeight", errors.New("required field is missing from input"))...)
	}
	// Field "mergeValues"
	if fields["mergeValues"] != nil {
		if string(fields["mergeValues"]) != "null" {
			if err := json.Unmarshal(fields["mergeValues"], &resource.MergeValues); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mergeValues", err)...)
			}

		}
		delete(fields, "mergeValues")

	}
	// Field "legend"
	if fields["legend"] != nil {
		if string(fields["legend"]) != "null" {

			resource.Legend = common.VizLegendOptions{}
			if err := resource.Legend.UnmarshalJSONStrict(fields["legend"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legend", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is null"))...)

		}
		delete(fields, "legend")
	} else {
		errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is missing from input"))...)
	}
	// Field "tooltip"
	if fields["tooltip"] != nil {
		if string(fields["tooltip"]) != "null" {

			resource.Tooltip = common.VizTooltipOptions{}
			if err := resource.Tooltip.UnmarshalJSONStrict(fields["tooltip"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is null"))...)

		}
		delete(fields, "tooltip")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is missing from input"))...)
	}
	// Field "timezone"
	if fields["timezone"] != nil {
		if string(fields["timezone"]) != "null" {

			if err := json.Unmarshal(fields["timezone"], &resource.Timezone); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timezone", err)...)
			}

		}
		delete(fields, "timezone")

	}
	// Field "alignValue"
	if fields["alignValue"] != nil {
		if string(fields["alignValue"]) != "null" {
			if err := json.Unmarshal(fields["alignValue"], &resource.AlignValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alignValue", err)...)
			}

		}
		delete(fields, "alignValue")

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
	if resource.ShowValue != other.ShowValue {
		return false
	}
	if resource.RowHeight != other.RowHeight {
		return false
	}
	if resource.MergeValues == nil && other.MergeValues != nil || resource.MergeValues != nil && other.MergeValues == nil {
		return false
	}

	if resource.MergeValues != nil {
		if *resource.MergeValues != *other.MergeValues {
			return false
		}
	}
	if !resource.Legend.Equals(other.Legend) {
		return false
	}
	if !resource.Tooltip.Equals(other.Tooltip) {
		return false
	}

	if len(resource.Timezone) != len(other.Timezone) {
		return false
	}

	for i1 := range resource.Timezone {
		if resource.Timezone[i1] != other.Timezone[i1] {
			return false
		}
	}
	if resource.AlignValue == nil && other.AlignValue != nil || resource.AlignValue != nil && other.AlignValue == nil {
		return false
	}

	if resource.AlignValue != nil {
		if *resource.AlignValue != *other.AlignValue {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if !(resource.RowHeight <= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"rowHeight",
			errors.New("must be <= 1"),
		)...)
	}
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}
	if err := resource.Tooltip.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type FieldConfig struct {
	LineWidth   *uint32                  `json:"lineWidth,omitempty"`
	HideFrom    *common.HideSeriesConfig `json:"hideFrom,omitempty"`
	FillOpacity *uint32                  `json:"fillOpacity,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FieldConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "lineWidth"
	if fields["lineWidth"] != nil {
		if string(fields["lineWidth"]) != "null" {
			if err := json.Unmarshal(fields["lineWidth"], &resource.LineWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineWidth", err)...)
			}

		}
		delete(fields, "lineWidth")

	}
	// Field "hideFrom"
	if fields["hideFrom"] != nil {
		if string(fields["hideFrom"]) != "null" {

			resource.HideFrom = &common.HideSeriesConfig{}
			if err := resource.HideFrom.UnmarshalJSONStrict(fields["hideFrom"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
			}

		}
		delete(fields, "hideFrom")

	}
	// Field "fillOpacity"
	if fields["fillOpacity"] != nil {
		if string(fields["fillOpacity"]) != "null" {
			if err := json.Unmarshal(fields["fillOpacity"], &resource.FillOpacity); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillOpacity", err)...)
			}

		}
		delete(fields, "fillOpacity")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FieldConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FieldConfig` objects.
func (resource FieldConfig) Equals(other FieldConfig) bool {
	if resource.LineWidth == nil && other.LineWidth != nil || resource.LineWidth != nil && other.LineWidth == nil {
		return false
	}

	if resource.LineWidth != nil {
		if *resource.LineWidth != *other.LineWidth {
			return false
		}
	}
	if resource.HideFrom == nil && other.HideFrom != nil || resource.HideFrom != nil && other.HideFrom == nil {
		return false
	}

	if resource.HideFrom != nil {
		if !resource.HideFrom.Equals(*other.HideFrom) {
			return false
		}
	}
	if resource.FillOpacity == nil && other.FillOpacity != nil || resource.FillOpacity != nil && other.FillOpacity == nil {
		return false
	}

	if resource.FillOpacity != nil {
		if *resource.FillOpacity != *other.FillOpacity {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `FieldConfig` fields for violations and returns them.
func (resource FieldConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.LineWidth != nil {
		if !(*resource.LineWidth <= 10) {
			errs = append(errs, cog.MakeBuildErrors(
				"lineWidth",
				errors.New("must be <= 10"),
			)...)
		}
	}
	if resource.HideFrom != nil {
		if err := resource.HideFrom.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
		}
	}
	if resource.FillOpacity != nil {
		if !(*resource.FillOpacity <= 100) {
			errs = append(errs, cog.MakeBuildErrors(
				"fillOpacity",
				errors.New("must be <= 100"),
			)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to state-timeline panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "state-timeline",
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
