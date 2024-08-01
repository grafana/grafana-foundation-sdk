// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package timeseries

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	Timezone []common.TimeZone        `json:"timezone,omitempty"`
	Legend   common.VizLegendOptions  `json:"legend"`
	Tooltip  common.VizTooltipOptions `json:"tooltip"`
}

type FieldConfig = common.GraphFieldConfig

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "timeseries",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
		FieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := FieldConfig{}

			if err := json.Unmarshal(raw, &fieldConfig); err != nil {
				return nil, err
			}

			return fieldConfig, nil
		},
	}
}
