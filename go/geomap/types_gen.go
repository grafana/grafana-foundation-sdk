// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

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

type Options struct {
	View     MapViewConfig            `json:"view"`
	Controls ControlsOptions          `json:"controls"`
	Basemap  common.MapLayerOptions   `json:"basemap"`
	Layers   []common.MapLayerOptions `json:"layers"`
	Tooltip  TooltipOptions           `json:"tooltip"`
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
	// Field "view"
	if fields["view"] != nil {
		if string(fields["view"]) != "null" {

			resource.View = MapViewConfig{}
			if err := resource.View.UnmarshalJSONStrict(fields["view"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("view", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("view", errors.New("required field is null"))...)

		}
		delete(fields, "view")
	} else {
		errs = append(errs, cog.MakeBuildErrors("view", errors.New("required field is missing from input"))...)
	}
	// Field "controls"
	if fields["controls"] != nil {
		if string(fields["controls"]) != "null" {

			resource.Controls = ControlsOptions{}
			if err := resource.Controls.UnmarshalJSONStrict(fields["controls"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("controls", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("controls", errors.New("required field is null"))...)

		}
		delete(fields, "controls")
	} else {
		errs = append(errs, cog.MakeBuildErrors("controls", errors.New("required field is missing from input"))...)
	}
	// Field "basemap"
	if fields["basemap"] != nil {
		if string(fields["basemap"]) != "null" {

			resource.Basemap = common.MapLayerOptions{}
			if err := resource.Basemap.UnmarshalJSONStrict(fields["basemap"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("basemap", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("basemap", errors.New("required field is null"))...)

		}
		delete(fields, "basemap")
	} else {
		errs = append(errs, cog.MakeBuildErrors("basemap", errors.New("required field is missing from input"))...)
	}
	// Field "layers"
	if fields["layers"] != nil {
		if string(fields["layers"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["layers"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 common.MapLayerOptions

				result1 = common.MapLayerOptions{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("layers["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Layers = append(resource.Layers, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("layers", errors.New("required field is null"))...)

		}
		delete(fields, "layers")
	} else {
		errs = append(errs, cog.MakeBuildErrors("layers", errors.New("required field is missing from input"))...)
	}
	// Field "tooltip"
	if fields["tooltip"] != nil {
		if string(fields["tooltip"]) != "null" {

			resource.Tooltip = TooltipOptions{}
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

// Validate checks all the validation constraints that may be defined on `Options` fields for violations and returns them.
func (resource Options) Validate() error {
	var errs cog.BuildErrors
	if err := resource.View.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("view", err)...)
	}
	if err := resource.Controls.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("controls", err)...)
	}
	if err := resource.Basemap.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("basemap", err)...)
	}

	for i1 := range resource.Layers {
		if err := resource.Layers[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("layers["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if err := resource.Tooltip.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("tooltip", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MapViewConfig` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MapViewConfig) UnmarshalJSONStrict(raw []byte) error {
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
	} else {
		errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is missing from input"))...)
	}
	// Field "lat"
	if fields["lat"] != nil {
		if string(fields["lat"]) != "null" {
			if err := json.Unmarshal(fields["lat"], &resource.Lat); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lat", err)...)
			}

		}
		delete(fields, "lat")

	}
	// Field "lon"
	if fields["lon"] != nil {
		if string(fields["lon"]) != "null" {
			if err := json.Unmarshal(fields["lon"], &resource.Lon); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lon", err)...)
			}

		}
		delete(fields, "lon")

	}
	// Field "zoom"
	if fields["zoom"] != nil {
		if string(fields["zoom"]) != "null" {
			if err := json.Unmarshal(fields["zoom"], &resource.Zoom); err != nil {
				errs = append(errs, cog.MakeBuildErrors("zoom", err)...)
			}

		}
		delete(fields, "zoom")

	}
	// Field "minZoom"
	if fields["minZoom"] != nil {
		if string(fields["minZoom"]) != "null" {
			if err := json.Unmarshal(fields["minZoom"], &resource.MinZoom); err != nil {
				errs = append(errs, cog.MakeBuildErrors("minZoom", err)...)
			}

		}
		delete(fields, "minZoom")

	}
	// Field "maxZoom"
	if fields["maxZoom"] != nil {
		if string(fields["maxZoom"]) != "null" {
			if err := json.Unmarshal(fields["maxZoom"], &resource.MaxZoom); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxZoom", err)...)
			}

		}
		delete(fields, "maxZoom")

	}
	// Field "padding"
	if fields["padding"] != nil {
		if string(fields["padding"]) != "null" {
			if err := json.Unmarshal(fields["padding"], &resource.Padding); err != nil {
				errs = append(errs, cog.MakeBuildErrors("padding", err)...)
			}

		}
		delete(fields, "padding")

	}
	// Field "allLayers"
	if fields["allLayers"] != nil {
		if string(fields["allLayers"]) != "null" {
			if err := json.Unmarshal(fields["allLayers"], &resource.AllLayers); err != nil {
				errs = append(errs, cog.MakeBuildErrors("allLayers", err)...)
			}

		}
		delete(fields, "allLayers")

	}
	// Field "lastOnly"
	if fields["lastOnly"] != nil {
		if string(fields["lastOnly"]) != "null" {
			if err := json.Unmarshal(fields["lastOnly"], &resource.LastOnly); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lastOnly", err)...)
			}

		}
		delete(fields, "lastOnly")

	}
	// Field "layer"
	if fields["layer"] != nil {
		if string(fields["layer"]) != "null" {
			if err := json.Unmarshal(fields["layer"], &resource.Layer); err != nil {
				errs = append(errs, cog.MakeBuildErrors("layer", err)...)
			}

		}
		delete(fields, "layer")

	}
	// Field "shared"
	if fields["shared"] != nil {
		if string(fields["shared"]) != "null" {
			if err := json.Unmarshal(fields["shared"], &resource.Shared); err != nil {
				errs = append(errs, cog.MakeBuildErrors("shared", err)...)
			}

		}
		delete(fields, "shared")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MapViewConfig", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MapViewConfig` objects.
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

// Validate checks all the validation constraints that may be defined on `MapViewConfig` fields for violations and returns them.
func (resource MapViewConfig) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ControlsOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ControlsOptions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "showZoom"
	if fields["showZoom"] != nil {
		if string(fields["showZoom"]) != "null" {
			if err := json.Unmarshal(fields["showZoom"], &resource.ShowZoom); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showZoom", err)...)
			}

		}
		delete(fields, "showZoom")

	}
	// Field "mouseWheelZoom"
	if fields["mouseWheelZoom"] != nil {
		if string(fields["mouseWheelZoom"]) != "null" {
			if err := json.Unmarshal(fields["mouseWheelZoom"], &resource.MouseWheelZoom); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mouseWheelZoom", err)...)
			}

		}
		delete(fields, "mouseWheelZoom")

	}
	// Field "showAttribution"
	if fields["showAttribution"] != nil {
		if string(fields["showAttribution"]) != "null" {
			if err := json.Unmarshal(fields["showAttribution"], &resource.ShowAttribution); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showAttribution", err)...)
			}

		}
		delete(fields, "showAttribution")

	}
	// Field "showScale"
	if fields["showScale"] != nil {
		if string(fields["showScale"]) != "null" {
			if err := json.Unmarshal(fields["showScale"], &resource.ShowScale); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showScale", err)...)
			}

		}
		delete(fields, "showScale")

	}
	// Field "showDebug"
	if fields["showDebug"] != nil {
		if string(fields["showDebug"]) != "null" {
			if err := json.Unmarshal(fields["showDebug"], &resource.ShowDebug); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showDebug", err)...)
			}

		}
		delete(fields, "showDebug")

	}
	// Field "showMeasure"
	if fields["showMeasure"] != nil {
		if string(fields["showMeasure"]) != "null" {
			if err := json.Unmarshal(fields["showMeasure"], &resource.ShowMeasure); err != nil {
				errs = append(errs, cog.MakeBuildErrors("showMeasure", err)...)
			}

		}
		delete(fields, "showMeasure")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ControlsOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ControlsOptions` objects.
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

// Validate checks all the validation constraints that may be defined on `ControlsOptions` fields for violations and returns them.
func (resource ControlsOptions) Validate() error {
	return nil
}

type TooltipOptions struct {
	Mode TooltipMode `json:"mode"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TooltipOptions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TooltipOptions) UnmarshalJSONStrict(raw []byte) error {
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

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TooltipOptions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TooltipOptions` objects.
func (resource TooltipOptions) Equals(other TooltipOptions) bool {
	if resource.Mode != other.Mode {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `TooltipOptions` fields for violations and returns them.
func (resource TooltipOptions) Validate() error {
	return nil
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

// VariantConfig returns the configuration related to geomap panels.
// This configuration describes how to unmarshal it, convert it to code, …
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
