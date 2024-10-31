// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class ColorDimensionConfigBuilder implements cog.Builder<common.ColorDimensionConfig> {
    protected readonly internal: common.ColorDimensionConfig;

    constructor() {
        this.internal = common.defaultColorDimensionConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.ColorDimensionConfig {
        return this.internal;
    }

    // color value
    fixed(fixed: string): this {
        this.internal.fixed = fixed;
        return this;
    }

    // fixed: T -- will be added by each element
    field(field: string): this {
        this.internal.field = field;
        return this;
    }
}
