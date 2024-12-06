// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class HeatmapCalculationOptionsBuilder implements cog.Builder<common.HeatmapCalculationOptions> {
    protected readonly internal: common.HeatmapCalculationOptions;

    constructor() {
        this.internal = common.defaultHeatmapCalculationOptions();
    }

    /**
     * Builds the object.
     */
    build(): common.HeatmapCalculationOptions {
        return this.internal;
    }

    // The number of buckets to use for the xAxis in the heatmap
    xBuckets(xBuckets: cog.Builder<common.HeatmapCalculationBucketConfig>): this {
        const xBucketsResource = xBuckets.build();
        this.internal.xBuckets = xBucketsResource;
        return this;
    }

    // The number of buckets to use for the yAxis in the heatmap
    yBuckets(yBuckets: cog.Builder<common.HeatmapCalculationBucketConfig>): this {
        const yBucketsResource = yBuckets.build();
        this.internal.yBuckets = yBucketsResource;
        return this;
    }
}
