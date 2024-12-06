// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class ScaleDistributionConfigBuilder implements cog.Builder<common.ScaleDistributionConfig> {
    protected readonly internal: common.ScaleDistributionConfig;

    constructor() {
        this.internal = common.defaultScaleDistributionConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.ScaleDistributionConfig {
        return this.internal;
    }

    type(type: common.ScaleDistribution): this {
        this.internal.type = type;
        return this;
    }

    log(log: number): this {
        this.internal.log = log;
        return this;
    }

    linearThreshold(linearThreshold: number): this {
        this.internal.linearThreshold = linearThreshold;
        return this;
    }
}
