// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"encoding/json"
	"reflect"

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

type Placement struct {
	Top      *float64 `json:"top,omitempty"`
	Left     *float64 `json:"left,omitempty"`
	Right    *float64 `json:"right,omitempty"`
	Bottom   *float64 `json:"bottom,omitempty"`
	Width    *float64 `json:"width,omitempty"`
	Height   *float64 `json:"height,omitempty"`
	Rotation *float64 `json:"rotation,omitempty"`
}

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
	if resource.Rotation == nil && other.Rotation != nil || resource.Rotation != nil && other.Rotation == nil {
		return false
	}

	if resource.Rotation != nil {
		if *resource.Rotation != *other.Rotation {
			return false
		}
	}

	return true
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

type LineConfig struct {
	Color  *common.ColorDimensionConfig `json:"color,omitempty"`
	Width  *float64                     `json:"width,omitempty"`
	Radius *float64                     `json:"radius,omitempty"`
}

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
	if resource.Radius == nil && other.Radius != nil || resource.Radius != nil && other.Radius == nil {
		return false
	}

	if resource.Radius != nil {
		if *resource.Radius != *other.Radius {
			return false
		}
	}

	return true
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

func (resource ConnectionCoordinates) Equals(other ConnectionCoordinates) bool {
	if resource.X != other.X {
		return false
	}
	if resource.Y != other.Y {
		return false
	}

	return true
}

type ConnectionPath string

const (
	ConnectionPathStraight ConnectionPath = "straight"
)

type CanvasConnection struct {
	Source         ConnectionCoordinates        `json:"source"`
	Target         ConnectionCoordinates        `json:"target"`
	TargetName     *string                      `json:"targetName,omitempty"`
	Path           ConnectionPath               `json:"path"`
	Color          *common.ColorDimensionConfig `json:"color,omitempty"`
	Size           *common.ScaleDimensionConfig `json:"size,omitempty"`
	Vertices       []ConnectionCoordinates      `json:"vertices,omitempty"`
	SourceOriginal *ConnectionCoordinates       `json:"sourceOriginal,omitempty"`
	TargetOriginal *ConnectionCoordinates       `json:"targetOriginal,omitempty"`
}

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

	if len(resource.Vertices) != len(other.Vertices) {
		return false
	}

	for i1 := range resource.Vertices {
		if !resource.Vertices[i1].Equals(other.Vertices[i1]) {
			return false
		}
	}
	if resource.SourceOriginal == nil && other.SourceOriginal != nil || resource.SourceOriginal != nil && other.SourceOriginal == nil {
		return false
	}

	if resource.SourceOriginal != nil {
		if !resource.SourceOriginal.Equals(*other.SourceOriginal) {
			return false
		}
	}
	if resource.TargetOriginal == nil && other.TargetOriginal != nil || resource.TargetOriginal != nil && other.TargetOriginal == nil {
		return false
	}

	if resource.TargetOriginal != nil {
		if !resource.TargetOriginal.Equals(*other.TargetOriginal) {
			return false
		}
	}

	return true
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

type Options struct {
	// Enable inline editing
	InlineEditing bool `json:"inlineEditing"`
	// Show all available element types
	ShowAdvancedTypes bool `json:"showAdvancedTypes"`
	// Enable pan and zoom
	PanZoom bool `json:"panZoom"`
	// Enable infinite pan
	InfinitePan bool `json:"infinitePan"`
	// The root element of canvas (frame), where all canvas elements are nested
	// TODO: Figure out how to define a default value for this
	Root CanvasOptionsRoot `json:"root"`
}

func (resource Options) Equals(other Options) bool {
	if resource.InlineEditing != other.InlineEditing {
		return false
	}
	if resource.ShowAdvancedTypes != other.ShowAdvancedTypes {
		return false
	}
	if resource.PanZoom != other.PanZoom {
		return false
	}
	if resource.InfinitePan != other.InfinitePan {
		return false
	}
	if !resource.Root.Equals(other.Root) {
		return false
	}

	return true
}

type CanvasOptionsRoot struct {
	// Name of the root element
	Name string `json:"name"`
	// Type of root element (frame)
	Type string `json:"type"`
	// The list of canvas elements attached to the root element
	Elements []CanvasElementOptions `json:"elements"`
}

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
		GoConverter: func(inputPanel any) string {
			if panel, ok := inputPanel.(*dashboard.Panel); ok {
				return PanelConverter(*panel)
			}

			return PanelConverter(inputPanel.(dashboard.Panel))
		},
	}
}
