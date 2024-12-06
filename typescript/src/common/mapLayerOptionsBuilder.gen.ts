// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class MapLayerOptionsBuilder implements cog.Builder<common.MapLayerOptions> {
    protected readonly internal: common.MapLayerOptions;

    constructor() {
        this.internal = common.defaultMapLayerOptions();
    }

    /**
     * Builds the object.
     */
    build(): common.MapLayerOptions {
        return this.internal;
    }

    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    // configured unique display name
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // Custom options depending on the type
    config(config: any): this {
        this.internal.config = config;
        return this;
    }

    // Common method to define geometry fields
    location(location: cog.Builder<common.FrameGeometrySource>): this {
        const locationResource = location.build();
        this.internal.location = locationResource;
        return this;
    }

    // Defines a frame MatcherConfig that may filter data for the given layer
    filterData(filterData: any): this {
        this.internal.filterData = filterData;
        return this;
    }

    // Common properties:
    // https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
    // Layer opacity (0-1)
    opacity(opacity: number): this {
        this.internal.opacity = opacity;
        return this;
    }

    // Check tooltip (defaults to true)
    tooltip(tooltip: boolean): this {
        this.internal.tooltip = tooltip;
        return this;
    }
}
