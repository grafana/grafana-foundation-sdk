// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class HeatmapCalculationBucketConfigBuilder implements cog.Builder<common.HeatmapCalculationBucketConfig> {
    protected readonly internal: common.HeatmapCalculationBucketConfig;

    constructor() {
        this.internal = common.defaultHeatmapCalculationBucketConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.HeatmapCalculationBucketConfig {
        return this.internal;
    }

    // Sets the bucket calculation mode
    mode(mode: common.HeatmapCalculationMode): this {
        this.internal.mode = mode;
        return this;
    }

    // The number of buckets to use for the axis in the heatmap
    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    // Controls the scale of the buckets
    scale(scale: cog.Builder<common.ScaleDistributionConfig>): this {
        const scaleResource = scale.build();
        this.internal.scale = scaleResource;
        return this;
    }
}
