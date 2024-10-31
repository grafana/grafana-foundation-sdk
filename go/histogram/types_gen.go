// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package histogram

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
	// Size of each bucket
	BucketSize *int32 `json:"bucketSize,omitempty"`
	// Offset buckets by this amount
	BucketOffset *int32                   `json:"bucketOffset,omitempty"`
	Legend       common.VizLegendOptions  `json:"legend"`
	Tooltip      common.VizTooltipOptions `json:"tooltip"`
	// Combines multiple series into a single histogram
	Combine *bool `json:"combine,omitempty"`
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
	// Field "bucketSize"
	if fields["bucketSize"] != nil {
		if string(fields["bucketSize"]) != "null" {
			if err := json.Unmarshal(fields["bucketSize"], &resource.BucketSize); err != nil {
				errs = append(errs, cog.MakeBuildErrors("bucketSize", err)...)
			}

		}
		delete(fields, "bucketSize")

	}
	// Field "bucketOffset"
	if fields["bucketOffset"] != nil {
		if string(fields["bucketOffset"]) != "null" {
			if err := json.Unmarshal(fields["bucketOffset"], &resource.BucketOffset); err != nil {
				errs = append(errs, cog.MakeBuildErrors("bucketOffset", err)...)
			}

		}
		delete(fields, "bucketOffset")

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
	// Field "combine"
	if fields["combine"] != nil {
		if string(fields["combine"]) != "null" {
			if err := json.Unmarshal(fields["combine"], &resource.Combine); err != nil {
				errs = append(errs, cog.MakeBuildErrors("combine", err)...)
			}

		}
		delete(fields, "combine")

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
	if resource.BucketSize == nil && other.BucketSize != nil || resource.BucketSize != nil && other.BucketSize == nil {
		return false
	}

	if resource.BucketSize != nil {
		if *resource.BucketSize != *other.BucketSize {
			return false
		}
	}
	if resource.BucketOffset == nil && other.BucketOffset != nil || resource.BucketOffset != nil && other.BucketOffset == nil {
		return false
	}

	if resource.BucketOffset != nil {
		if *resource.BucketOffset != *other.BucketOffset {
			return false
		}
	}
	if !resource.Legend.Equals(other.Legend) {
		return false
	}
	if !resource.Tooltip.Equals(other.Tooltip) {
		return false
	}
	if resource.Combine == nil && other.Combine != nil || resource.Combine != nil && other.Combine == nil {
		return false
	}

	if resource.Combine != nil {
		if *resource.Combine != *other.Combine {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}
	if err := resource.Tooltip.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
	// Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
	// Gradient appearance is influenced by the Fill opacity setting.
	GradientMode     *common.GraphGradientMode `json:"gradientMode,omitempty"`
	AxisCenteredZero *bool                     `json:"axisCenteredZero,omitempty"`
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
	// Field "gradientMode"
	if fields["gradientMode"] != nil {
		if string(fields["gradientMode"]) != "null" {
			if err := json.Unmarshal(fields["gradientMode"], &resource.GradientMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("gradientMode", err)...)
			}

		}
		delete(fields, "gradientMode")

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
	if resource.HideFrom == nil && other.HideFrom != nil || resource.HideFrom != nil && other.HideFrom == nil {
		return false
	}

	if resource.HideFrom != nil {
		if !resource.HideFrom.Equals(*other.HideFrom) {
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
	if resource.AxisCenteredZero == nil && other.AxisCenteredZero != nil || resource.AxisCenteredZero != nil && other.AxisCenteredZero == nil {
		return false
	}

	if resource.AxisCenteredZero != nil {
		if *resource.AxisCenteredZero != *other.AxisCenteredZero {
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

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to histogram panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "histogram",
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
