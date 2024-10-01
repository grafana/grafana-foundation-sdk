// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statetimeline

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	// Show timeline values on chart
	ShowValue common.VisibilityMode `json:"showValue"`
	// Controls the row height
	RowHeight float64 `json:"rowHeight"`
	// Merge equal consecutive values
	MergeValues *bool `json:"mergeValues,omitempty"`
	// Controls value alignment on the timelines
	AlignValue *common.TimelineValueAlignment `json:"alignValue,omitempty"`
	Legend     common.VizLegendOptions        `json:"legend"`
	Tooltip    common.VizTooltipOptions       `json:"tooltip"`
	Timezone   []common.TimeZone              `json:"timezone,omitempty"`
	// Enables pagination when > 0
	PerPage *float64 `json:"perPage,omitempty"`
}

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
	if resource.AlignValue == nil && other.AlignValue != nil || resource.AlignValue != nil && other.AlignValue == nil {
		return false
	}

	if resource.AlignValue != nil {
		if *resource.AlignValue != *other.AlignValue {
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
	if resource.PerPage == nil && other.PerPage != nil || resource.PerPage != nil && other.PerPage == nil {
		return false
	}

	if resource.PerPage != nil {
		if *resource.PerPage != *other.PerPage {
			return false
		}
	}

	return true
}

type FieldConfig struct {
	LineWidth   *uint32                  `json:"lineWidth,omitempty"`
	HideFrom    *common.HideSeriesConfig `json:"hideFrom,omitempty"`
	FillOpacity *uint32                  `json:"fillOpacity,omitempty"`
}

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
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := json.Unmarshal(raw, fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
	}
}
