// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class ScalarDimensionConfigBuilder implements cog.Builder<common.ScalarDimensionConfig> {
    protected readonly internal: common.ScalarDimensionConfig;

    constructor() {
        this.internal = common.defaultScalarDimensionConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.ScalarDimensionConfig {
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

    mode(mode: common.ScalarDimensionMode): this {
        this.internal.mode = mode;
        return this;
    }
}
