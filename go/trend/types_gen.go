// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package trend

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// Identical to timeseries... except it does not have timezone settings
type Options struct {
	Legend  common.VizLegendOptions  `json:"legend"`
	Tooltip common.VizTooltipOptions `json:"tooltip"`
	// Name of the x field to use (defaults to first number)
	XField *string `json:"xField,omitempty"`
}

func (resource Options) Equals(other Options) bool {
	if !resource.Legend.Equals(other.Legend) {
		return false
	}
	if !resource.Tooltip.Equals(other.Tooltip) {
		return false
	}
	if resource.XField == nil && other.XField != nil || resource.XField != nil && other.XField == nil {
		return false
	}

	if resource.XField != nil {
		if *resource.XField != *other.XField {
			return false
		}
	}

	return true
}

type FieldConfig = common.GraphFieldConfig

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "trend",
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
