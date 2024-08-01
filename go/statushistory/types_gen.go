// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package statushistory

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
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

type FieldConfig struct {
	LineWidth   *uint32                  `json:"lineWidth,omitempty"`
	HideFrom    *common.HideSeriesConfig `json:"hideFrom,omitempty"`
	FillOpacity *uint32                  `json:"fillOpacity,omitempty"`
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "status-history",
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
