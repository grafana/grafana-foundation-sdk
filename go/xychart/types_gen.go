// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `XYDimensionConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *XYDimensionConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "frame"
	if fields["frame"] != nil {
		if string(fields["frame"]) != "null" {
			if err := json.Unmarshal(fields["frame"], &resource.Frame); err != nil {
				errs = append(errs, cog.MakeBuildErrors("frame", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("frame", errors.New("required field is null"))...)

		}
		delete(fields, "frame")
	} else {
		errs = append(errs, cog.MakeBuildErrors("frame", errors.New("required field is missing from input"))...)
	}
	// Field "x"
	if fields["x"] != nil {
		if string(fields["x"]) != "null" {
			if err := json.Unmarshal(fields["x"], &resource.X); err != nil {
				errs = append(errs, cog.MakeBuildErrors("x", err)...)
			}

		}
		delete(fields, "x")

	}
	// Field "exclude"
	if fields["exclude"] != nil {
		if string(fields["exclude"]) != "null" {

			if err := json.Unmarshal(fields["exclude"], &resource.Exclude); err != nil {
				errs = append(errs, cog.MakeBuildErrors("exclude", err)...)
			}

		}
		delete(fields, "exclude")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("XYDimensionConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `XYDimensionConfig` objects.
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

// Validate checks all the validation constraints that may be defined on `XYDimensionConfig` fields for violations and returns them.
func (resource XYDimensionConfig) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Frame >= 0) {
		errs = append(errs, cog.MakeBuildErrors(
			"frame",
			errors.New("must be >= 0"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

			resource.PointSize = &common.ScaleDimensionConfig{}
			if err := resource.PointSize.UnmarshalJSONStrict(fields["pointSize"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
			}

		}
		delete(fields, "pointSize")

	}
	// Field "pointColor"
	if fields["pointColor"] != nil {
		if string(fields["pointColor"]) != "null" {

			resource.PointColor = &common.ColorDimensionConfig{}
			if err := resource.PointColor.UnmarshalJSONStrict(fields["pointColor"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointColor", err)...)
			}

		}
		delete(fields, "pointColor")

	}
	// Field "lineColor"
	if fields["lineColor"] != nil {
		if string(fields["lineColor"]) != "null" {

			resource.LineColor = &common.ColorDimensionConfig{}
			if err := resource.LineColor.UnmarshalJSONStrict(fields["lineColor"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineColor", err)...)
			}

		}
		delete(fields, "lineColor")

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
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

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
	// Field "labelValue"
	if fields["labelValue"] != nil {
		if string(fields["labelValue"]) != "null" {

			resource.LabelValue = &common.TextDimensionConfig{}
			if err := resource.LabelValue.UnmarshalJSONStrict(fields["labelValue"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labelValue", err)...)
			}

		}
		delete(fields, "labelValue")

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

// Validate checks all the validation constraints that may be defined on `FieldConfig` fields for violations and returns them.
func (resource FieldConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.PointSize != nil {
		if err := resource.PointSize.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
		}
	}
	if resource.PointColor != nil {
		if err := resource.PointColor.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pointColor", err)...)
		}
	}
	if resource.LineColor != nil {
		if err := resource.LineColor.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("lineColor", err)...)
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
	if resource.LineStyle != nil {
		if err := resource.LineStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
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
	if resource.LabelValue != nil {
		if err := resource.LabelValue.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("labelValue", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ScatterSeriesConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ScatterSeriesConfig) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "x"
	if fields["x"] != nil {
		if string(fields["x"]) != "null" {
			if err := json.Unmarshal(fields["x"], &resource.X); err != nil {
				errs = append(errs, cog.MakeBuildErrors("x", err)...)
			}

		}
		delete(fields, "x")

	}
	// Field "y"
	if fields["y"] != nil {
		if string(fields["y"]) != "null" {
			if err := json.Unmarshal(fields["y"], &resource.Y); err != nil {
				errs = append(errs, cog.MakeBuildErrors("y", err)...)
			}

		}
		delete(fields, "y")

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

			resource.PointSize = &common.ScaleDimensionConfig{}
			if err := resource.PointSize.UnmarshalJSONStrict(fields["pointSize"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
			}

		}
		delete(fields, "pointSize")

	}
	// Field "pointColor"
	if fields["pointColor"] != nil {
		if string(fields["pointColor"]) != "null" {

			resource.PointColor = &common.ColorDimensionConfig{}
			if err := resource.PointColor.UnmarshalJSONStrict(fields["pointColor"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("pointColor", err)...)
			}

		}
		delete(fields, "pointColor")

	}
	// Field "lineColor"
	if fields["lineColor"] != nil {
		if string(fields["lineColor"]) != "null" {

			resource.LineColor = &common.ColorDimensionConfig{}
			if err := resource.LineColor.UnmarshalJSONStrict(fields["lineColor"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lineColor", err)...)
			}

		}
		delete(fields, "lineColor")

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
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

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
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "labelValue"
	if fields["labelValue"] != nil {
		if string(fields["labelValue"]) != "null" {

			resource.LabelValue = &common.TextDimensionConfig{}
			if err := resource.LabelValue.UnmarshalJSONStrict(fields["labelValue"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labelValue", err)...)
			}

		}
		delete(fields, "labelValue")

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
		errs = append(errs, cog.MakeBuildErrors("ScatterSeriesConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ScatterSeriesConfig` objects.
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

// Validate checks all the validation constraints that may be defined on `ScatterSeriesConfig` fields for violations and returns them.
func (resource ScatterSeriesConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.PointSize != nil {
		if err := resource.PointSize.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pointSize", err)...)
		}
	}
	if resource.PointColor != nil {
		if err := resource.PointColor.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("pointColor", err)...)
		}
	}
	if resource.LineColor != nil {
		if err := resource.LineColor.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("lineColor", err)...)
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
	if resource.LineStyle != nil {
		if err := resource.LineStyle.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("lineStyle", err)...)
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
	if resource.LabelValue != nil {
		if err := resource.LabelValue.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("labelValue", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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
	// Field "seriesMapping"
	if fields["seriesMapping"] != nil {
		if string(fields["seriesMapping"]) != "null" {
			if err := json.Unmarshal(fields["seriesMapping"], &resource.SeriesMapping); err != nil {
				errs = append(errs, cog.MakeBuildErrors("seriesMapping", err)...)
			}

		}
		delete(fields, "seriesMapping")

	}
	// Field "dims"
	if fields["dims"] != nil {
		if string(fields["dims"]) != "null" {

			resource.Dims = XYDimensionConfig{}
			if err := resource.Dims.UnmarshalJSONStrict(fields["dims"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("dims", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("dims", errors.New("required field is null"))...)

		}
		delete(fields, "dims")
	} else {
		errs = append(errs, cog.MakeBuildErrors("dims", errors.New("required field is missing from input"))...)
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
				var result1 ScatterSeriesConfig

				result1 = ScatterSeriesConfig{}
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

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Dims.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("dims", err)...)
	}
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

// VariantConfig returns the configuration related to xychart panels.
// This configuration describes how to unmarshal it, convert it to code, …
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
