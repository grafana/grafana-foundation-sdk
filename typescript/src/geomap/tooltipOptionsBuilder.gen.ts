// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as geomap from '../geomap';

export class TooltipOptionsBuilder implements cog.Builder<geomap.TooltipOptions> {
    protected readonly internal: geomap.TooltipOptions;

    constructor() {
        this.internal = geomap.defaultTooltipOptions();
    }

    /**
     * Builds the object.
     */
    build(): geomap.TooltipOptions {
        return this.internal;
    }

    mode(mode: geomap.TooltipMode): this {
        this.internal.mode = mode;
        return this;
    }
}
