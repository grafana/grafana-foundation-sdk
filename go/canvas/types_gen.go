// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

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
)

type HorizontalConstraint string

const (
	HorizontalConstraintLeft      HorizontalConstraint = "left"
	HorizontalConstraintRight     HorizontalConstraint = "right"
	HorizontalConstraintLeftRight HorizontalConstraint = "leftright"
	HorizontalConstraintCenter    HorizontalConstraint = "center"
	HorizontalConstraintScale     HorizontalConstraint = "scale"
)

type VerticalConstraint string

const (
	VerticalConstraintTop       VerticalConstraint = "top"
	VerticalConstraintBottom    VerticalConstraint = "bottom"
	VerticalConstraintTopBottom VerticalConstraint = "topbottom"
	VerticalConstraintCenter    VerticalConstraint = "center"
	VerticalConstraintScale     VerticalConstraint = "scale"
)

type Constraint struct {
	Horizontal *HorizontalConstraint `json:"horizontal,omitempty"`
	Vertical   *VerticalConstraint   `json:"vertical,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Constraint` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Constraint) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "horizontal"
	if fields["horizontal"] != nil {
		if string(fields["horizontal"]) != "null" {
			if err := json.Unmarshal(fields["horizontal"], &resource.Horizontal); err != nil {
				errs = append(errs, cog.MakeBuildErrors("horizontal", err)...)
			}

		}
		delete(fields, "horizontal")

	}
	// Field "vertical"
	if fields["vertical"] != nil {
		if string(fields["vertical"]) != "null" {
			if err := json.Unmarshal(fields["vertical"], &resource.Vertical); err != nil {
				errs = append(errs, cog.MakeBuildErrors("vertical", err)...)
			}

		}
		delete(fields, "vertical")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Constraint", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Constraint` objects.
func (resource Constraint) Equals(other Constraint) bool {
	if resource.Horizontal == nil && other.Horizontal != nil || resource.Horizontal != nil && other.Horizontal == nil {
		return false
	}

	if resource.Horizontal != nil {
		if *resource.Horizontal != *other.Horizontal {
			return false
		}
	}
	if resource.Vertical == nil && other.Vertical != nil || resource.Vertical != nil && other.Vertical == nil {
		return false
	}

	if resource.Vertical != nil {
		if *resource.Vertical != *other.Vertical {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Constraint` fields for violations and returns them.
func (resource Constraint) Validate() error {
	return nil
}

type Placement struct {
	Top    *float64 `json:"top,omitempty"`
	Left   *float64 `json:"left,omitempty"`
	Right  *float64 `json:"right,omitempty"`
	Bottom *float64 `json:"bottom,omitempty"`
	Width  *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Placement` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Placement) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "top"
	if fields["top"] != nil {
		if string(fields["top"]) != "null" {
			if err := json.Unmarshal(fields["top"], &resource.Top); err != nil {
				errs = append(errs, cog.MakeBuildErrors("top", err)...)
			}

		}
		delete(fields, "top")

	}
	// Field "left"
	if fields["left"] != nil {
		if string(fields["left"]) != "null" {
			if err := json.Unmarshal(fields["left"], &resource.Left); err != nil {
				errs = append(errs, cog.MakeBuildErrors("left", err)...)
			}

		}
		delete(fields, "left")

	}
	// Field "right"
	if fields["right"] != nil {
		if string(fields["right"]) != "null" {
			if err := json.Unmarshal(fields["right"], &resource.Right); err != nil {
				errs = append(errs, cog.MakeBuildErrors("right", err)...)
			}

		}
		delete(fields, "right")

	}
	// Field "bottom"
	if fields["bottom"] != nil {
		if string(fields["bottom"]) != "null" {
			if err := json.Unmarshal(fields["bottom"], &resource.Bottom); err != nil {
				errs = append(errs, cog.MakeBuildErrors("bottom", err)...)
			}

		}
		delete(fields, "bottom")

	}
	// Field "width"
	if fields["width"] != nil {
		if string(fields["width"]) != "null" {
			if err := json.Unmarshal(fields["width"], &resource.Width); err != nil {
				errs = append(errs, cog.MakeBuildErrors("width", err)...)
			}

		}
		delete(fields, "width")

	}
	// Field "height"
	if fields["height"] != nil {
		if string(fields["height"]) != "null" {
			if err := json.Unmarshal(fields["height"], &resource.Height); err != nil {
				errs = append(errs, cog.MakeBuildErrors("height", err)...)
			}

		}
		delete(fields, "height")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Placement", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Placement` objects.
func (resource Placement) Equals(other Placement) bool {
	if resource.Top == nil && other.Top != nil || resource.Top != nil && other.Top == nil {
		return false
	}

	if resource.Top != nil {
		if *resource.Top != *other.Top {
			return false
		}
	}
	if resource.Left == nil && other.Left != nil || resource.Left != nil && other.Left == nil {
		return false
	}

	if resource.Left != nil {
		if *resource.Left != *other.Left {
			return false
		}
	}
	if resource.Right == nil && other.Right != nil || resource.Right != nil && other.Right == nil {
		return false
	}

	if resource.Right != nil {
		if *resource.Right != *other.Right {
			return false
		}
	}
	if resource.Bottom == nil && other.Bottom != nil || resource.Bottom != nil && other.Bottom == nil {
		return false
	}

	if resource.Bottom != nil {
		if *resource.Bottom != *other.Bottom {
			return false
		}
	}
	if resource.Width == nil && other.Width != nil || resource.Width != nil && other.Width == nil {
		return false
	}

	if resource.Width != nil {
		if *resource.Width != *other.Width {
			return false
		}
	}
	if resource.Height == nil && other.Height != nil || resource.Height != nil && other.Height == nil {
		return false
	}

	if resource.Height != nil {
		if *resource.Height != *other.Height {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Placement` fields for violations and returns them.
func (resource Placement) Validate() error {
	return nil
}

type BackgroundImageSize string

const (
	BackgroundImageSizeOriginal BackgroundImageSize = "original"
	BackgroundImageSizeContain  BackgroundImageSize = "contain"
	BackgroundImageSizeCover    BackgroundImageSize = "cover"
	BackgroundImageSizeFill     BackgroundImageSize = "fill"
	BackgroundImageSizeTile     BackgroundImageSize = "tile"
)

type BackgroundConfig struct {
	Color *common.ColorDimensionConfig    `json:"color,omitempty"`
	Image *common.ResourceDimensionConfig `json:"image,omitempty"`
	Size  *BackgroundImageSize            `json:"size,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BackgroundConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *BackgroundConfig) UnmarshalJSONStrict(raw []byte) error {
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

			resource.Color = &common.ColorDimensionConfig{}
			if err := resource.Color.UnmarshalJSONStrict(fields["color"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}

		}
		delete(fields, "color")

	}
	// Field "image"
	if fields["image"] != nil {
		if string(fields["image"]) != "null" {

			resource.Image = &common.ResourceDimensionConfig{}
			if err := resource.Image.UnmarshalJSONStrict(fields["image"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("image", err)...)
			}

		}
		delete(fields, "image")

	}
	// Field "size"
	if fields["size"] != nil {
		if string(fields["size"]) != "null" {
			if err := json.Unmarshal(fields["size"], &resource.Size); err != nil {
				errs = append(errs, cog.MakeBuildErrors("size", err)...)
			}

		}
		delete(fields, "size")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("BackgroundConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `BackgroundConfig` objects.
func (resource BackgroundConfig) Equals(other BackgroundConfig) bool {
	if resource.Color == nil && other.Color != nil || resource.Color != nil && other.Color == nil {
		return false
	}

	if resource.Color != nil {
		if !resource.Color.Equals(*other.Color) {
			return false
		}
	}
	if resource.Image == nil && other.Image != nil || resource.Image != nil && other.Image == nil {
		return false
	}

	if resource.Image != nil {
		if !resource.Image.Equals(*other.Image) {
			return false
		}
	}
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `BackgroundConfig` fields for violations and returns them.
func (resource BackgroundConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.Color != nil {
		if err := resource.Color.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("color", err)...)
		}
	}
	if resource.Image != nil {
		if err := resource.Image.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("image", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type LineConfig struct {
	Color *common.ColorDimensionConfig `json:"color,omitempty"`
	Width *float64                     `json:"width,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LineConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LineConfig) UnmarshalJSONStrict(raw []byte) error {
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

			resource.Color = &common.ColorDimensionConfig{}
			if err := resource.Color.UnmarshalJSONStrict(fields["color"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}

		}
		delete(fields, "color")

	}
	// Field "width"
	if fields["width"] != nil {
		if string(fields["width"]) != "null" {
			if err := json.Unmarshal(fields["width"], &resource.Width); err != nil {
				errs = append(errs, cog.MakeBuildErrors("width", err)...)
			}

		}
		delete(fields, "width")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LineConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LineConfig` objects.
func (resource LineConfig) Equals(other LineConfig) bool {
	if resource.Color == nil && other.Color != nil || resource.Color != nil && other.Color == nil {
		return false
	}

	if resource.Color != nil {
		if !resource.Color.Equals(*other.Color) {
			return false
		}
	}
	if resource.Width == nil && other.Width != nil || resource.Width != nil && other.Width == nil {
		return false
	}

	if resource.Width != nil {
		if *resource.Width != *other.Width {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `LineConfig` fields for violations and returns them.
func (resource LineConfig) Validate() error {
	var errs cog.BuildErrors
	if resource.Color != nil {
		if err := resource.Color.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("color", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type HttpRequestMethod string

const (
	HttpRequestMethodGET  HttpRequestMethod = "GET"
	HttpRequestMethodPOST HttpRequestMethod = "POST"
	HttpRequestMethodPUT  HttpRequestMethod = "PUT"
)

type ConnectionCoordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConnectionCoordinates` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ConnectionCoordinates) UnmarshalJSONStrict(raw []byte) error {
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
		} else {
			errs = append(errs, cog.MakeBuildErrors("x", errors.New("required field is null"))...)

		}
		delete(fields, "x")
	} else {
		errs = append(errs, cog.MakeBuildErrors("x", errors.New("required field is missing from input"))...)
	}
	// Field "y"
	if fields["y"] != nil {
		if string(fields["y"]) != "null" {
			if err := json.Unmarshal(fields["y"], &resource.Y); err != nil {
				errs = append(errs, cog.MakeBuildErrors("y", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("y", errors.New("required field is null"))...)

		}
		delete(fields, "y")
	} else {
		errs = append(errs, cog.MakeBuildErrors("y", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ConnectionCoordinates", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ConnectionCoordinates` objects.
func (resource ConnectionCoordinates) Equals(other ConnectionCoordinates) bool {
	if resource.X != other.X {
		return false
	}
	if resource.Y != other.Y {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ConnectionCoordinates` fields for violations and returns them.
func (resource ConnectionCoordinates) Validate() error {
	return nil
}

type ConnectionPath string

const (
	ConnectionPathStraight ConnectionPath = "straight"
)

type CanvasConnection struct {
	Source     ConnectionCoordinates        `json:"source"`
	Target     ConnectionCoordinates        `json:"target"`
	TargetName *string                      `json:"targetName,omitempty"`
	Path       ConnectionPath               `json:"path"`
	Color      *common.ColorDimensionConfig `json:"color,omitempty"`
	Size       *common.ScaleDimensionConfig `json:"size,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CanvasConnection` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CanvasConnection) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "source"
	if fields["source"] != nil {
		if string(fields["source"]) != "null" {

			resource.Source = ConnectionCoordinates{}
			if err := resource.Source.UnmarshalJSONStrict(fields["source"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("source", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("source", errors.New("required field is null"))...)

		}
		delete(fields, "source")
	} else {
		errs = append(errs, cog.MakeBuildErrors("source", errors.New("required field is missing from input"))...)
	}
	// Field "target"
	if fields["target"] != nil {
		if string(fields["target"]) != "null" {

			resource.Target = ConnectionCoordinates{}
			if err := resource.Target.UnmarshalJSONStrict(fields["target"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("target", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("target", errors.New("required field is null"))...)

		}
		delete(fields, "target")
	} else {
		errs = append(errs, cog.MakeBuildErrors("target", errors.New("required field is missing from input"))...)
	}
	// Field "targetName"
	if fields["targetName"] != nil {
		if string(fields["targetName"]) != "null" {
			if err := json.Unmarshal(fields["targetName"], &resource.TargetName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("targetName", err)...)
			}

		}
		delete(fields, "targetName")

	}
	// Field "path"
	if fields["path"] != nil {
		if string(fields["path"]) != "null" {
			if err := json.Unmarshal(fields["path"], &resource.Path); err != nil {
				errs = append(errs, cog.MakeBuildErrors("path", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("path", errors.New("required field is null"))...)

		}
		delete(fields, "path")
	} else {
		errs = append(errs, cog.MakeBuildErrors("path", errors.New("required field is missing from input"))...)
	}
	// Field "color"
	if fields["color"] != nil {
		if string(fields["color"]) != "null" {

			resource.Color = &common.ColorDimensionConfig{}
			if err := resource.Color.UnmarshalJSONStrict(fields["color"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("color", err)...)
			}

		}
		delete(fields, "color")

	}
	// Field "size"
	if fields["size"] != nil {
		if string(fields["size"]) != "null" {

			resource.Size = &common.ScaleDimensionConfig{}
			if err := resource.Size.UnmarshalJSONStrict(fields["size"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("size", err)...)
			}

		}
		delete(fields, "size")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CanvasConnection", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CanvasConnection` objects.
func (resource CanvasConnection) Equals(other CanvasConnection) bool {
	if !resource.Source.Equals(other.Source) {
		return false
	}
	if !resource.Target.Equals(other.Target) {
		return false
	}
	if resource.TargetName == nil && other.TargetName != nil || resource.TargetName != nil && other.TargetName == nil {
		return false
	}

	if resource.TargetName != nil {
		if *resource.TargetName != *other.TargetName {
			return false
		}
	}
	if resource.Path != other.Path {
		return false
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

// Validate checks all the validation constraints that may be defined on `CanvasConnection` fields for violations and returns them.
func (resource CanvasConnection) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Source.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("source", err)...)
	}
	if err := resource.Target.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("target", err)...)
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

type CanvasElementOptions struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// TODO: figure out how to define this (element config(s))
	Config      any                `json:"config,omitempty"`
	Constraint  *Constraint        `json:"constraint,omitempty"`
	Placement   *Placement         `json:"placement,omitempty"`
	Background  *BackgroundConfig  `json:"background,omitempty"`
	Border      *LineConfig        `json:"border,omitempty"`
	Connections []CanvasConnection `json:"connections,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CanvasElementOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CanvasElementOptions) UnmarshalJSONStrict(raw []byte) error {
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
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is null"))...)

		}
		delete(fields, "name")
	} else {
		errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is missing from input"))...)
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "config"
	if fields["config"] != nil {
		if string(fields["config"]) != "null" {
			if err := json.Unmarshal(fields["config"], &resource.Config); err != nil {
				errs = append(errs, cog.MakeBuildErrors("config", err)...)
			}

		}
		delete(fields, "config")

	}
	// Field "constraint"
	if fields["constraint"] != nil {
		if string(fields["constraint"]) != "null" {

			resource.Constraint = &Constraint{}
			if err := resource.Constraint.UnmarshalJSONStrict(fields["constraint"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("constraint", err)...)
			}

		}
		delete(fields, "constraint")

	}
	// Field "placement"
	if fields["placement"] != nil {
		if string(fields["placement"]) != "null" {

			resource.Placement = &Placement{}
			if err := resource.Placement.UnmarshalJSONStrict(fields["placement"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("placement", err)...)
			}

		}
		delete(fields, "placement")

	}
	// Field "background"
	if fields["background"] != nil {
		if string(fields["background"]) != "null" {

			resource.Background = &BackgroundConfig{}
			if err := resource.Background.UnmarshalJSONStrict(fields["background"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("background", err)...)
			}

		}
		delete(fields, "background")

	}
	// Field "border"
	if fields["border"] != nil {
		if string(fields["border"]) != "null" {

			resource.Border = &LineConfig{}
			if err := resource.Border.UnmarshalJSONStrict(fields["border"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("border", err)...)
			}

		}
		delete(fields, "border")

	}
	// Field "connections"
	if fields["connections"] != nil {
		if string(fields["connections"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["connections"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 CanvasConnection

				result1 = CanvasConnection{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("connections["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Connections = append(resource.Connections, result1)
			}

		}
		delete(fields, "connections")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CanvasElementOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CanvasElementOptions` objects.
func (resource CanvasElementOptions) Equals(other CanvasElementOptions) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Config, other.Config) {
		return false
	}
	if resource.Constraint == nil && other.Constraint != nil || resource.Constraint != nil && other.Constraint == nil {
		return false
	}

	if resource.Constraint != nil {
		if !resource.Constraint.Equals(*other.Constraint) {
			return false
		}
	}
	if resource.Placement == nil && other.Placement != nil || resource.Placement != nil && other.Placement == nil {
		return false
	}

	if resource.Placement != nil {
		if !resource.Placement.Equals(*other.Placement) {
			return false
		}
	}
	if resource.Background == nil && other.Background != nil || resource.Background != nil && other.Background == nil {
		return false
	}

	if resource.Background != nil {
		if !resource.Background.Equals(*other.Background) {
			return false
		}
	}
	if resource.Border == nil && other.Border != nil || resource.Border != nil && other.Border == nil {
		return false
	}

	if resource.Border != nil {
		if !resource.Border.Equals(*other.Border) {
			return false
		}
	}

	if len(resource.Connections) != len(other.Connections) {
		return false
	}

	for i1 := range resource.Connections {
		if !resource.Connections[i1].Equals(other.Connections[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CanvasElementOptions` fields for violations and returns them.
func (resource CanvasElementOptions) Validate() error {
	var errs cog.BuildErrors
	if resource.Constraint != nil {
		if err := resource.Constraint.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("constraint", err)...)
		}
	}
	if resource.Placement != nil {
		if err := resource.Placement.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("placement", err)...)
		}
	}
	if resource.Background != nil {
		if err := resource.Background.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("background", err)...)
		}
	}
	if resource.Border != nil {
		if err := resource.Border.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("border", err)...)
		}
	}

	for i1 := range resource.Connections {
		if err := resource.Connections[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("connections["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type Options struct {
	// Enable inline editing
	InlineEditing bool `json:"inlineEditing"`
	// Show all available element types
	ShowAdvancedTypes bool `json:"showAdvancedTypes"`
	// The root element of canvas (frame), where all canvas elements are nested
	// TODO: Figure out how to define a default value for this
	Root CanvasOptionsRoot `json:"root"`
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
	// Field "inlineEditing"
	if fields["inlineEditing"] != nil {
		if string(fields["inlineEditing"]) != "null" {
			if err := json.Unmarshal(fields["inlineEditing"], &resource.InlineEditing); err != nil {
				errs = append(errs, cog.MakeBuildErrors("inlineEditing", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("inlineEditing", errors.New("required field is null"))...)

		}
		delete(fields, "inlineEditing")
	} else {
		errs = append(errs, cog.MakeBuildErrors("inlineEditing", errors.New("required field is missing from input"))...)
	}
	// Field "showAdvancedTypes"
	if fields["showAdvancedTypes"] != nil {
		if string(fields["showAdvancedTypes"]) != "null" {
			if err := json.Unmarshal(fields["showAdvancedTypes"], &resource.ShowAdvancedTypes); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showAdvancedTypes", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("showAdvancedTypes", errors.New("required field is null"))...)

		}
		delete(fields, "showAdvancedTypes")
	} else {
		errs = append(errs, cog.MakeBuildErrors("showAdvancedTypes", errors.New("required field is missing from input"))...)
	}
	// Field "root"
	if fields["root"] != nil {
		if string(fields["root"]) != "null" {

			resource.Root = CanvasOptionsRoot{}
			if err := resource.Root.UnmarshalJSONStrict(fields["root"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("root", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("root", errors.New("required field is null"))...)

		}
		delete(fields, "root")
	} else {
		errs = append(errs, cog.MakeBuildErrors("root", errors.New("required field is missing from input"))...)
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
	if resource.InlineEditing != other.InlineEditing {
		return false
	}
	if resource.ShowAdvancedTypes != other.ShowAdvancedTypes {
		return false
	}
	if !resource.Root.Equals(other.Root) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Root.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("root", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type CanvasOptionsRoot struct {
	// Name of the root element
	Name string `json:"name"`
	// Type of root element (frame)
	Type string `json:"type"`
	// The list of canvas elements attached to the root element
	Elements []CanvasElementOptions `json:"elements"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CanvasOptionsRoot` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CanvasOptionsRoot) UnmarshalJSONStrict(raw []byte) error {
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
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is null"))...)

		}
		delete(fields, "name")
	} else {
		errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is missing from input"))...)
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "elements"
	if fields["elements"] != nil {
		if string(fields["elements"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["elements"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 CanvasElementOptions

				result1 = CanvasElementOptions{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("elements["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Elements = append(resource.Elements, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("elements", errors.New("required field is null"))...)

		}
		delete(fields, "elements")
	} else {
		errs = append(errs, cog.MakeBuildErrors("elements", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CanvasOptionsRoot", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `CanvasOptionsRoot` objects.
func (resource CanvasOptionsRoot) Equals(other CanvasOptionsRoot) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.Type != other.Type {
		return false
	}

	if len(resource.Elements) != len(other.Elements) {
		return false
	}

	for i1 := range resource.Elements {
		if !resource.Elements[i1].Equals(other.Elements[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CanvasOptionsRoot` fields for violations and returns them.
func (resource CanvasOptionsRoot) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Elements {
		if err := resource.Elements[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("elements["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// VariantConfig returns the configuration related to canvas panels.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "canvas",
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
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
