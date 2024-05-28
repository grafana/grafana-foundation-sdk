// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
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

type Placement struct {
	Top    *float64 `json:"top,omitempty"`
	Left   *float64 `json:"left,omitempty"`
	Right  *float64 `json:"right,omitempty"`
	Bottom *float64 `json:"bottom,omitempty"`
	Width  *float64 `json:"width,omitempty"`
	Height *float64 `json:"height,omitempty"`
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

type LineConfig struct {
	Color *common.ColorDimensionConfig `json:"color,omitempty"`
	Width *float64                     `json:"width,omitempty"`
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

type Options struct {
	// Enable inline editing
	InlineEditing bool `json:"inlineEditing"`
	// Show all available element types
	ShowAdvancedTypes bool `json:"showAdvancedTypes"`
	// Enable pan and zoom
	PanZoom bool `json:"panZoom"`
	// The root element of canvas (frame), where all canvas elements are nested
	// TODO: Figure out how to define a default value for this
	Root struct {
		// Name of the root element
		Name string `json:"name"`
		// Type of root element (frame)
		Type string `json:"type"`
		// The list of canvas elements attached to the root element
		Elements []CanvasElementOptions `json:"elements"`
	} `json:"root"`
}

func VariantConfig() cogvariants.PanelcfgConfig {
	return cogvariants.PanelcfgConfig{
		Identifier: "canvas",
		OptionsUnmarshaler: func(raw []byte) (any, error) {
			options := Options{}

			if err := json.Unmarshal(raw, &options); err != nil {
				return nil, err
			}

			return options, nil
		},
	}
}
