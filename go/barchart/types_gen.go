// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package barchart

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	// Manually select which field from the dataset to represent the x field.
	XField *string `json:"xField,omitempty"`
	// Use the color value for a sibling field to color each bar value.
	ColorByField *string `json:"colorByField,omitempty"`
	// Controls the orientation of the bar chart, either vertical or horizontal.
	Orientation common.VizOrientation `json:"orientation"`
	// Controls the radius of each bar.
	BarRadius *float64 `json:"barRadius,omitempty"`
	// Controls the rotation of the x axis labels.
	XTickLabelRotation int32 `json:"xTickLabelRotation"`
	// Sets the max length that a label can have before it is truncated.
	XTickLabelMaxLength int32 `json:"xTickLabelMaxLength"`
	// Controls the spacing between x axis labels.
	// negative values indicate backwards skipping behavior
	XTickLabelSpacing *int32 `json:"xTickLabelSpacing,omitempty"`
	// Controls whether bars are stacked or not, either normally or in percent mode.
	Stacking common.StackingMode `json:"stacking"`
	// This controls whether values are shown on top or to the left of bars.
	ShowValue common.VisibilityMode `json:"showValue"`
	// Controls the width of bars. 1 = Max width, 0 = Min width.
	BarWidth float64 `json:"barWidth"`
	// Controls the width of groups. 1 = max with, 0 = min width.
	GroupWidth float64                       `json:"groupWidth"`
	Legend     common.VizLegendOptions       `json:"legend"`
	Tooltip    common.VizTooltipOptions      `json:"tooltip"`
	Text       *common.VizTextDisplayOptions `json:"text,omitempty"`
	// Enables mode which highlights the entire bar area and shows tooltip when cursor
	// hovers over highlighted area
	FullHighlight bool `json:"fullHighlight"`
}

type FieldConfig struct {
	// Controls line width of the bars.
	LineWidth *int32 `json:"lineWidth,omitempty"`
	// Controls the fill opacity of the bars.
	FillOpacity *int32 `json:"fillOpacity,omitempty"`
	// Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
	// Gradient appearance is influenced by the Fill opacity setting.
	GradientMode      *common.GraphGradientMode       `json:"gradientMode,omitempty"`
	AxisPlacement     *common.AxisPlacement           `json:"axisPlacement,omitempty"`
	AxisColorMode     *common.AxisColorMode           `json:"axisColorMode,omitempty"`
	AxisLabel         *string                         `json:"axisLabel,omitempty"`
	AxisWidth         *float64                        `json:"axisWidth,omitempty"`
	AxisSoftMin       *float64                        `json:"axisSoftMin,omitempty"`
	AxisSoftMax       *float64                        `json:"axisSoftMax,omitempty"`
	AxisGridShow      *bool                           `json:"axisGridShow,omitempty"`
	ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
	// Threshold rendering
	ThresholdsStyle  *common.GraphThresholdsStyleConfig `json:"thresholdsStyle,omitempty"`
	AxisCenteredZero *bool                              `json:"axisCenteredZero,omitempty"`
}
