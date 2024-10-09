// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package barchart

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
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

func (resource Options) Equals(other Options) bool {
	if resource.XField == nil && other.XField != nil || resource.XField != nil && other.XField == nil {
		return false
	}

	if resource.XField != nil {
		if *resource.XField != *other.XField {
			return false
		}
	}
	if resource.ColorByField == nil && other.ColorByField != nil || resource.ColorByField != nil && other.ColorByField == nil {
		return false
	}

	if resource.ColorByField != nil {
		if *resource.ColorByField != *other.ColorByField {
			return false
		}
	}
	if resource.Orientation != other.Orientation {
		return false
	}
	if resource.BarRadius == nil && other.BarRadius != nil || resource.BarRadius != nil && other.BarRadius == nil {
		return false
	}

	if resource.BarRadius != nil {
		if *resource.BarRadius != *other.BarRadius {
			return false
		}
	}
	if resource.XTickLabelRotation != other.XTickLabelRotation {
		return false
	}
	if resource.XTickLabelMaxLength != other.XTickLabelMaxLength {
		return false
	}
	if resource.XTickLabelSpacing == nil && other.XTickLabelSpacing != nil || resource.XTickLabelSpacing != nil && other.XTickLabelSpacing == nil {
		return false
	}

	if resource.XTickLabelSpacing != nil {
		if *resource.XTickLabelSpacing != *other.XTickLabelSpacing {
			return false
		}
	}
	if resource.Stacking != other.Stacking {
		return false
	}
	if resource.ShowValue != other.ShowValue {
		return false
	}
	if resource.BarWidth != other.BarWidth {
		return false
	}
	if resource.GroupWidth != other.GroupWidth {
		return false
	}
	if !resource.Legend.Equals(other.Legend) {
		return false
	}
	if !resource.Tooltip.Equals(other.Tooltip) {
		return false
	}
	if resource.Text == nil && other.Text != nil || resource.Text != nil && other.Text == nil {
		return false
	}

	if resource.Text != nil {
		if !resource.Text.Equals(*other.Text) {
			return false
		}
	}
	if resource.FullHighlight != other.FullHighlight {
		return false
	}

	return true
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
	AxisCenteredZero  *bool                           `json:"axisCenteredZero,omitempty"`
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
	// Threshold rendering
	ThresholdsStyle *common.GraphThresholdsStyleConfig `json:"thresholdsStyle,omitempty"`
	AxisBorderShow  *bool                              `json:"axisBorderShow,omitempty"`
}

func (resource FieldConfig) Equals(other FieldConfig) bool {
	if resource.LineWidth == nil && other.LineWidth != nil || resource.LineWidth != nil && other.LineWidth == nil {
		return false
	}

	if resource.LineWidth != nil {
		if *resource.LineWidth != *other.LineWidth {
			return false
		}
	}
	if resource.FillOpacity == nil && other.FillOpacity != nil || resource.FillOpacity != nil && other.FillOpacity == nil {
		return false
	}

	if resource.FillOpacity != nil {
		if *resource.FillOpacity != *other.FillOpacity {
			return false
		}
	}
	if resource.GradientMode == nil && other.GradientMode != nil || resource.GradientMode != nil && other.GradientMode == nil {
		return false
	}

	if resource.GradientMode != nil {
		if *resource.GradientMode != *other.GradientMode {
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
	if resource.HideFrom == nil && other.HideFrom != nil || resource.HideFrom != nil && other.HideFrom == nil {
		return false
	}

	if resource.HideFrom != nil {
		if !resource.HideFrom.Equals(*other.HideFrom) {
			return false
		}
	}
	if resource.ThresholdsStyle == nil && other.ThresholdsStyle != nil || resource.ThresholdsStyle != nil && other.ThresholdsStyle == nil {
		return false
	}

	if resource.ThresholdsStyle != nil {
		if !resource.ThresholdsStyle.Equals(*other.ThresholdsStyle) {
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

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "barchart",
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
