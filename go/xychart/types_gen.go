// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

type PointShape string

const (
	PointShapeCircle PointShape = "circle"
	PointShapeSquare PointShape = "square"
)

type SeriesMapping string

const (
	SeriesMappingAuto   SeriesMapping = "auto"
	SeriesMappingManual SeriesMapping = "manual"
)

type XYShowMode string

const (
	XYShowModePoints         XYShowMode = "points"
	XYShowModeLines          XYShowMode = "lines"
	XYShowModePointsAndLines XYShowMode = "points+lines"
)

// NOTE: (copied from dashboard_kind.cue, since not exported)
// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
type MatcherConfig struct {
	// The matcher id. This is used to find the matcher implementation from registry.
	Id string `json:"id"`
	// The matcher options. This is specific to the matcher implementation.
	Options any `json:"options,omitempty"`
}

// NewMatcherConfig creates a new MatcherConfig object.
func NewMatcherConfig() *MatcherConfig {
	return &MatcherConfig{
		Id: "",
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MatcherConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MatcherConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "id"
	if fields["id"] != nil {
		if string(fields["id"]) != "null" {
			if err := json.Unmarshal(fields["id"], &resource.Id); err != nil {
				errs = append(errs, cog.MakeBuildErrors("id", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is null"))...)

		}
		delete(fields, "id")

	}
	// Field "options"
	if fields["options"] != nil {
		if string(fields["options"]) != "null" {
			if err := json.Unmarshal(fields["options"], &resource.Options); err != nil {
				errs = append(errs, cog.MakeBuildErrors("options", err)...)
			}

		}
		delete(fields, "options")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MatcherConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MatcherConfig` objects.
func (resource MatcherConfig) Equals(other MatcherConfig) bool {
	if resource.Id != other.Id {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Options, other.Options) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MatcherConfig` fields for violations and returns them.
func (resource MatcherConfig) Validate() error {
	return nil
}

type FieldConfig struct {
	Show              *XYShowMode                     `json:"show,omitempty"`
	PointSize         *XychartFieldConfigPointSize    `json:"pointSize,omitempty"`
	PointShape        *PointShape                     `json:"pointShape,omitempty"`
	PointStrokeWidth  *int32                          `json:"pointStrokeWidth,omitempty"`
	FillOpacity       *uint32                         `json:"fillOpacity,omitempty"`
	LineWidth         *int32                          `json:"lineWidth,omitempty"`
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
	LineStyle         *common.LineStyle               `json:"lineStyle,omitempty"`
	AxisBorderShow    *bool                           `json:"axisBorderShow,omitempty"`
}

// NewFieldConfig creates a new FieldConfig object.
func NewFieldConfig() *FieldConfig {
	return &FieldConfig{
		Show:        (func(input XYShowMode) *XYShowMode { return &input })(XYShowModePoints),
		FillOpacity: (func(input uint32) *uint32 { return &input })(50),
	}
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
	// Field "show"
	if fields["show"] != nil {
		if string(fields["show"]) != "null" {
			if err := json.Unmarshal(fields["show"], &resource.Show); err != nil {
				errs = append(errs, cog.MakeBuildErrors("show", err)...)
			}

		}
		delete(fields, "show")

	}
	// Field "pointSize"
	if fields["pointSize"] != nil {
		if string(fields["pointSize"]) != "null" {

			resource.PointSize = &XychartFieldConfigPointSize{}
			if err := resource.PointSize.UnmarshalJSONStrict(fields["pointSize"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
			}

		}
		delete(fields, "pointSize")

	}
	// Field "pointShape"
	if fields["pointShape"] != nil {
		if string(fields["pointShape"]) != "null" {
			if err := json.Unmarshal(fields["pointShape"], &resource.PointShape); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointShape", err)...)
			}

		}
		delete(fields, "pointShape")

	}
	// Field "pointStrokeWidth"
	if fields["pointStrokeWidth"] != nil {
		if string(fields["pointStrokeWidth"]) != "null" {
			if err := json.Unmarshal(fields["pointStrokeWidth"], &resource.PointStrokeWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointStrokeWidth", err)...)
			}

		}
		delete(fields, "pointStrokeWidth")

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
	// Field "lineWidth"
	if fields["lineWidth"] != nil {
		if string(fields["lineWidth"]) != "null" {
			if err := json.Unmarshal(fields["lineWidth"], &resource.LineWidth); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineWidth", err)...)
			}

		}
		delete(fields, "lineWidth")

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
	// Field "lineStyle"
	if fields["lineStyle"] != nil {
		if string(fields["lineStyle"]) != "null" {

			resource.LineStyle = &common.LineStyle{}
			if err := resource.LineStyle.UnmarshalJSONStrict(fields["lineStyle"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
			}

		}
		delete(fields, "lineStyle")

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
	if resource.PointShape == nil && other.PointShape != nil || resource.PointShape != nil && other.PointShape == nil {
		return false
	}

	if resource.PointShape != nil {
		if *resource.PointShape != *other.PointShape {
			return false
		}
	}
	if resource.PointStrokeWidth == nil && other.PointStrokeWidth != nil || resource.PointStrokeWidth != nil && other.PointStrokeWidth == nil {
		return false
	}

	if resource.PointStrokeWidth != nil {
		if *resource.PointStrokeWidth != *other.PointStrokeWidth {
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
	if resource.LineWidth == nil && other.LineWidth != nil || resource.LineWidth != nil && other.LineWidth == nil {
		return false
	}

	if resource.LineWidth != nil {
		if *resource.LineWidth != *other.LineWidth {
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
	if resource.LineStyle == nil && other.LineStyle != nil || resource.LineStyle != nil && other.LineStyle == nil {
		return false
	}

	if resource.LineStyle != nil {
		if !resource.LineStyle.Equals(*other.LineStyle) {
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
	if resource.PointSize != nil {
		if err := resource.PointSize.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
		}
	}
	if resource.PointStrokeWidth != nil {
		if !(*resource.PointStrokeWidth >= 0) {
			errs = append(errs, cog.MakeBuildErrors(
				"pointStrokeWidth",
				errors.New("must be >= 0"),
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
	if resource.LineWidth != nil {
		if !(*resource.LineWidth >= 0) {
			errs = append(errs, cog.MakeBuildErrors(
				"lineWidth",
				errors.New("must be >= 0"),
			)...)
		}
	}
	if resource.HideFrom != nil {
		if err := resource.HideFrom.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("hideFrom", err)...)
		}
	}
	if resource.ScaleDistribution != nil {
		if err := resource.ScaleDistribution.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("scaleDistribution", err)...)
		}
	}
	if resource.LineStyle != nil {
		if err := resource.LineStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type XYSeriesConfig struct {
	Name  *XychartXYSeriesConfigName  `json:"name,omitempty"`
	Frame *XychartXYSeriesConfigFrame `json:"frame,omitempty"`
	X     *XychartXYSeriesConfigX     `json:"x,omitempty"`
	Y     *XychartXYSeriesConfigY     `json:"y,omitempty"`
	Color *XychartXYSeriesConfigColor `json:"color,omitempty"`
	Size  *XychartXYSeriesConfigSize  `json:"size,omitempty"`
}

// NewXYSeriesConfig creates a new XYSeriesConfig object.
func NewXYSeriesConfig() *XYSeriesConfig {
	return &XYSeriesConfig{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XYSeriesConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XYSeriesConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {

			resource.Name = &XychartXYSeriesConfigName{}
			if err := resource.Name.UnmarshalJSONStrict(fields["name"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "frame"
	if fields["frame"] != nil {
		if string(fields["frame"]) != "null" {

			resource.Frame = &XychartXYSeriesConfigFrame{}
			if err := resource.Frame.UnmarshalJSONStrict(fields["frame"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("frame", err)...)
			}

		}
		delete(fields, "frame")

	}
	// Field "x"
	if fields["x"] != nil {
		if string(fields["x"]) != "null" {

			resource.X = &XychartXYSeriesConfigX{}
			if err := resource.X.UnmarshalJSONStrict(fields["x"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("x", err)...)
			}

		}
		delete(fields, "x")

	}
	// Field "y"
	if fields["y"] != nil {
		if string(fields["y"]) != "null" {

			resource.Y = &XychartXYSeriesConfigY{}
			if err := resource.Y.UnmarshalJSONStrict(fields["y"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("y", err)...)
			}

		}
		delete(fields, "y")

	}
	// Field "color"
	if fields["color"] != nil {
		if string(fields["color"]) != "null" {

			resource.Color = &XychartXYSeriesConfigColor{}
			if err := resource.Color.UnmarshalJSONStrict(fields["color"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}

		}
		delete(fields, "color")

	}
	// Field "size"
	if fields["size"] != nil {
		if string(fields["size"]) != "null" {

			resource.Size = &XychartXYSeriesConfigSize{}
			if err := resource.Size.UnmarshalJSONStrict(fields["size"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("size", err)...)
			}

		}
		delete(fields, "size")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("XYSeriesConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XYSeriesConfig` objects.
func (resource XYSeriesConfig) Equals(other XYSeriesConfig) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if !resource.Name.Equals(*other.Name) {
			return false
		}
	}
	if resource.Frame == nil && other.Frame != nil || resource.Frame != nil && other.Frame == nil {
		return false
	}

	if resource.Frame != nil {
		if !resource.Frame.Equals(*other.Frame) {
			return false
		}
	}
	if resource.X == nil && other.X != nil || resource.X != nil && other.X == nil {
		return false
	}

	if resource.X != nil {
		if !resource.X.Equals(*other.X) {
			return false
		}
	}
	if resource.Y == nil && other.Y != nil || resource.Y != nil && other.Y == nil {
		return false
	}

	if resource.Y != nil {
		if !resource.Y.Equals(*other.Y) {
			return false
		}
	}
	if resource.Color == nil && other.Color != nil || resource.Color != nil && other.Color == nil {
		return false
	}

	if resource.Color != nil {
		if !resource.Color.Equals(*other.Color) {
			return false
		}
	}
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if !resource.Size.Equals(*other.Size) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `XYSeriesConfig` fields for violations and returns them.
func (resource XYSeriesConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.Name != nil {
		if err := resource.Name.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("name", err)...)
		}
	}
	if resource.Frame != nil {
		if err := resource.Frame.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("frame", err)...)
		}
	}
	if resource.X != nil {
		if err := resource.X.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("x", err)...)
		}
	}
	if resource.Y != nil {
		if err := resource.Y.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("y", err)...)
		}
	}
	if resource.Color != nil {
		if err := resource.Color.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("color", err)...)
		}
	}
	if resource.Size != nil {
		if err := resource.Size.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("size", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Options struct {
	Mapping SeriesMapping            `json:"mapping"`
	Legend  common.VizLegendOptions  `json:"legend"`
	Tooltip common.VizTooltipOptions `json:"tooltip"`
	Series  []XYSeriesConfig         `json:"series"`
}

// NewOptions creates a new Options object.
func NewOptions() *Options {
	return &Options{
		Legend:  *common.NewVizLegendOptions(),
		Tooltip: *common.NewVizTooltipOptions(),
		Series:  []XYSeriesConfig{},
	}
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
	// Field "mapping"
	if fields["mapping"] != nil {
		if string(fields["mapping"]) != "null" {
			if err := json.Unmarshal(fields["mapping"], &resource.Mapping); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mapping", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mapping", errors.New("required field is null"))...)

		}
		delete(fields, "mapping")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mapping", errors.New("required field is missing from input"))...)
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
	// Field "series"
	if fields["series"] != nil {
		if string(fields["series"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["series"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 XYSeriesConfig

				result1 = XYSeriesConfig{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("series["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Series = append(resource.Series, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("series", errors.New("required field is null"))...)

		}
		delete(fields, "series")
	} else {
		errs = append(errs, cog.MakeBuildErrors("series", errors.New("required field is missing from input"))...)
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
	if resource.Mapping != other.Mapping {
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

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Legend.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("legend", err)...)
	}
	if err := resource.Tooltip.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
	}

	for i1 := range resource.Series {
		if err := resource.Series[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("series["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type XychartFieldConfigPointSize struct {
	Fixed *int32 `json:"fixed,omitempty"`
	Min   *int32 `json:"min,omitempty"`
	Max   *int32 `json:"max,omitempty"`
}

// NewXychartFieldConfigPointSize creates a new XychartFieldConfigPointSize object.
func NewXychartFieldConfigPointSize() *XychartFieldConfigPointSize {
	return &XychartFieldConfigPointSize{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartFieldConfigPointSize` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XychartFieldConfigPointSize) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "fixed"
	if fields["fixed"] != nil {
		if string(fields["fixed"]) != "null" {
			if err := json.Unmarshal(fields["fixed"], &resource.Fixed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fixed", err)...)
			}

		}
		delete(fields, "fixed")

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
		errs = append(errs, cog.MakeBuildErrors("XychartFieldConfigPointSize", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XychartFieldConfigPointSize` objects.
func (resource XychartFieldConfigPointSize) Equals(other XychartFieldConfigPointSize) bool {
	if resource.Fixed == nil && other.Fixed != nil || resource.Fixed != nil && other.Fixed == nil {
		return false
	}

	if resource.Fixed != nil {
		if *resource.Fixed != *other.Fixed {
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

// Validate checks all the validation constraints that may be defined on `XychartFieldConfigPointSize` fields for violations and returns them.
func (resource XychartFieldConfigPointSize) Validate() error {
	var errs cog.BuildErrors
	if resource.Fixed != nil {
		if !(*resource.Fixed >= 0) {
			errs = append(errs, cog.MakeBuildErrors(
				"fixed",
				errors.New("must be >= 0"),
			)...)
		}
	}
	if resource.Min != nil {
		if !(*resource.Min >= 0) {
			errs = append(errs, cog.MakeBuildErrors(
				"min",
				errors.New("must be >= 0"),
			)...)
		}
	}
	if resource.Max != nil {
		if !(*resource.Max >= 0) {
			errs = append(errs, cog.MakeBuildErrors(
				"max",
				errors.New("must be >= 0"),
			)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type XychartXYSeriesConfigName struct {
	Fixed *string `json:"fixed,omitempty"`
}

// NewXychartXYSeriesConfigName creates a new XychartXYSeriesConfigName object.
func NewXychartXYSeriesConfigName() *XychartXYSeriesConfigName {
	return &XychartXYSeriesConfigName{}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigName` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XychartXYSeriesConfigName) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "fixed"
	if fields["fixed"] != nil {
		if string(fields["fixed"]) != "null" {
			if err := json.Unmarshal(fields["fixed"], &resource.Fixed); err != nil {
				errs = append(errs, cog.MakeBuildErrors("fixed", err)...)
			}

		}
		delete(fields, "fixed")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("XychartXYSeriesConfigName", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XychartXYSeriesConfigName` objects.
func (resource XychartXYSeriesConfigName) Equals(other XychartXYSeriesConfigName) bool {
	if resource.Fixed == nil && other.Fixed != nil || resource.Fixed != nil && other.Fixed == nil {
		return false
	}

	if resource.Fixed != nil {
		if *resource.Fixed != *other.Fixed {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigName` fields for violations and returns them.
func (resource XychartXYSeriesConfigName) Validate() error {
	return nil
}

type XychartXYSeriesConfigFrame struct {
	Matcher MatcherConfig `json:"matcher"`
}

// NewXychartXYSeriesConfigFrame creates a new XychartXYSeriesConfigFrame object.
func NewXychartXYSeriesConfigFrame() *XychartXYSeriesConfigFrame {
	return &XychartXYSeriesConfigFrame{
		Matcher: *NewMatcherConfig(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigFrame` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XychartXYSeriesConfigFrame) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "matcher"
	if fields["matcher"] != nil {
		if string(fields["matcher"]) != "null" {

			resource.Matcher = MatcherConfig{}
			if err := resource.Matcher.UnmarshalJSONStrict(fields["matcher"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is null"))...)

		}
		delete(fields, "matcher")
	} else {
		errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("XychartXYSeriesConfigFrame", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XychartXYSeriesConfigFrame` objects.
func (resource XychartXYSeriesConfigFrame) Equals(other XychartXYSeriesConfigFrame) bool {
	if !resource.Matcher.Equals(other.Matcher) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigFrame` fields for violations and returns them.
func (resource XychartXYSeriesConfigFrame) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Matcher.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type XychartXYSeriesConfigX struct {
	Matcher MatcherConfig `json:"matcher"`
}

// NewXychartXYSeriesConfigX creates a new XychartXYSeriesConfigX object.
func NewXychartXYSeriesConfigX() *XychartXYSeriesConfigX {
	return &XychartXYSeriesConfigX{
		Matcher: *NewMatcherConfig(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigX` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XychartXYSeriesConfigX) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "matcher"
	if fields["matcher"] != nil {
		if string(fields["matcher"]) != "null" {

			resource.Matcher = MatcherConfig{}
			if err := resource.Matcher.UnmarshalJSONStrict(fields["matcher"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is null"))...)

		}
		delete(fields, "matcher")
	} else {
		errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("XychartXYSeriesConfigX", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XychartXYSeriesConfigX` objects.
func (resource XychartXYSeriesConfigX) Equals(other XychartXYSeriesConfigX) bool {
	if !resource.Matcher.Equals(other.Matcher) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigX` fields for violations and returns them.
func (resource XychartXYSeriesConfigX) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Matcher.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type XychartXYSeriesConfigY struct {
	Matcher MatcherConfig `json:"matcher"`
}

// NewXychartXYSeriesConfigY creates a new XychartXYSeriesConfigY object.
func NewXychartXYSeriesConfigY() *XychartXYSeriesConfigY {
	return &XychartXYSeriesConfigY{
		Matcher: *NewMatcherConfig(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigY` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XychartXYSeriesConfigY) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "matcher"
	if fields["matcher"] != nil {
		if string(fields["matcher"]) != "null" {

			resource.Matcher = MatcherConfig{}
			if err := resource.Matcher.UnmarshalJSONStrict(fields["matcher"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is null"))...)

		}
		delete(fields, "matcher")
	} else {
		errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("XychartXYSeriesConfigY", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XychartXYSeriesConfigY` objects.
func (resource XychartXYSeriesConfigY) Equals(other XychartXYSeriesConfigY) bool {
	if !resource.Matcher.Equals(other.Matcher) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigY` fields for violations and returns them.
func (resource XychartXYSeriesConfigY) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Matcher.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type XychartXYSeriesConfigColor struct {
	Matcher MatcherConfig `json:"matcher"`
}

// NewXychartXYSeriesConfigColor creates a new XychartXYSeriesConfigColor object.
func NewXychartXYSeriesConfigColor() *XychartXYSeriesConfigColor {
	return &XychartXYSeriesConfigColor{
		Matcher: *NewMatcherConfig(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigColor` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XychartXYSeriesConfigColor) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "matcher"
	if fields["matcher"] != nil {
		if string(fields["matcher"]) != "null" {

			resource.Matcher = MatcherConfig{}
			if err := resource.Matcher.UnmarshalJSONStrict(fields["matcher"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is null"))...)

		}
		delete(fields, "matcher")
	} else {
		errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("XychartXYSeriesConfigColor", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XychartXYSeriesConfigColor` objects.
func (resource XychartXYSeriesConfigColor) Equals(other XychartXYSeriesConfigColor) bool {
	if !resource.Matcher.Equals(other.Matcher) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigColor` fields for violations and returns them.
func (resource XychartXYSeriesConfigColor) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Matcher.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type XychartXYSeriesConfigSize struct {
	Matcher MatcherConfig `json:"matcher"`
}

// NewXychartXYSeriesConfigSize creates a new XychartXYSeriesConfigSize object.
func NewXychartXYSeriesConfigSize() *XychartXYSeriesConfigSize {
	return &XychartXYSeriesConfigSize{
		Matcher: *NewMatcherConfig(),
	}
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XychartXYSeriesConfigSize` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XychartXYSeriesConfigSize) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "matcher"
	if fields["matcher"] != nil {
		if string(fields["matcher"]) != "null" {

			resource.Matcher = MatcherConfig{}
			if err := resource.Matcher.UnmarshalJSONStrict(fields["matcher"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is null"))...)

		}
		delete(fields, "matcher")
	} else {
		errs = append(errs, cog.MakeBuildErrors("matcher", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("XychartXYSeriesConfigSize", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XychartXYSeriesConfigSize` objects.
func (resource XychartXYSeriesConfigSize) Equals(other XychartXYSeriesConfigSize) bool {
	if !resource.Matcher.Equals(other.Matcher) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `XychartXYSeriesConfigSize` fields for violations and returns them.
func (resource XychartXYSeriesConfigSize) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Matcher.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("matcher", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to xychart panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "xychart", OptionsUnmarshaler: func(raw []byte) (any, error) {
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
		}, FieldConfigUnmarshaler: func(raw []byte) (any, error) {
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
		}, GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}
			if panel, ok := inputPanel.(dashboard.Panel); ok {
				return PanelConverter(panel)
			}
			if panel, ok := inputPanel.(*dashboardv2beta1.VizConfigKind); ok {
				return VisualizationConverter(*panel)
			}

			return VisualizationConverter(inputPanel.(dashboardv2beta1.VizConfigKind))
		},
	}
}
