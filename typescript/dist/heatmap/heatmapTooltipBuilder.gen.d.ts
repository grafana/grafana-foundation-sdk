import * as cog from '../cog';
import * as heatmap from '../heatmap';
import * as common from '../common';
export declare class HeatmapTooltipBuilder implements cog.Builder<heatmap.HeatmapTooltip> {
    protected readonly internal: heatmap.HeatmapTooltip;
    constructor();
    /**
     * Builds the object.
     */
    build(): heatmap.HeatmapTooltip;
    mode(mode: common.TooltipDisplayMode): this;
    maxHeight(maxHeight: number): this;
    maxWidth(maxWidth: number): this;
    yHistogram(yHistogram: boolean): this;
    showColorScale(showColorScale: boolean): this;
}
