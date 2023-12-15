// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	View     MapViewConfig            `json:"view"`
	Controls ControlsOptions          `json:"controls"`
	Basemap  common.MapLayerOptions   `json:"basemap"`
	Layers   []common.MapLayerOptions `json:"layers"`
	Tooltip  TooltipOptions           `json:"tooltip"`
}

type MapViewConfig struct {
	Id        string  `json:"id"`
	Lat       *int64  `json:"lat,omitempty"`
	Lon       *int64  `json:"lon,omitempty"`
	Zoom      *int64  `json:"zoom,omitempty"`
	MinZoom   *int64  `json:"minZoom,omitempty"`
	MaxZoom   *int64  `json:"maxZoom,omitempty"`
	Padding   *int64  `json:"padding,omitempty"`
	AllLayers *bool   `json:"allLayers,omitempty"`
	LastOnly  *bool   `json:"lastOnly,omitempty"`
	Layer     *string `json:"layer,omitempty"`
	Shared    *bool   `json:"shared,omitempty"`
}

type ControlsOptions struct {
	// Zoom (upper left)
	ShowZoom *bool `json:"showZoom,omitempty"`
	// let the mouse wheel zoom
	MouseWheelZoom *bool `json:"mouseWheelZoom,omitempty"`
	// Lower right
	ShowAttribution *bool `json:"showAttribution,omitempty"`
	// Scale options
	ShowScale *bool `json:"showScale,omitempty"`
	// Show debug
	ShowDebug *bool `json:"showDebug,omitempty"`
	// Show measure
	ShowMeasure *bool `json:"showMeasure,omitempty"`
}

type TooltipOptions struct {
	Mode TooltipMode `json:"mode"`
}

type TooltipMode string

const (
	TooltipModeNone    TooltipMode = "none"
	TooltipModeDetails TooltipMode = "details"
)

type MapCenterID string

const (
	MapCenterIDZero   MapCenterID = "zero"
	MapCenterIDCoords MapCenterID = "coords"
	MapCenterIDFit    MapCenterID = "fit"
)
