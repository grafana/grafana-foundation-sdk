// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Options struct {
	View     MapViewConfig            `json:"view"`
	Controls ControlsOptions          `json:"controls"`
	Basemap  common.MapLayerOptions   `json:"basemap"`
	Layers   []common.MapLayerOptions `json:"layers"`
	Tooltip  TooltipOptions           `json:"tooltip"`
}

func (resource Options) Equals(other Options) bool {
	if !resource.View.Equals(other.View) {
		return false
	}
	if !resource.Controls.Equals(other.Controls) {
		return false
	}
	if !resource.Basemap.Equals(other.Basemap) {
		return false
	}

	if len(resource.Layers) != len(other.Layers) {
		return false
	}

	for i1 := range resource.Layers {
		if !resource.Layers[i1].Equals(other.Layers[i1]) {
			return false
		}
	}
	if !resource.Tooltip.Equals(other.Tooltip) {
		return false
	}

	return true
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

func (resource MapViewConfig) Equals(other MapViewConfig) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Lat == nil && other.Lat != nil || resource.Lat != nil && other.Lat == nil {
		return false
	}

	if resource.Lat != nil {
		if *resource.Lat != *other.Lat {
			return false
		}
	}
	if resource.Lon == nil && other.Lon != nil || resource.Lon != nil && other.Lon == nil {
		return false
	}

	if resource.Lon != nil {
		if *resource.Lon != *other.Lon {
			return false
		}
	}
	if resource.Zoom == nil && other.Zoom != nil || resource.Zoom != nil && other.Zoom == nil {
		return false
	}

	if resource.Zoom != nil {
		if *resource.Zoom != *other.Zoom {
			return false
		}
	}
	if resource.MinZoom == nil && other.MinZoom != nil || resource.MinZoom != nil && other.MinZoom == nil {
		return false
	}

	if resource.MinZoom != nil {
		if *resource.MinZoom != *other.MinZoom {
			return false
		}
	}
	if resource.MaxZoom == nil && other.MaxZoom != nil || resource.MaxZoom != nil && other.MaxZoom == nil {
		return false
	}

	if resource.MaxZoom != nil {
		if *resource.MaxZoom != *other.MaxZoom {
			return false
		}
	}
	if resource.Padding == nil && other.Padding != nil || resource.Padding != nil && other.Padding == nil {
		return false
	}

	if resource.Padding != nil {
		if *resource.Padding != *other.Padding {
			return false
		}
	}
	if resource.AllLayers == nil && other.AllLayers != nil || resource.AllLayers != nil && other.AllLayers == nil {
		return false
	}

	if resource.AllLayers != nil {
		if *resource.AllLayers != *other.AllLayers {
			return false
		}
	}
	if resource.LastOnly == nil && other.LastOnly != nil || resource.LastOnly != nil && other.LastOnly == nil {
		return false
	}

	if resource.LastOnly != nil {
		if *resource.LastOnly != *other.LastOnly {
			return false
		}
	}
	if resource.Layer == nil && other.Layer != nil || resource.Layer != nil && other.Layer == nil {
		return false
	}

	if resource.Layer != nil {
		if *resource.Layer != *other.Layer {
			return false
		}
	}
	if resource.Shared == nil && other.Shared != nil || resource.Shared != nil && other.Shared == nil {
		return false
	}

	if resource.Shared != nil {
		if *resource.Shared != *other.Shared {
			return false
		}
	}

	return true
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

func (resource ControlsOptions) Equals(other ControlsOptions) bool {
	if resource.ShowZoom == nil && other.ShowZoom != nil || resource.ShowZoom != nil && other.ShowZoom == nil {
		return false
	}

	if resource.ShowZoom != nil {
		if *resource.ShowZoom != *other.ShowZoom {
			return false
		}
	}
	if resource.MouseWheelZoom == nil && other.MouseWheelZoom != nil || resource.MouseWheelZoom != nil && other.MouseWheelZoom == nil {
		return false
	}

	if resource.MouseWheelZoom != nil {
		if *resource.MouseWheelZoom != *other.MouseWheelZoom {
			return false
		}
	}
	if resource.ShowAttribution == nil && other.ShowAttribution != nil || resource.ShowAttribution != nil && other.ShowAttribution == nil {
		return false
	}

	if resource.ShowAttribution != nil {
		if *resource.ShowAttribution != *other.ShowAttribution {
			return false
		}
	}
	if resource.ShowScale == nil && other.ShowScale != nil || resource.ShowScale != nil && other.ShowScale == nil {
		return false
	}

	if resource.ShowScale != nil {
		if *resource.ShowScale != *other.ShowScale {
			return false
		}
	}
	if resource.ShowDebug == nil && other.ShowDebug != nil || resource.ShowDebug != nil && other.ShowDebug == nil {
		return false
	}

	if resource.ShowDebug != nil {
		if *resource.ShowDebug != *other.ShowDebug {
			return false
		}
	}
	if resource.ShowMeasure == nil && other.ShowMeasure != nil || resource.ShowMeasure != nil && other.ShowMeasure == nil {
		return false
	}

	if resource.ShowMeasure != nil {
		if *resource.ShowMeasure != *other.ShowMeasure {
			return false
		}
	}

	return true
}

type TooltipOptions struct {
	Mode TooltipMode `json:"mode"`
}

func (resource TooltipOptions) Equals(other TooltipOptions) bool {
	if resource.Mode != other.Mode {
		return false
	}

	return true
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

func VariantConfig() variants.PanelcfgConfig {
	return variants.PanelcfgConfig{
		Identifier: "geomap",
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
