// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type SeriesMapping string

const (
	SeriesMappingAuto   SeriesMapping = "auto"
	SeriesMappingManual SeriesMapping = "manual"
)

type ScatterShow string

const (
	ScatterShowPoints         ScatterShow = "points"
	ScatterShowLines          ScatterShow = "lines"
	ScatterShowPointsAndLines ScatterShow = "points+lines"
)

type XYDimensionConfig struct {
	Frame   int32    `json:"frame"`
	X       *string  `json:"x,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type FieldConfig struct {
	Show              *ScatterShow                    `json:"show,omitempty"`
	PointSize         *common.ScaleDimensionConfig    `json:"pointSize,omitempty"`
	PointColor        *common.ColorDimensionConfig    `json:"pointColor,omitempty"`
	LineColor         *common.ColorDimensionConfig    `json:"lineColor,omitempty"`
	LineWidth         *int32                          `json:"lineWidth,omitempty"`
	LineStyle         *common.LineStyle               `json:"lineStyle,omitempty"`
	Label             *common.VisibilityMode          `json:"label,omitempty"`
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
	AxisPlacement     *common.AxisPlacement           `json:"axisPlacement,omitempty"`
	AxisColorMode     *common.AxisColorMode           `json:"axisColorMode,omitempty"`
	AxisLabel         *string                         `json:"axisLabel,omitempty"`
	AxisWidth         *float64                        `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                        `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                        `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                           `json:"axisGridShow,omitempty"`
	ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	LabelValue        *common.TextDimensionConfig     `json:"labelValue,omitempty"`
	AxisCenteredZero  *bool                           `json:"axisCenteredZero,omitempty"`
}

type ScatterSeriesConfig struct {
	X                 *string                         `json:"x,omitempty"`
	Y                 *string                         `json:"y,omitempty"`
	Show              *ScatterShow                    `json:"show,omitempty"`
	PointSize         *common.ScaleDimensionConfig    `json:"pointSize,omitempty"`
	PointColor        *common.ColorDimensionConfig    `json:"pointColor,omitempty"`
	LineColor         *common.ColorDimensionConfig    `json:"lineColor,omitempty"`
	LineWidth         *int32                          `json:"lineWidth,omitempty"`
	LineStyle         *common.LineStyle               `json:"lineStyle,omitempty"`
	Label             *common.VisibilityMode          `json:"label,omitempty"`
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
	AxisPlacement     *common.AxisPlacement           `json:"axisPlacement,omitempty"`
	AxisColorMode     *common.AxisColorMode           `json:"axisColorMode,omitempty"`
	AxisLabel         *string                         `json:"axisLabel,omitempty"`
	AxisWidth         *float64                        `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                        `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                        `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                           `json:"axisGridShow,omitempty"`
	ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	Name              *string                         `json:"name,omitempty"`
	LabelValue        *common.TextDimensionConfig     `json:"labelValue,omitempty"`
	AxisCenteredZero  *bool                           `json:"axisCenteredZero,omitempty"`
}

type Options struct {
	SeriesMapping *SeriesMapping           `json:"seriesMapping,omitempty"`
	Dims          XYDimensionConfig        `json:"dims"`
	Legend        common.VizLegendOptions  `json:"legend"`
	Tooltip       common.VizTooltipOptions `json:"tooltip"`
	Series        []ScatterSeriesConfig    `json:"series"`
}

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "xychart",
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
