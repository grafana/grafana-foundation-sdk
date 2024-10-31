// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';
import * as common from '../common';

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

    // Controls how the tooltip is shown
    mode(mode: common.TooltipDisplayMode): this {
        this.internal.mode = mode;
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
