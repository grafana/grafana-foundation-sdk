import * as cog from '../cog';
import * as common from '../common';
export declare class HeatmapCalculationOptionsBuilder implements cog.Builder<common.HeatmapCalculationOptions> {
    protected readonly internal: common.HeatmapCalculationOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.HeatmapCalculationOptions;
    xBuckets(xBuckets: cog.Builder<common.HeatmapCalculationBucketConfig>): this;
    yBuckets(yBuckets: cog.Builder<common.HeatmapCalculationBucketConfig>): this;
}
