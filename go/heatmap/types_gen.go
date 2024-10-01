// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

// Controls the color mode of the heatmap
type HeatmapColorMode string

const (
	HeatmapColorModeOpacity HeatmapColorMode = "opacity"
	HeatmapColorModeScheme  HeatmapColorMode = "scheme"
)

// Controls the color scale of the heatmap
type HeatmapColorScale string

const (
	HeatmapColorScaleLinear      HeatmapColorScale = "linear"
	HeatmapColorScaleExponential HeatmapColorScale = "exponential"
)

// Controls various color options
type HeatmapColorOptions struct {
	// Sets the color mode
	Mode *HeatmapColorMode `json:"mode,omitempty"`
	// Controls the color scheme used
	Scheme string `json:"scheme"`
	// Controls the color fill when in opacity mode
	Fill string `json:"fill"`
	// Controls the color scale
	Scale *HeatmapColorScale `json:"scale,omitempty"`
	// Controls the exponent when scale is set to exponential
	Exponent float32 `json:"exponent"`
	// Controls the number of color steps
	Steps uint64 `json:"steps"`
	// Reverses the color scheme
	Reverse bool `json:"reverse"`
	// Sets the minimum value for the color scale
	Min *float32 `json:"min,omitempty"`
	// Sets the maximum value for the color scale
	Max *float32 `json:"max,omitempty"`
}

func (resource HeatmapColorOptions) Equals(other HeatmapColorOptions) bool {
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
			return false
		}
	}
	if resource.Scheme != other.Scheme {
		return false
	}
	if resource.Fill != other.Fill {
		return false
	}
	if resource.Scale == nil && other.Scale != nil || resource.Scale != nil && other.Scale == nil {
		return false
	}

	if resource.Scale != nil {
		if *resource.Scale != *other.Scale {
			return false
		}
	}
	if resource.Exponent != other.Exponent {
		return false
	}
	if resource.Steps != other.Steps {
		return false
	}
	if resource.Reverse != other.Reverse {
		return false
	}
	if resource.Min == nil && other.Min != nil || resource.Min != nil && other.Min == nil {
		return false
	}

	if resource.Min != nil {
		if *resource.Min != *other.Min {
			return false
		}
	}
	if resource.Max == nil && other.Max != nil || resource.Max != nil && other.Max == nil {
		return false
	}

	if resource.Max != nil {
		if *resource.Max != *other.Max {
			return false
		}
	}

	return true
}

// Configuration options for the yAxis
type YAxisConfig struct {
	// Sets the yAxis unit
	Unit *string `json:"unit,omitempty"`
	// Reverses the yAxis
	Reverse *bool `json:"reverse,omitempty"`
	// Controls the number of decimals for yAxis values
	Decimals *float32 `json:"decimals,omitempty"`
	// Sets the minimum value for the yAxis
	Min               *float32                        `json:"min,omitempty"`
	AxisPlacement     *common.AxisPlacement           `json:"axisPlacement,omitempty"`
	AxisColorMode     *common.AxisColorMode           `json:"axisColorMode,omitempty"`
	AxisLabel         *string                         `json:"axisLabel,omitempty"`
	AxisWidth         *float64                        `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                        `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                        `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                           `json:"axisGridShow,omitempty"`
	ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	AxisCenteredZero  *bool                           `json:"axisCenteredZero,omitempty"`
	// Sets the maximum value for the yAxis
	Max            *float32 `json:"max,omitempty"`
	AxisBorderShow *bool    `json:"axisBorderShow,omitempty"`
}

func (resource YAxisConfig) Equals(other YAxisConfig) bool {
	if resource.Unit == nil && other.Unit != nil || resource.Unit != nil && other.Unit == nil {
		return false
	}

	if resource.Unit != nil {
		if *resource.Unit != *other.Unit {
			return false
		}
	}
	if resource.Reverse == nil && other.Reverse != nil || resource.Reverse != nil && other.Reverse == nil {
		return false
	}

	if resource.Reverse != nil {
		if *resource.Reverse != *other.Reverse {
			return false
		}
	}
	if resource.Decimals == nil && other.Decimals != nil || resource.Decimals != nil && other.Decimals == nil {
		return false
	}

	if resource.Decimals != nil {
		if *resource.Decimals != *other.Decimals {
			return false
		}
	}
	if resource.Min == nil && other.Min != nil || resource.Min != nil && other.Min == nil {
		return false
	}

	if resource.Min != nil {
		if *resource.Min != *other.Min {
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
	if resource.Max == nil && other.Max != nil || resource.Max != nil && other.Max == nil {
		return false
	}

	if resource.Max != nil {
		if *resource.Max != *other.Max {
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

// Controls cell value options
type CellValues struct {
	// Controls the cell value unit
	Unit *string `json:"unit,omitempty"`
	// Controls the number of decimals for cell values
	Decimals *float32 `json:"decimals,omitempty"`
}

func (resource CellValues) Equals(other CellValues) bool {
	if resource.Unit == nil && other.Unit != nil || resource.Unit != nil && other.Unit == nil {
		return false
	}

	if resource.Unit != nil {
		if *resource.Unit != *other.Unit {
			return false
		}
	}
	if resource.Decimals == nil && other.Decimals != nil || resource.Decimals != nil && other.Decimals == nil {
		return false
	}

	if resource.Decimals != nil {
		if *resource.Decimals != *other.Decimals {
			return false
		}
	}

	return true
}

// Controls the value filter range
type FilterValueRange struct {
	// Sets the filter range to values less than or equal to the given value
	Le *float32 `json:"le,omitempty"`
	// Sets the filter range to values greater than or equal to the given value
	Ge *float32 `json:"ge,omitempty"`
}

func (resource FilterValueRange) Equals(other FilterValueRange) bool {
	if resource.Le == nil && other.Le != nil || resource.Le != nil && other.Le == nil {
		return false
	}

	if resource.Le != nil {
		if *resource.Le != *other.Le {
			return false
		}
	}
	if resource.Ge == nil && other.Ge != nil || resource.Ge != nil && other.Ge == nil {
		return false
	}

	if resource.Ge != nil {
		if *resource.Ge != *other.Ge {
			return false
		}
	}

	return true
}

// Controls tooltip options
type HeatmapTooltip struct {
	// Controls how the tooltip is shown
	Mode      common.TooltipDisplayMode `json:"mode"`
	MaxHeight *float64                  `json:"maxHeight,omitempty"`
	MaxWidth  *float64                  `json:"maxWidth,omitempty"`
	// Controls if the tooltip shows a histogram of the y-axis values
	YHistogram *bool `json:"yHistogram,omitempty"`
	// Controls if the tooltip shows a color scale in header
	ShowColorScale *bool `json:"showColorScale,omitempty"`
}

func (resource HeatmapTooltip) Equals(other HeatmapTooltip) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.MaxHeight == nil && other.MaxHeight != nil || resource.MaxHeight != nil && other.MaxHeight == nil {
		return false
	}

	if resource.MaxHeight != nil {
		if *resource.MaxHeight != *other.MaxHeight {
			return false
		}
	}
	if resource.MaxWidth == nil && other.MaxWidth != nil || resource.MaxWidth != nil && other.MaxWidth == nil {
		return false
	}

	if resource.MaxWidth != nil {
		if *resource.MaxWidth != *other.MaxWidth {
			return false
		}
	}
	if resource.YHistogram == nil && other.YHistogram != nil || resource.YHistogram != nil && other.YHistogram == nil {
		return false
	}

	if resource.YHistogram != nil {
		if *resource.YHistogram != *other.YHistogram {
			return false
		}
	}
	if resource.ShowColorScale == nil && other.ShowColorScale != nil || resource.ShowColorScale != nil && other.ShowColorScale == nil {
		return false
	}

	if resource.ShowColorScale != nil {
		if *resource.ShowColorScale != *other.ShowColorScale {
			return false
		}
	}

	return true
}

// Controls legend options
type HeatmapLegend struct {
	// Controls if the legend is shown
	Show bool `json:"show"`
}

func (resource HeatmapLegend) Equals(other HeatmapLegend) bool {
	if resource.Show != other.Show {
		return false
	}

	return true
}

// Controls exemplar options
type ExemplarConfig struct {
	// Sets the color of the exemplar markers
	Color string `json:"color"`
}

func (resource ExemplarConfig) Equals(other ExemplarConfig) bool {
	if resource.Color != other.Color {
		return false
	}

	return true
}

// Controls frame rows options
type RowsHeatmapOptions struct {
	// Sets the name of the cell when not calculating from data
	Value *string `json:"value,omitempty"`
	// Controls tick alignment when not calculating from data
	Layout *common.HeatmapCellLayout `json:"layout,omitempty"`
}

func (resource RowsHeatmapOptions) Equals(other RowsHeatmapOptions) bool {
	if resource.Value == nil && other.Value != nil || resource.Value != nil && other.Value == nil {
		return false
	}

	if resource.Value != nil {
		if *resource.Value != *other.Value {
			return false
		}
	}
	if resource.Layout == nil && other.Layout != nil || resource.Layout != nil && other.Layout == nil {
		return false
	}

	if resource.Layout != nil {
		if *resource.Layout != *other.Layout {
			return false
		}
	}

	return true
}

type Options struct {
	// Controls if the heatmap should be calculated from data
	Calculate *bool `json:"calculate,omitempty"`
	// Calculation options for the heatmap
	Calculation *common.HeatmapCalculationOptions `json:"calculation,omitempty"`
	// Controls the color options
	Color HeatmapColorOptions `json:"color"`
	// Filters values between a given range
	FilterValues *FilterValueRange `json:"filterValues,omitempty"`
	// Controls tick alignment and value name when not calculating from data
	RowsFrame *RowsHeatmapOptions `json:"rowsFrame,omitempty"`
	// | *{
	// 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
	// }
	// Controls the display of the value in the cell
	ShowValue common.VisibilityMode `json:"showValue"`
	// Controls gap between cells
	CellGap *uint8 `json:"cellGap,omitempty"`
	// Controls cell radius
	CellRadius *float32 `json:"cellRadius,omitempty"`
	// Controls cell value unit
	CellValues *CellValues `json:"cellValues,omitempty"`
	// Controls yAxis placement
	YAxis YAxisConfig `json:"yAxis"`
	// | *{
	// 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
	// }
	// Controls legend options
	Legend HeatmapLegend `json:"legend"`
	// Controls tooltip options
	Tooltip HeatmapTooltip `json:"tooltip"`
	// Controls exemplar options
	Exemplars ExemplarConfig `json:"exemplars"`
}

func (resource Options) Equals(other Options) bool {
	if resource.Calculate == nil && other.Calculate != nil || resource.Calculate != nil && other.Calculate == nil {
		return false
	}

	if resource.Calculate != nil {
		if *resource.Calculate != *other.Calculate {
			return false
		}
	}
	if resource.Calculation == nil && other.Calculation != nil || resource.Calculation != nil && other.Calculation == nil {
		return false
	}

	if resource.Calculation != nil {
		if !resource.Calculation.Equals(*other.Calculation) {
			return false
		}
	}
	if !resource.Color.Equals(other.Color) {
		return false
	}
	if resource.FilterValues == nil && other.FilterValues != nil || resource.FilterValues != nil && other.FilterValues == nil {
		return false
	}

	if resource.FilterValues != nil {
		if !resource.FilterValues.Equals(*other.FilterValues) {
			return false
		}
	}
	if resource.RowsFrame == nil && other.RowsFrame != nil || resource.RowsFrame != nil && other.RowsFrame == nil {
		return false
	}

	if resource.RowsFrame != nil {
		if !resource.RowsFrame.Equals(*other.RowsFrame) {
			return false
		}
	}
	if resource.ShowValue != other.ShowValue {
		return false
	}
	if resource.CellGap == nil && other.CellGap != nil || resource.CellGap != nil && other.CellGap == nil {
		return false
	}

	if resource.CellGap != nil {
		if *resource.CellGap != *other.CellGap {
			return false
		}
	}
	if resource.CellRadius == nil && other.CellRadius != nil || resource.CellRadius != nil && other.CellRadius == nil {
		return false
	}

	if resource.CellRadius != nil {
		if *resource.CellRadius != *other.CellRadius {
			return false
		}
	}
	if resource.CellValues == nil && other.CellValues != nil || resource.CellValues != nil && other.CellValues == nil {
		return false
	}

	if resource.CellValues != nil {
		if !resource.CellValues.Equals(*other.CellValues) {
			return false
		}
	}
	if !resource.YAxis.Equals(other.YAxis) {
		return false
	}
	if !resource.Legend.Equals(other.Legend) {
		return false
	}
	if !resource.Tooltip.Equals(other.Tooltip) {
		return false
	}
	if !resource.Exemplars.Equals(other.Exemplars) {
		return false
	}

	return true
}

type FieldConfig struct {
	ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
}

func (resource FieldConfig) Equals(other FieldConfig) bool {
	if resource.ScaleDistribution == nil && other.ScaleDistribution != nil || resource.ScaleDistribution != nil && other.ScaleDistribution == nil {
		return false
	}

	if resource.ScaleDistribution != nil {
		if !resource.ScaleDistribution.Equals(*other.ScaleDistribution) {
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

	return true
}

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "heatmap",
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
