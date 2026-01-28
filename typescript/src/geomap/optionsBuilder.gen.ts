// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as geomap from '../geomap';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<geomap.Options> {
    protected readonly internal: geomap.Options;

    constructor() {
        this.internal = geomap.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): geomap.Options {
        return this.internal;
    }

    view(view: cog.Builder<geomap.MapViewConfig>): this {
        const viewResource = view.build();
        this.internal.view = viewResource;
        return this;
    }

    controls(controls: cog.Builder<geomap.ControlsOptions>): this {
        const controlsResource = controls.build();
        this.internal.controls = controlsResource;
        return this;
    }

    basemap(basemap: cog.Builder<common.MapLayerOptions>): this {
        const basemapResource = basemap.build();
        this.internal.basemap = basemapResource;
        return this;
    }

    layers(layers: cog.Builder<common.MapLayerOptions>[]): this {
        const layersResources = layers.map(builder1 => builder1.build());
        this.internal.layers = layersResources;
        return this;
    }

    tooltip(tooltip: cog.Builder<geomap.TooltipOptions>): this {
        const tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }
}

