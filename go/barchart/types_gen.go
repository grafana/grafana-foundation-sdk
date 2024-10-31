// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package barchart

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Options` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Options) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "xField"
	if fields["xField"] != nil {
		if string(fields["xField"]) != "null" {
			if err := json.Unmarshal(fields["xField"], &resource.XField); err != nil {
				errs = append(errs, cog.MakeBuildErrors("xField", err)...)
			}

		}
		delete(fields, "xField")

	}
	// Field "colorByField"
	if fields["colorByField"] != nil {
		if string(fields["colorByField"]) != "null" {
			if err := json.Unmarshal(fields["colorByField"], &resource.ColorByField); err != nil {
				errs = append(errs, cog.MakeBuildErrors("colorByField", err)...)
			}

		}
		delete(fields, "colorByField")

	}
	// Field "orientation"
	if fields["orientation"] != nil {
		if string(fields["orientation"]) != "null" {
			if err := json.Unmarshal(fields["orientation"], &resource.Orientation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orientation", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("orientation", errors.New("required field is null"))...)

		}
		delete(fields, "orientation")
	} else {
		errs = append(errs, cog.MakeBuildErrors("orientation", errors.New("required field is missing from input"))...)
	}
	// Field "barRadius"
	if fields["barRadius"] != nil {
		if string(fields["barRadius"]) != "null" {
			if err := json.Unmarshal(fields["barRadius"], &resource.BarRadius); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barRadius", err)...)
			}

		}
		delete(fields, "barRadius")

	}
	// Field "xTickLabelRotation"
	if fields["xTickLabelRotation"] != nil {
		if string(fields["xTickLabelRotation"]) != "null" {
			if err := json.Unmarshal(fields["xTickLabelRotation"], &resource.XTickLabelRotation); err != nil {
				errs = append(errs, cog.MakeBuildErrors("xTickLabelRotation", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("xTickLabelRotation", errors.New("required field is null"))...)

		}
		delete(fields, "xTickLabelRotation")
	} else {
		errs = append(errs, cog.MakeBuildErrors("xTickLabelRotation", errors.New("required field is missing from input"))...)
	}
	// Field "xTickLabelMaxLength"
	if fields["xTickLabelMaxLength"] != nil {
		if string(fields["xTickLabelMaxLength"]) != "null" {
			if err := json.Unmarshal(fields["xTickLabelMaxLength"], &resource.XTickLabelMaxLength); err != nil {
				errs = append(errs, cog.MakeBuildErrors("xTickLabelMaxLength", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("xTickLabelMaxLength", errors.New("required field is null"))...)

		}
		delete(fields, "xTickLabelMaxLength")
	} else {
		errs = append(errs, cog.MakeBuildErrors("xTickLabelMaxLength", errors.New("required field is missing from input"))...)
	}
	// Field "xTickLabelSpacing"
	if fields["xTickLabelSpacing"] != nil {
		if string(fields["xTickLabelSpacing"]) != "null" {
			if err := json.Unmarshal(fields["xTickLabelSpacing"], &resource.XTickLabelSpacing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("xTickLabelSpacing", err)...)
			}

		}
		delete(fields, "xTickLabelSpacing")

	}
	// Field "stacking"
	if fields["stacking"] != nil {
		if string(fields["stacking"]) != "null" {
			if err := json.Unmarshal(fields["stacking"], &resource.Stacking); err != nil {
				errs = append(errs, cog.MakeBuildErrors("stacking", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("stacking", errors.New("required field is null"))...)

		}
		delete(fields, "stacking")
	} else {
		errs = append(errs, cog.MakeBuildErrors("stacking", errors.New("required field is missing from input"))...)
	}
	// Field "showValue"
	if fields["showValue"] != nil {
		if string(fields["showValue"]) != "null" {
			if err := json.Unmarshal(fields["showValue"], &resource.ShowValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showValue", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showValue", errors.New("required field is null"))...)

		}
		delete(fields, "showValue")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showValue", errors.New("required field is missing from input"))...)
	}
	// Field "barWidth"
	if fields["barWidth"] != nil {
		if string(fields["barWidth"]) != "null" {
			if err := json.Unmarshal(fields["barWidth"], &resource.BarWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("barWidth", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("barWidth", errors.New("required field is null"))...)

		}
		delete(fields, "barWidth")
	} else {
		errs = append(errs, cog.MakeBuildErrors("barWidth", errors.New("required field is missing from input"))...)
	}
	// Field "groupWidth"
	if fields["groupWidth"] != nil {
		if string(fields["groupWidth"]) != "null" {
			if err := json.Unmarshal(fields["groupWidth"], &resource.GroupWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupWidth", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("groupWidth", errors.New("required field is null"))...)

		}
		delete(fields, "groupWidth")
	} else {
		errs = append(errs, cog.MakeBuildErrors("groupWidth", errors.New("required field is missing from input"))...)
	}
	// Field "legend"
	if fields["legend"] != nil {
		if string(fields["legend"]) != "null" {

			resource.Legend = common.VizLegendOptions{}
			if err := resource.Legend.UnmarshalJSONStrict(fields["legend"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("legend", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is null"))...)

		}
		delete(fields, "legend")
	} else {
		errs = append(errs, cog.MakeBuildErrors("legend", errors.New("required field is missing from input"))...)
	}
	// Field "tooltip"
	if fields["tooltip"] != nil {
		if string(fields["tooltip"]) != "null" {

			resource.Tooltip = common.VizTooltipOptions{}
			if err := resource.Tooltip.UnmarshalJSONStrict(fields["tooltip"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is null"))...)

		}
		delete(fields, "tooltip")
	} else {
		errs = append(errs, cog.MakeBuildErrors("tooltip", errors.New("required field is missing from input"))...)
	}
	// Field "text"
	if fields["text"] != nil {
		if string(fields["text"]) != "null" {

			resource.Text = &common.VizTextDisplayOptions{}
			if err := resource.Text.UnmarshalJSONStrict(fields["text"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}

		}
		delete(fields, "text")

	}
	// Field "fullHighlight"
	if fields["fullHighlight"] != nil {
		if string(fields["fullHighlight"]) != "null" {
			if err := json.Unmarshal(fields["fullHighlight"], &resource.FullHighlight); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fullHighlight", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("fullHighlight", errors.New("required field is null"))...)

		}
		delete(fields, "fullHighlight")
	} else {
		errs = append(errs, cog.MakeBuildErrors("fullHighlight", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Options", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Options` objects.
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

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if resource.BarRadius != nil {
		if !(*resource.BarRadius >= 0) {
			errs = append(errs, cog.MakeBuildErrors(
				"barRadius",
				errors.New("must be >= 0"),
			)...)
		}
		if !(*resource.BarRadius <= 0.5) {
			errs = append(errs, cog.MakeBuildErrors(
				"barRadius",
				errors.New("must be <= 0.5"),
			)...)
		}
	}
	if !(resource.XTickLabelRotation >= -90) {
		errs = append(errs, cog.MakeBuildErrors(
			"xTickLabelRotation",
			errors.New("must be >= -90"),
		)...)
	}
	if !(resource.XTickLabelRotation <= 90) {
		errs = append(errs, cog.MakeBuildErrors(
			"xTickLabelRotation",
			errors.New("must be <= 90"),
		)...)
	}
	if !(resource.XTickLabelMaxLength >= 0) {
		errs = append(errs, cog.MakeBuildErrors(
			"xTickLabelMaxLength",
			errors.New("must be >= 0"),
		)...)
	}
	if !(resource.BarWidth >= 0) {
		errs = append(errs, cog.MakeBuildErrors(
			"barWidth",
			errors.New("must be >= 0"),
		)...)
	}
	if !(resource.BarWidth <= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"barWidth",
			errors.New("must be <= 1"),
		)...)
	}
	if !(resource.GroupWidth >= 0) {
		errs = append(errs, cog.MakeBuildErrors(
			"groupWidth",
			errors.New("must be >= 0"),
		)...)
	}
	if !(resource.GroupWidth <= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"groupWidth",
			errors.New("must be <= 1"),
		)...)
	}
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}
	if err := resource.Tooltip.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
	}
	if resource.Text != nil {
		if err := resource.Text.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("text", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FieldConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FieldConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "lineWidth"
	if fields["lineWidth"] != nil {
		if string(fields["lineWidth"]) != "null" {
			if err := json.Unmarshal(fields["lineWidth"], &resource.LineWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineWidth", err)...)
			}

		}
		delete(fields, "lineWidth")

	}
	// Field "fillOpacity"
	if fields["fillOpacity"] != nil {
		if string(fields["fillOpacity"]) != "null" {
			if err := json.Unmarshal(fields["fillOpacity"], &resource.FillOpacity); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fillOpacity", err)...)
			}

		}
		delete(fields, "fillOpacity")

	}
	// Field "gradientMode"
	if fields["gradientMode"] != nil {
		if string(fields["gradientMode"]) != "null" {
			if err := json.Unmarshal(fields["gradientMode"], &resource.GradientMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("gradientMode", err)...)
			}

		}
		delete(fields, "gradientMode")

	}
	// Field "axisPlacement"
	if fields["axisPlacement"] != nil {
		if string(fields["axisPlacement"]) != "null" {
			if err := json.Unmarshal(fields["axisPlacement"], &resource.AxisPlacement); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisPlacement", err)...)
			}

		}
		delete(fields, "axisPlacement")

	}
	// Field "axisColorMode"
	if fields["axisColorMode"] != nil {
		if string(fields["axisColorMode"]) != "null" {
			if err := json.Unmarshal(fields["axisColorMode"], &resource.AxisColorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisColorMode", err)...)
			}

		}
		delete(fields, "axisColorMode")

	}
	// Field "axisLabel"
	if fields["axisLabel"] != nil {
		if string(fields["axisLabel"]) != "null" {
			if err := json.Unmarshal(fields["axisLabel"], &resource.AxisLabel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisLabel", err)...)
			}

		}
		delete(fields, "axisLabel")

	}
	// Field "axisWidth"
	if fields["axisWidth"] != nil {
		if string(fields["axisWidth"]) != "null" {
			if err := json.Unmarshal(fields["axisWidth"], &resource.AxisWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisWidth", err)...)
			}

		}
		delete(fields, "axisWidth")

	}
	// Field "axisSoftMin"
	if fields["axisSoftMin"] != nil {
		if string(fields["axisSoftMin"]) != "null" {
			if err := json.Unmarshal(fields["axisSoftMin"], &resource.AxisSoftMin); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisSoftMin", err)...)
			}

		}
		delete(fields, "axisSoftMin")

	}
	// Field "axisSoftMax"
	if fields["axisSoftMax"] != nil {
		if string(fields["axisSoftMax"]) != "null" {
			if err := json.Unmarshal(fields["axisSoftMax"], &resource.AxisSoftMax); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisSoftMax", err)...)
			}

		}
		delete(fields, "axisSoftMax")

	}
	// Field "axisGridShow"
	if fields["axisGridShow"] != nil {
		if string(fields["axisGridShow"]) != "null" {
			if err := json.Unmarshal(fields["axisGridShow"], &resource.AxisGridShow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisGridShow", err)...)
			}

		}
		delete(fields, "axisGridShow")

	}
	// Field "scaleDistribution"
	if fields["scaleDistribution"] != nil {
		if string(fields["scaleDistribution"]) != "null" {

			resource.ScaleDistribution = &common.ScaleDistributionConfig{}
			if err := resource.ScaleDistribution.UnmarshalJSONStrict(fields["scaleDistribution"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
			}

		}
		delete(fields, "scaleDistribution")

	}
	// Field "axisCenteredZero"
	if fields["axisCenteredZero"] != nil {
		if string(fields["axisCenteredZero"]) != "null" {
			if err := json.Unmarshal(fields["axisCenteredZero"], &resource.AxisCenteredZero); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisCenteredZero", err)...)
			}

		}
		delete(fields, "axisCenteredZero")

	}
	// Field "hideFrom"
	if fields["hideFrom"] != nil {
		if string(fields["hideFrom"]) != "null" {

			resource.HideFrom = &common.HideSeriesConfig{}
			if err := resource.HideFrom.UnmarshalJSONStrict(fields["hideFrom"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
			}

		}
		delete(fields, "hideFrom")

	}
	// Field "thresholdsStyle"
	if fields["thresholdsStyle"] != nil {
		if string(fields["thresholdsStyle"]) != "null" {

			resource.ThresholdsStyle = &common.GraphThresholdsStyleConfig{}
			if err := resource.ThresholdsStyle.UnmarshalJSONStrict(fields["thresholdsStyle"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("thresholdsStyle", err)...)
			}

		}
		delete(fields, "thresholdsStyle")

	}
	// Field "axisBorderShow"
	if fields["axisBorderShow"] != nil {
		if string(fields["axisBorderShow"]) != "null" {
			if err := json.Unmarshal(fields["axisBorderShow"], &resource.AxisBorderShow); err != nil {
				errs = append(errs, cog.MakeBuildErrors("axisBorderShow", err)...)
			}

		}
		delete(fields, "axisBorderShow")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FieldConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FieldConfig` objects.
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

// Validate checks all the validation constraints that may be defined on `FieldConfig` fields for violations and returns them.
func (resource FieldConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.LineWidth != nil {
		if !(*resource.LineWidth <= 10) {
			errs = append(errs, cog.MakeBuildErrors(
				"lineWidth",
				errors.New("must be <= 10"),
			)...)
		}
	}
	if resource.FillOpacity != nil {
		if !(*resource.FillOpacity <= 100) {
			errs = append(errs, cog.MakeBuildErrors(
				"fillOpacity",
				errors.New("must be <= 100"),
			)...)
		}
	}
	if resource.ScaleDistribution != nil {
		if err := resource.ScaleDistribution.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
		}
	}
	if resource.HideFrom != nil {
		if err := resource.HideFrom.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
		}
	}
	if resource.ThresholdsStyle != nil {
		if err := resource.ThresholdsStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("thresholdsStyle", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to barchart panels.
// This configuration describes how to unmarshal it, convert it to code, …
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
		StrictOptionsUnmarshaler: func(raw []byte) (any, error) {
			options := &Options{}

			if err := options.UnmarshalJSONStrict(raw); err != nil {
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
		StrictFieldConfigUnmarshaler: func(raw []byte) (any, error) {
			fieldConfig := &FieldConfig{}

			if err := fieldConfig.UnmarshalJSONStrict(raw); err != nil {
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
