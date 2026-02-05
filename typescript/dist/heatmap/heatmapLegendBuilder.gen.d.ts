import * as cog from '../cog';
import * as heatmap from '../heatmap';
export declare class HeatmapLegendBuilder implements cog.Builder<heatmap.HeatmapLegend> {
    protected readonly internal: heatmap.HeatmapLegend;
    constructor();
    /**
     * Builds the object.
     */
    build(): heatmap.HeatmapLegend;
    show(show: boolean): this;
}
