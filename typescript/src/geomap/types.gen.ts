// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as common from '../common';


export interface Options {
	view: MapViewConfig;
	controls: ControlsOptions;
	basemap: common.MapLayerOptions;
	layers: common.MapLayerOptions[];
	tooltip: TooltipOptions;
}

export const defaultOptions = (): Options => ({
	view: defaultMapViewConfig(),
	controls: defaultControlsOptions(),
	basemap: common.defaultMapLayerOptions(),
	layers: [],
	tooltip: defaultTooltipOptions(),
});

export interface MapViewConfig {
	id: string;
	lat?: number;
	lon?: number;
	zoom?: number;
	minZoom?: number;
	maxZoom?: number;
	padding?: number;
	allLayers?: boolean;
	lastOnly?: boolean;
	layer?: string;
	shared?: boolean;
}

export const defaultMapViewConfig = (): MapViewConfig => ({
	id: "zero",
	lat: 0,
	lon: 0,
	zoom: 1,
	allLayers: true,
});

export interface ControlsOptions {
	// Zoom (upper left)
	showZoom?: boolean;
	// let the mouse wheel zoom
	mouseWheelZoom?: boolean;
	// Lower right
	showAttribution?: boolean;
	// Scale options
	showScale?: boolean;
	// Show debug
	showDebug?: boolean;
	// Show measure
	showMeasure?: boolean;
}

export const defaultControlsOptions = (): ControlsOptions => ({
});

export interface TooltipOptions {
	mode: TooltipMode;
}

export const defaultTooltipOptions = (): TooltipOptions => ({
	mode: TooltipMode.None,
});

export enum TooltipMode {
	None = "none",
	Details = "details",
}

export const defaultTooltipMode = (): TooltipMode => (TooltipMode.None);

export enum MapCenterID {
	Zero = "zero",
	Coords = "coords",
	Fit = "fit",
}

export const defaultMapCenterID = (): MapCenterID => (MapCenterID.Zero);

