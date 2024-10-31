// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';

// Controls legend options
export class HeatmapLegendBuilder implements cog.Builder<heatmap.HeatmapLegend> {
    protected readonly internal: heatmap.HeatmapLegend;

    constructor() {
        this.internal = heatmap.defaultHeatmapLegend();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.HeatmapLegend {
        return this.internal;
    }

    // Controls if the legend is shown
    show(show: boolean): this {
        this.internal.show = show;
        return this;
    }
}
