// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';

// Controls tooltip options
export class HeatmapTooltipBuilder implements cog.Builder<heatmap.HeatmapTooltip> {
    protected readonly internal: heatmap.HeatmapTooltip;

    constructor() {
        this.internal = heatmap.defaultHeatmapTooltip();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.HeatmapTooltip {
        return this.internal;
    }

    // Controls if the tooltip is shown
    show(show: boolean): this {
        this.internal.show = show;
        return this;
    }

    // Controls if the tooltip shows a histogram of the y-axis values
    yHistogram(yHistogram: boolean): this {
        this.internal.yHistogram = yHistogram;
        return this;
    }

    // Controls if the tooltip shows a color scale in header
    showColorScale(showColorScale: boolean): this {
        this.internal.showColorScale = showColorScale;
        return this;
    }
}
