// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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

// Controls which axis to allow selection on
type HeatmapSelectionMode string

const (
	HeatmapSelectionModeX  HeatmapSelectionMode = "x"
	HeatmapSelectionModeY  HeatmapSelectionMode = "y"
	HeatmapSelectionModeXy HeatmapSelectionMode = "xy"
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

// Controls cell value options
type CellValues struct {
	// Controls the cell value unit
	Unit *string `json:"unit,omitempty"`
	// Controls the number of decimals for cell values
	Decimals *float32 `json:"decimals,omitempty"`
}

// Controls the value filter range
type FilterValueRange struct {
	// Sets the filter range to values less than or equal to the given value
	Le *float32 `json:"le,omitempty"`
	// Sets the filter range to values greater than or equal to the given value
	Ge *float32 `json:"ge,omitempty"`
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

// Controls legend options
type HeatmapLegend struct {
	// Controls if the legend is shown
	Show bool `json:"show"`
}

// Controls exemplar options
type ExemplarConfig struct {
	// Sets the color of the exemplar markers
	Color string `json:"color"`
}

// Controls frame rows options
type RowsHeatmapOptions struct {
	// Sets the name of the cell when not calculating from data
	Value *string `json:"value,omitempty"`
	// Controls tick alignment when not calculating from data
	Layout *common.HeatmapCellLayout `json:"layout,omitempty"`
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
	// Controls which axis to allow selection on
	SelectionMode *HeatmapSelectionMode `json:"selectionMode,omitempty"`
}

type FieldConfig struct {
	ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
}

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "heatmap",
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
