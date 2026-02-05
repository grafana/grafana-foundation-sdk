import * as common from '../common';
export interface Options {
    view: MapViewConfig;
    controls: ControlsOptions;
    basemap: common.MapLayerOptions;
    layers: common.MapLayerOptions[];
    tooltip: TooltipOptions;
}
export declare const defaultOptions: () => Options;
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
export declare const defaultMapViewConfig: () => MapViewConfig;
export interface ControlsOptions {
    showZoom?: boolean;
    mouseWheelZoom?: boolean;
    showAttribution?: boolean;
    showScale?: boolean;
    showDebug?: boolean;
    showMeasure?: boolean;
}
export declare const defaultControlsOptions: () => ControlsOptions;
export interface TooltipOptions {
    mode: TooltipMode;
}
export declare const defaultTooltipOptions: () => TooltipOptions;
export declare enum TooltipMode {
    None = "none",
    Details = "details"
}
export declare const defaultTooltipMode: () => TooltipMode;
export declare enum MapCenterID {
    Zero = "zero",
    Coords = "coords",
    Fit = "fit"
}
export declare const defaultMapCenterID: () => MapCenterID;
