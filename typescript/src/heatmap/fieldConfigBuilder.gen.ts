// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';
import * as common from '../common';

export class FieldConfigBuilder implements cog.Builder<heatmap.FieldConfig> {
    protected readonly internal: heatmap.FieldConfig;

    constructor() {
        this.internal = heatmap.defaultFieldConfig();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.FieldConfig {
        return this.internal;
    }

    scaleDistribution(scaleDistribution: cog.Builder<common.ScaleDistributionConfig>): this {
        const scaleDistributionResource = scaleDistribution.build();
        this.internal.scaleDistribution = scaleDistributionResource;
        return this;
    }

    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this {
        const hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }
}

