// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as geomap from '../geomap';

export class ControlsOptionsBuilder implements cog.Builder<geomap.ControlsOptions> {
    protected readonly internal: geomap.ControlsOptions;

    constructor() {
        this.internal = geomap.defaultControlsOptions();
    }

    /**
     * Builds the object.
     */
    build(): geomap.ControlsOptions {
        return this.internal;
    }

    // Zoom (upper left)
    showZoom(showZoom: boolean): this {
        this.internal.showZoom = showZoom;
        return this;
    }

    // let the mouse wheel zoom
    mouseWheelZoom(mouseWheelZoom: boolean): this {
        this.internal.mouseWheelZoom = mouseWheelZoom;
        return this;
    }

    // Lower right
    showAttribution(showAttribution: boolean): this {
        this.internal.showAttribution = showAttribution;
        return this;
    }

    // Scale options
    showScale(showScale: boolean): this {
        this.internal.showScale = showScale;
        return this;
    }

    // Show debug
    showDebug(showDebug: boolean): this {
        this.internal.showDebug = showDebug;
        return this;
    }

    // Show measure
    showMeasure(showMeasure: boolean): this {
        this.internal.showMeasure = showMeasure;
        return this;
    }
}
