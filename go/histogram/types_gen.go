// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package histogram

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	// Size of each bucket
	BucketSize *int32 `json:"bucketSize,omitempty"`
	// Offset buckets by this amount
	BucketOffset *int32                   `json:"bucketOffset,omitempty"`
	Legend       common.VizLegendOptions  `json:"legend"`
	Tooltip      common.VizTooltipOptions `json:"tooltip"`
	// Combines multiple series into a single histogram
	Combine *bool `json:"combine,omitempty"`
}

type FieldConfig struct {
	// Controls line width of the bars.
	LineWidth *uint32 `json:"lineWidth,omitempty"`
	// Controls the fill opacity of the bars.
	FillOpacity       *uint32                         `json:"fillOpacity,omitempty"`
	AxisPlacement     *common.AxisPlacement           `json:"axisPlacement,omitempty"`
	AxisColorMode     *common.AxisColorMode           `json:"axisColorMode,omitempty"`
	AxisLabel         *string                         `json:"axisLabel,omitempty"`
	AxisWidth         *float64                        `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                        `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                        `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                           `json:"axisGridShow,omitempty"`
	ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	AxisCenteredZero  *bool                           `json:"axisCenteredZero,omitempty"`
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
	// Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
	// Gradient appearance is influenced by the Fill opacity setting.
	GradientMode   *common.GraphGradientMode `json:"gradientMode,omitempty"`
	AxisBorderShow *bool                     `json:"axisBorderShow,omitempty"`
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "histogram",
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
