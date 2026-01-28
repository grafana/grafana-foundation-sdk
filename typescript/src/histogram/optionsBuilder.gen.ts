// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as histogram from '../histogram';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<histogram.Options> {
    protected readonly internal: histogram.Options;

    constructor() {
        this.internal = histogram.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): histogram.Options {
        return this.internal;
    }

    // Bucket count (approx)
    bucketCount(bucketCount: number): this {
        if (!(bucketCount > 0)) {
            throw new Error("bucketCount must be > 0");
        }
        this.internal.bucketCount = bucketCount;
        return this;
    }

    // Size of each bucket
    bucketSize(bucketSize: number): this {
        this.internal.bucketSize = bucketSize;
        return this;
    }

    // Offset buckets by this amount
    bucketOffset(bucketOffset: number): this {
        this.internal.bucketOffset = bucketOffset;
        return this;
    }

    legend(legend: cog.Builder<common.VizLegendOptions>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }

    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this {
        const tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }

    // Combines multiple series into a single histogram
    combine(combine: boolean): this {
        this.internal.combine = combine;
        return this;
    }
}

