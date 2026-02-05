import * as cog from '../cog';
import * as heatmap from '../heatmap';
import * as common from '../common';
export declare class RowsHeatmapOptionsBuilder implements cog.Builder<heatmap.RowsHeatmapOptions> {
    protected readonly internal: heatmap.RowsHeatmapOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): heatmap.RowsHeatmapOptions;
    value(value: string): this;
    layout(layout: common.HeatmapCellLayout): this;
}
