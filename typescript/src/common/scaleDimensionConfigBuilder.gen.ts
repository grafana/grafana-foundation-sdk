// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class ScaleDimensionConfigBuilder implements cog.Builder<common.ScaleDimensionConfig> {
    protected readonly internal: common.ScaleDimensionConfig;

    constructor() {
        this.internal = common.defaultScaleDimensionConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.ScaleDimensionConfig {
        return this.internal;
    }

    min(min: number): this {
        this.internal.min = min;
        return this;
    }

    max(max: number): this {
        this.internal.max = max;
        return this;
    }

    fixed(fixed: number): this {
        this.internal.fixed = fixed;
        return this;
    }

    // fixed: T -- will be added by each element
    field(field: string): this {
        this.internal.field = field;
        return this;
    }

    // | *"linear"
    mode(mode: common.ScaleDimensionMode): this {
        this.internal.mode = mode;
        return this;
    }
}
