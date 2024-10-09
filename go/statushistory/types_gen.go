// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statushistory

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	// Set the height of the rows
	RowHeight float32 `json:"rowHeight"`
	// Show values on the columns
	ShowValue common.VisibilityMode    `json:"showValue"`
	Legend    common.VizLegendOptions  `json:"legend"`
	Tooltip   common.VizTooltipOptions `json:"tooltip"`
	Timezone  []common.TimeZone        `json:"timezone,omitempty"`
	// Controls the column width
	ColWidth *float64 `json:"colWidth,omitempty"`
}

func (resource Options) Equals(other Options) bool {
	if resource.RowHeight != other.RowHeight {
		return false
	}
	if resource.ShowValue != other.ShowValue {
		return false
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
	if resource.ColWidth == nil && other.ColWidth != nil || resource.ColWidth != nil && other.ColWidth == nil {
		return false
	}

	if resource.ColWidth != nil {
		if *resource.ColWidth != *other.ColWidth {
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
		Identifier: "status-history",
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
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
