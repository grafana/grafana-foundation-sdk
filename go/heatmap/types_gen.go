// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapColorOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *HeatmapColorOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}

		}
		delete(fields, "mode")

	}
	// Field "scheme"
	if fields["scheme"] != nil {
		if string(fields["scheme"]) != "null" {
			if err := json.Unmarshal(fields["scheme"], &resource.Scheme); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scheme", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("scheme", errors.New("required field is null"))...)

		}
		delete(fields, "scheme")
	} else {
		errs = append(errs, cog.MakeBuildErrors("scheme", errors.New("required field is missing from input"))...)
	}
	// Field "fill"
	if fields["fill"] != nil {
		if string(fields["fill"]) != "null" {
			if err := json.Unmarshal(fields["fill"], &resource.Fill); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fill", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("fill", errors.New("required field is null"))...)

		}
		delete(fields, "fill")
	} else {
		errs = append(errs, cog.MakeBuildErrors("fill", errors.New("required field is missing from input"))...)
	}
	// Field "scale"
	if fields["scale"] != nil {
		if string(fields["scale"]) != "null" {
			if err := json.Unmarshal(fields["scale"], &resource.Scale); err != nil {
				errs = append(errs, cog.MakeBuildErrors("scale", err)...)
			}

		}
		delete(fields, "scale")

	}
	// Field "exponent"
	if fields["exponent"] != nil {
		if string(fields["exponent"]) != "null" {
			if err := json.Unmarshal(fields["exponent"], &resource.Exponent); err != nil {
				errs = append(errs, cog.MakeBuildErrors("exponent", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("exponent", errors.New("required field is null"))...)

		}
		delete(fields, "exponent")
	} else {
		errs = append(errs, cog.MakeBuildErrors("exponent", errors.New("required field is missing from input"))...)
	}
	// Field "steps"
	if fields["steps"] != nil {
		if string(fields["steps"]) != "null" {
			if err := json.Unmarshal(fields["steps"], &resource.Steps); err != nil {
				errs = append(errs, cog.MakeBuildErrors("steps", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("steps", errors.New("required field is null"))...)

		}
		delete(fields, "steps")
	} else {
		errs = append(errs, cog.MakeBuildErrors("steps", errors.New("required field is missing from input"))...)
	}
	// Field "reverse"
	if fields["reverse"] != nil {
		if string(fields["reverse"]) != "null" {
			if err := json.Unmarshal(fields["reverse"], &resource.Reverse); err != nil {
				errs = append(errs, cog.MakeBuildErrors("reverse", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("reverse", errors.New("required field is null"))...)

		}
		delete(fields, "reverse")
	} else {
		errs = append(errs, cog.MakeBuildErrors("reverse", errors.New("required field is missing from input"))...)
	}
	// Field "min"
	if fields["min"] != nil {
		if string(fields["min"]) != "null" {
			if err := json.Unmarshal(fields["min"], &resource.Min); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min", err)...)
			}

		}
		delete(fields, "min")

	}
	// Field "max"
	if fields["max"] != nil {
		if string(fields["max"]) != "null" {
			if err := json.Unmarshal(fields["max"], &resource.Max); err != nil {
				errs = append(errs, cog.MakeBuildErrors("max", err)...)
			}

		}
		delete(fields, "max")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("HeatmapColorOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `HeatmapColorOptions` objects.
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

// Validate checks all the validation constraints that may be defined on `HeatmapColorOptions` fields for violations and returns them.
func (resource HeatmapColorOptions) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Steps >= 2) {
		errs = append(errs, cog.MakeBuildErrors(
			"steps",
			errors.New("must be >= 2"),
		)...)
	}
	if !(resource.Steps <= 128) {
		errs = append(errs, cog.MakeBuildErrors(
			"steps",
			errors.New("must be <= 128"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `YAxisConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *YAxisConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "unit"
	if fields["unit"] != nil {
		if string(fields["unit"]) != "null" {
			if err := json.Unmarshal(fields["unit"], &resource.Unit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("unit", err)...)
			}

		}
		delete(fields, "unit")

	}
	// Field "reverse"
	if fields["reverse"] != nil {
		if string(fields["reverse"]) != "null" {
			if err := json.Unmarshal(fields["reverse"], &resource.Reverse); err != nil {
				errs = append(errs, cog.MakeBuildErrors("reverse", err)...)
			}

		}
		delete(fields, "reverse")

	}
	// Field "decimals"
	if fields["decimals"] != nil {
		if string(fields["decimals"]) != "null" {
			if err := json.Unmarshal(fields["decimals"], &resource.Decimals); err != nil {
				errs = append(errs, cog.MakeBuildErrors("decimals", err)...)
			}

		}
		delete(fields, "decimals")

	}
	// Field "min"
	if fields["min"] != nil {
		if string(fields["min"]) != "null" {
			if err := json.Unmarshal(fields["min"], &resource.Min); err != nil {
				errs = append(errs, cog.MakeBuildErrors("min", err)...)
			}

		}
		delete(fields, "min")

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
	// Field "max"
	if fields["max"] != nil {
		if string(fields["max"]) != "null" {
			if err := json.Unmarshal(fields["max"], &resource.Max); err != nil {
				errs = append(errs, cog.MakeBuildErrors("max", err)...)
			}

		}
		delete(fields, "max")

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
		errs = append(errs, cog.MakeBuildErrors("YAxisConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `YAxisConfig` objects.
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

// Validate checks all the validation constraints that may be defined on `YAxisConfig` fields for violations and returns them.
func (resource YAxisConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.ScaleDistribution != nil {
		if err := resource.ScaleDistribution.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Controls cell value options
type CellValues struct {
	// Controls the cell value unit
	Unit *string `json:"unit,omitempty"`
	// Controls the number of decimals for cell values
	Decimals *float32 `json:"decimals,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CellValues` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CellValues) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "unit"
	if fields["unit"] != nil {
		if string(fields["unit"]) != "null" {
			if err := json.Unmarshal(fields["unit"], &resource.Unit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("unit", err)...)
			}

		}
		delete(fields, "unit")

	}
	// Field "decimals"
	if fields["decimals"] != nil {
		if string(fields["decimals"]) != "null" {
			if err := json.Unmarshal(fields["decimals"], &resource.Decimals); err != nil {
				errs = append(errs, cog.MakeBuildErrors("decimals", err)...)
			}

		}
		delete(fields, "decimals")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CellValues", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CellValues` objects.
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

// Validate checks all the validation constraints that may be defined on `CellValues` fields for violations and returns them.
func (resource CellValues) Validate() error {
	return nil
}

// Controls the value filter range
type FilterValueRange struct {
	// Sets the filter range to values less than or equal to the given value
	Le *float32 `json:"le,omitempty"`
	// Sets the filter range to values greater than or equal to the given value
	Ge *float32 `json:"ge,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `FilterValueRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *FilterValueRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "le"
	if fields["le"] != nil {
		if string(fields["le"]) != "null" {
			if err := json.Unmarshal(fields["le"], &resource.Le); err != nil {
				errs = append(errs, cog.MakeBuildErrors("le", err)...)
			}

		}
		delete(fields, "le")

	}
	// Field "ge"
	if fields["ge"] != nil {
		if string(fields["ge"]) != "null" {
			if err := json.Unmarshal(fields["ge"], &resource.Ge); err != nil {
				errs = append(errs, cog.MakeBuildErrors("ge", err)...)
			}

		}
		delete(fields, "ge")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("FilterValueRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `FilterValueRange` objects.
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

// Validate checks all the validation constraints that may be defined on `FilterValueRange` fields for violations and returns them.
func (resource FilterValueRange) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapTooltip` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *HeatmapTooltip) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "maxHeight"
	if fields["maxHeight"] != nil {
		if string(fields["maxHeight"]) != "null" {
			if err := json.Unmarshal(fields["maxHeight"], &resource.MaxHeight); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxHeight", err)...)
			}

		}
		delete(fields, "maxHeight")

	}
	// Field "maxWidth"
	if fields["maxWidth"] != nil {
		if string(fields["maxWidth"]) != "null" {
			if err := json.Unmarshal(fields["maxWidth"], &resource.MaxWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxWidth", err)...)
			}

		}
		delete(fields, "maxWidth")

	}
	// Field "yHistogram"
	if fields["yHistogram"] != nil {
		if string(fields["yHistogram"]) != "null" {
			if err := json.Unmarshal(fields["yHistogram"], &resource.YHistogram); err != nil {
				errs = append(errs, cog.MakeBuildErrors("yHistogram", err)...)
			}

		}
		delete(fields, "yHistogram")

	}
	// Field "showColorScale"
	if fields["showColorScale"] != nil {
		if string(fields["showColorScale"]) != "null" {
			if err := json.Unmarshal(fields["showColorScale"], &resource.ShowColorScale); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showColorScale", err)...)
			}

		}
		delete(fields, "showColorScale")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("HeatmapTooltip", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `HeatmapTooltip` objects.
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

// Validate checks all the validation constraints that may be defined on `HeatmapTooltip` fields for violations and returns them.
func (resource HeatmapTooltip) Validate() error {
	return nil
}

// Controls legend options
type HeatmapLegend struct {
	// Controls if the legend is shown
	Show bool `json:"show"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `HeatmapLegend` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *HeatmapLegend) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "show"
	if fields["show"] != nil {
		if string(fields["show"]) != "null" {
			if err := json.Unmarshal(fields["show"], &resource.Show); err != nil {
				errs = append(errs, cog.MakeBuildErrors("show", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("show", errors.New("required field is null"))...)

		}
		delete(fields, "show")
	} else {
		errs = append(errs, cog.MakeBuildErrors("show", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("HeatmapLegend", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `HeatmapLegend` objects.
func (resource HeatmapLegend) Equals(other HeatmapLegend) bool {
	if resource.Show != other.Show {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `HeatmapLegend` fields for violations and returns them.
func (resource HeatmapLegend) Validate() error {
	return nil
}

// Controls exemplar options
type ExemplarConfig struct {
	// Sets the color of the exemplar markers
	Color string `json:"color"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExemplarConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExemplarConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "color"
	if fields["color"] != nil {
		if string(fields["color"]) != "null" {
			if err := json.Unmarshal(fields["color"], &resource.Color); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("color", errors.New("required field is null"))...)

		}
		delete(fields, "color")
	} else {
		errs = append(errs, cog.MakeBuildErrors("color", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExemplarConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExemplarConfig` objects.
func (resource ExemplarConfig) Equals(other ExemplarConfig) bool {
	if resource.Color != other.Color {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExemplarConfig` fields for violations and returns them.
func (resource ExemplarConfig) Validate() error {
	return nil
}

// Controls frame rows options
type RowsHeatmapOptions struct {
	// Sets the name of the cell when not calculating from data
	Value *string `json:"value,omitempty"`
	// Controls tick alignment when not calculating from data
	Layout *common.HeatmapCellLayout `json:"layout,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsHeatmapOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *RowsHeatmapOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}

		}
		delete(fields, "value")

	}
	// Field "layout"
	if fields["layout"] != nil {
		if string(fields["layout"]) != "null" {
			if err := json.Unmarshal(fields["layout"], &resource.Layout); err != nil {
				errs = append(errs, cog.MakeBuildErrors("layout", err)...)
			}

		}
		delete(fields, "layout")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("RowsHeatmapOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `RowsHeatmapOptions` objects.
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

// Validate checks all the validation constraints that may be defined on `RowsHeatmapOptions` fields for violations and returns them.
func (resource RowsHeatmapOptions) Validate() error {
	return nil
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
	// Field "calculate"
	if fields["calculate"] != nil {
		if string(fields["calculate"]) != "null" {
			if err := json.Unmarshal(fields["calculate"], &resource.Calculate); err != nil {
				errs = append(errs, cog.MakeBuildErrors("calculate", err)...)
			}

		}
		delete(fields, "calculate")

	}
	// Field "calculation"
	if fields["calculation"] != nil {
		if string(fields["calculation"]) != "null" {

			resource.Calculation = &common.HeatmapCalculationOptions{}
			if err := resource.Calculation.UnmarshalJSONStrict(fields["calculation"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("calculation", err)...)
			}

		}
		delete(fields, "calculation")

	}
	// Field "color"
	if fields["color"] != nil {
		if string(fields["color"]) != "null" {

			resource.Color = HeatmapColorOptions{}
			if err := resource.Color.UnmarshalJSONStrict(fields["color"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("color", errors.New("required field is null"))...)

		}
		delete(fields, "color")
	} else {
		errs = append(errs, cog.MakeBuildErrors("color", errors.New("required field is missing from input"))...)
	}
	// Field "filterValues"
	if fields["filterValues"] != nil {
		if string(fields["filterValues"]) != "null" {

			resource.FilterValues = &FilterValueRange{}
			if err := resource.FilterValues.UnmarshalJSONStrict(fields["filterValues"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filterValues", err)...)
			}

		}
		delete(fields, "filterValues")

	}
	// Field "rowsFrame"
	if fields["rowsFrame"] != nil {
		if string(fields["rowsFrame"]) != "null" {

			resource.RowsFrame = &RowsHeatmapOptions{}
			if err := resource.RowsFrame.UnmarshalJSONStrict(fields["rowsFrame"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("rowsFrame", err)...)
			}

		}
		delete(fields, "rowsFrame")

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
	// Field "cellGap"
	if fields["cellGap"] != nil {
		if string(fields["cellGap"]) != "null" {
			if err := json.Unmarshal(fields["cellGap"], &resource.CellGap); err != nil {
				errs = append(errs, cog.MakeBuildErrors("cellGap", err)...)
			}

		}
		delete(fields, "cellGap")

	}
	// Field "cellRadius"
	if fields["cellRadius"] != nil {
		if string(fields["cellRadius"]) != "null" {
			if err := json.Unmarshal(fields["cellRadius"], &resource.CellRadius); err != nil {
				errs = append(errs, cog.MakeBuildErrors("cellRadius", err)...)
			}

		}
		delete(fields, "cellRadius")

	}
	// Field "cellValues"
	if fields["cellValues"] != nil {
		if string(fields["cellValues"]) != "null" {

			resource.CellValues = &CellValues{}
			if err := resource.CellValues.UnmarshalJSONStrict(fields["cellValues"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("cellValues", err)...)
			}

		}
		delete(fields, "cellValues")

	}
	// Field "yAxis"
	if fields["yAxis"] != nil {
		if string(fields["yAxis"]) != "null" {

			resource.YAxis = YAxisConfig{}
			if err := resource.YAxis.UnmarshalJSONStrict(fields["yAxis"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("yAxis", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("yAxis", errors.New("required field is null"))...)

		}
		delete(fields, "yAxis")
	} else {
		errs = append(errs, cog.MakeBuildErrors("yAxis", errors.New("required field is missing from input"))...)
	}
	// Field "legend"
	if fields["legend"] != nil {
		if string(fields["legend"]) != "null" {

			resource.Legend = HeatmapLegend{}
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

			resource.Tooltip = HeatmapTooltip{}
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
	// Field "exemplars"
	if fields["exemplars"] != nil {
		if string(fields["exemplars"]) != "null" {

			resource.Exemplars = ExemplarConfig{}
			if err := resource.Exemplars.UnmarshalJSONStrict(fields["exemplars"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("exemplars", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("exemplars", errors.New("required field is null"))...)

		}
		delete(fields, "exemplars")
	} else {
		errs = append(errs, cog.MakeBuildErrors("exemplars", errors.New("required field is missing from input"))...)
	}
	// Field "selectionMode"
	if fields["selectionMode"] != nil {
		if string(fields["selectionMode"]) != "null" {
			if err := json.Unmarshal(fields["selectionMode"], &resource.SelectionMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("selectionMode", err)...)
			}

		}
		delete(fields, "selectionMode")

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
	if resource.SelectionMode == nil && other.SelectionMode != nil || resource.SelectionMode != nil && other.SelectionMode == nil {
		return false
	}

	if resource.SelectionMode != nil {
		if *resource.SelectionMode != *other.SelectionMode {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if resource.Calculation != nil {
		if err := resource.Calculation.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("calculation", err)...)
		}
	}
	if err := resource.Color.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("color", err)...)
	}
	if resource.FilterValues != nil {
		if err := resource.FilterValues.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("filterValues", err)...)
		}
	}
	if resource.RowsFrame != nil {
		if err := resource.RowsFrame.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("rowsFrame", err)...)
		}
	}
	if resource.CellGap != nil {
		if !(*resource.CellGap <= 25) {
			errs = append(errs, cog.MakeBuildErrors(
				"cellGap",
				errors.New("must be <= 25"),
			)...)
		}
	}
	if resource.CellValues != nil {
		if err := resource.CellValues.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("cellValues", err)...)
		}
	}
	if err := resource.YAxis.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("yAxis", err)...)
	}
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}
	if err := resource.Tooltip.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
	}
	if err := resource.Exemplars.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("exemplars", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type FieldConfig struct {
	ScaleDistribution *common.ScaleDistributionConfig `json:"scaleDistribution,omitempty"`
	HideFrom          *common.HideSeriesConfig        `json:"hideFrom,omitempty"`
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

// Validate checks all the validation constraints that may be defined on `FieldConfig` fields for violations and returns them.
func (resource FieldConfig) Validate() error {
	var errs cog.BuildErrors
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

// VariantConfig returns the configuration related to heatmap panels.
// This configuration describes how to unmarshal it, convert it to code, …
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
