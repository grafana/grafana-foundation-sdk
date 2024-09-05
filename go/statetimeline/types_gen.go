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

type FieldConfig struct {
	LineWidth   *uint32                  `json:"lineWidth,omitempty"`
	HideFrom    *common.HideSeriesConfig `json:"hideFrom,omitempty"`
	FillOpacity *uint32                  `json:"fillOpacity,omitempty"`
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "state-timeline",
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
