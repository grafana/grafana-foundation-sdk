// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

// Auto is "table" in the UI
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

// Configuration for the Table/Auto mode
type XYDimensionConfig struct {
	Frame   int32    `json:"frame"`
	X       *string  `json:"x,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

func (resource XYDimensionConfig) Equals(other XYDimensionConfig) bool {
	if resource.Frame != other.Frame {
		return false
	}
	if resource.X == nil && other.X != nil || resource.X != nil && other.X == nil {
		return false
	}

	if resource.X != nil {
		if *resource.X != *other.X {
			return false
		}
	}

	if len(resource.Exclude) != len(other.Exclude) {
		return false
	}

	for i1 := range resource.Exclude {
		if resource.Exclude[i1] != other.Exclude[i1] {
			return false
		}
	}

	return true
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
	AxisCenteredZero  *bool                           `json:"axisCenteredZero,omitempty"`
	LabelValue        *common.TextDimensionConfig     `json:"labelValue,omitempty"`
	AxisBorderShow    *bool                           `json:"axisBorderShow,omitempty"`
}

func (resource FieldConfig) Equals(other FieldConfig) bool {
	if resource.Show == nil && other.Show != nil || resource.Show != nil && other.Show == nil {
		return false
	}

	if resource.Show != nil {
		if *resource.Show != *other.Show {
			return false
		}
	}
	if resource.PointSize == nil && other.PointSize != nil || resource.PointSize != nil && other.PointSize == nil {
		return false
	}

	if resource.PointSize != nil {
		if !resource.PointSize.Equals(*other.PointSize) {
			return false
		}
	}
	if resource.PointColor == nil && other.PointColor != nil || resource.PointColor != nil && other.PointColor == nil {
		return false
	}

	if resource.PointColor != nil {
		if !resource.PointColor.Equals(*other.PointColor) {
			return false
		}
	}
	if resource.LineColor == nil && other.LineColor != nil || resource.LineColor != nil && other.LineColor == nil {
		return false
	}

	if resource.LineColor != nil {
		if !resource.LineColor.Equals(*other.LineColor) {
			return false
		}
	}
	if resource.LineWidth == nil && other.LineWidth != nil || resource.LineWidth != nil && other.LineWidth == nil {
		return false
	}

	if resource.LineWidth != nil {
		if *resource.LineWidth != *other.LineWidth {
			return false
		}
	}
	if resource.LineStyle == nil && other.LineStyle != nil || resource.LineStyle != nil && other.LineStyle == nil {
		return false
	}

	if resource.LineStyle != nil {
		if !resource.LineStyle.Equals(*other.LineStyle) {
			return false
		}
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
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
	if resource.AxisPlacement == nil && other.AxisPlacement != nil || resource.AxisPlacement != nil && other.AxisPlacement == nil {
		return false
	}

	if resource.AxisPlacement != nil {
		if *resource.AxisPlacement != *other.AxisPlacement {
			return false
		}
	}
	if resource.AxisColorMode == nil && other.AxisColorMode != nil || resource.AxisColorMode != nil && other.AxisColorMode == nil {
		return false
	}

	if resource.AxisColorMode != nil {
		if *resource.AxisColorMode != *other.AxisColorMode {
			return false
		}
	}
	if resource.AxisLabel == nil && other.AxisLabel != nil || resource.AxisLabel != nil && other.AxisLabel == nil {
		return false
	}

	if resource.AxisLabel != nil {
		if *resource.AxisLabel != *other.AxisLabel {
			return false
		}
	}
	if resource.AxisWidth == nil && other.AxisWidth != nil || resource.AxisWidth != nil && other.AxisWidth == nil {
		return false
	}

	if resource.AxisWidth != nil {
		if *resource.AxisWidth != *other.AxisWidth {
			return false
		}
	}
	if resource.AxisSoftMin == nil && other.AxisSoftMin != nil || resource.AxisSoftMin != nil && other.AxisSoftMin == nil {
		return false
	}

	if resource.AxisSoftMin != nil {
		if *resource.AxisSoftMin != *other.AxisSoftMin {
			return false
		}
	}
	if resource.AxisSoftMax == nil && other.AxisSoftMax != nil || resource.AxisSoftMax != nil && other.AxisSoftMax == nil {
		return false
	}

	if resource.AxisSoftMax != nil {
		if *resource.AxisSoftMax != *other.AxisSoftMax {
			return false
		}
	}
	if resource.AxisGridShow == nil && other.AxisGridShow != nil || resource.AxisGridShow != nil && other.AxisGridShow == nil {
		return false
	}

	if resource.AxisGridShow != nil {
		if *resource.AxisGridShow != *other.AxisGridShow {
			return false
		}
	}
	if resource.ScaleDistribution == nil && other.ScaleDistribution != nil || resource.ScaleDistribution != nil && other.ScaleDistribution == nil {
		return false
	}

	if resource.ScaleDistribution != nil {
		if !resource.ScaleDistribution.Equals(*other.ScaleDistribution) {
			return false
		}
	}
	if resource.AxisCenteredZero == nil && other.AxisCenteredZero != nil || resource.AxisCenteredZero != nil && other.AxisCenteredZero == nil {
		return false
	}

	if resource.AxisCenteredZero != nil {
		if *resource.AxisCenteredZero != *other.AxisCenteredZero {
			return false
		}
	}
	if resource.LabelValue == nil && other.LabelValue != nil || resource.LabelValue != nil && other.LabelValue == nil {
		return false
	}

	if resource.LabelValue != nil {
		if !resource.LabelValue.Equals(*other.LabelValue) {
			return false
		}
	}
	if resource.AxisBorderShow == nil && other.AxisBorderShow != nil || resource.AxisBorderShow != nil && other.AxisBorderShow == nil {
		return false
	}

	if resource.AxisBorderShow != nil {
		if *resource.AxisBorderShow != *other.AxisBorderShow {
			return false
		}
	}

	return true
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
	AxisCenteredZero  *bool                           `json:"axisCenteredZero,omitempty"`
	Name              *string                         `json:"name,omitempty"`
	LabelValue        *common.TextDimensionConfig     `json:"labelValue,omitempty"`
	AxisBorderShow    *bool                           `json:"axisBorderShow,omitempty"`
}

func (resource ScatterSeriesConfig) Equals(other ScatterSeriesConfig) bool {
	if resource.X == nil && other.X != nil || resource.X != nil && other.X == nil {
		return false
	}

	if resource.X != nil {
		if *resource.X != *other.X {
			return false
		}
	}
	if resource.Y == nil && other.Y != nil || resource.Y != nil && other.Y == nil {
		return false
	}

	if resource.Y != nil {
		if *resource.Y != *other.Y {
			return false
		}
	}
	if resource.Show == nil && other.Show != nil || resource.Show != nil && other.Show == nil {
		return false
	}

	if resource.Show != nil {
		if *resource.Show != *other.Show {
			return false
		}
	}
	if resource.PointSize == nil && other.PointSize != nil || resource.PointSize != nil && other.PointSize == nil {
		return false
	}

	if resource.PointSize != nil {
		if !resource.PointSize.Equals(*other.PointSize) {
			return false
		}
	}
	if resource.PointColor == nil && other.PointColor != nil || resource.PointColor != nil && other.PointColor == nil {
		return false
	}

	if resource.PointColor != nil {
		if !resource.PointColor.Equals(*other.PointColor) {
			return false
		}
	}
	if resource.LineColor == nil && other.LineColor != nil || resource.LineColor != nil && other.LineColor == nil {
		return false
	}

	if resource.LineColor != nil {
		if !resource.LineColor.Equals(*other.LineColor) {
			return false
		}
	}
	if resource.LineWidth == nil && other.LineWidth != nil || resource.LineWidth != nil && other.LineWidth == nil {
		return false
	}

	if resource.LineWidth != nil {
		if *resource.LineWidth != *other.LineWidth {
			return false
		}
	}
	if resource.LineStyle == nil && other.LineStyle != nil || resource.LineStyle != nil && other.LineStyle == nil {
		return false
	}

	if resource.LineStyle != nil {
		if !resource.LineStyle.Equals(*other.LineStyle) {
			return false
		}
	}
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
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
	if resource.AxisPlacement == nil && other.AxisPlacement != nil || resource.AxisPlacement != nil && other.AxisPlacement == nil {
		return false
	}

	if resource.AxisPlacement != nil {
		if *resource.AxisPlacement != *other.AxisPlacement {
			return false
		}
	}
	if resource.AxisColorMode == nil && other.AxisColorMode != nil || resource.AxisColorMode != nil && other.AxisColorMode == nil {
		return false
	}

	if resource.AxisColorMode != nil {
		if *resource.AxisColorMode != *other.AxisColorMode {
			return false
		}
	}
	if resource.AxisLabel == nil && other.AxisLabel != nil || resource.AxisLabel != nil && other.AxisLabel == nil {
		return false
	}

	if resource.AxisLabel != nil {
		if *resource.AxisLabel != *other.AxisLabel {
			return false
		}
	}
	if resource.AxisWidth == nil && other.AxisWidth != nil || resource.AxisWidth != nil && other.AxisWidth == nil {
		return false
	}

	if resource.AxisWidth != nil {
		if *resource.AxisWidth != *other.AxisWidth {
			return false
		}
	}
	if resource.AxisSoftMin == nil && other.AxisSoftMin != nil || resource.AxisSoftMin != nil && other.AxisSoftMin == nil {
		return false
	}

	if resource.AxisSoftMin != nil {
		if *resource.AxisSoftMin != *other.AxisSoftMin {
			return false
		}
	}
	if resource.AxisSoftMax == nil && other.AxisSoftMax != nil || resource.AxisSoftMax != nil && other.AxisSoftMax == nil {
		return false
	}

	if resource.AxisSoftMax != nil {
		if *resource.AxisSoftMax != *other.AxisSoftMax {
			return false
		}
	}
	if resource.AxisGridShow == nil && other.AxisGridShow != nil || resource.AxisGridShow != nil && other.AxisGridShow == nil {
		return false
	}

	if resource.AxisGridShow != nil {
		if *resource.AxisGridShow != *other.AxisGridShow {
			return false
		}
	}
	if resource.ScaleDistribution == nil && other.ScaleDistribution != nil || resource.ScaleDistribution != nil && other.ScaleDistribution == nil {
		return false
	}

	if resource.ScaleDistribution != nil {
		if !resource.ScaleDistribution.Equals(*other.ScaleDistribution) {
			return false
		}
	}
	if resource.AxisCenteredZero == nil && other.AxisCenteredZero != nil || resource.AxisCenteredZero != nil && other.AxisCenteredZero == nil {
		return false
	}

	if resource.AxisCenteredZero != nil {
		if *resource.AxisCenteredZero != *other.AxisCenteredZero {
			return false
		}
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}
	if resource.LabelValue == nil && other.LabelValue != nil || resource.LabelValue != nil && other.LabelValue == nil {
		return false
	}

	if resource.LabelValue != nil {
		if !resource.LabelValue.Equals(*other.LabelValue) {
			return false
		}
	}
	if resource.AxisBorderShow == nil && other.AxisBorderShow != nil || resource.AxisBorderShow != nil && other.AxisBorderShow == nil {
		return false
	}

	if resource.AxisBorderShow != nil {
		if *resource.AxisBorderShow != *other.AxisBorderShow {
			return false
		}
	}

	return true
}

type Options struct {
	SeriesMapping *SeriesMapping `json:"seriesMapping,omitempty"`
	// Table Mode (auto)
	Dims    XYDimensionConfig        `json:"dims"`
	Legend  common.VizLegendOptions  `json:"legend"`
	Tooltip common.VizTooltipOptions `json:"tooltip"`
	// Manual Mode
	Series []ScatterSeriesConfig `json:"series"`
}

func (resource Options) Equals(other Options) bool {
	if resource.SeriesMapping == nil && other.SeriesMapping != nil || resource.SeriesMapping != nil && other.SeriesMapping == nil {
		return false
	}

	if resource.SeriesMapping != nil {
		if *resource.SeriesMapping != *other.SeriesMapping {
			return false
		}
	}
	if !resource.Dims.Equals(other.Dims) {
		return false
	}
	if !resource.Legend.Equals(other.Legend) {
		return false
	}
	if !resource.Tooltip.Equals(other.Tooltip) {
		return false
	}

	if len(resource.Series) != len(other.Series) {
		return false
	}

	for i1 := range resource.Series {
		if !resource.Series[i1].Equals(other.Series[i1]) {
			return false
		}
	}

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "xychart",
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
