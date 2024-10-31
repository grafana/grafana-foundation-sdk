// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class TextDimensionConfigBuilder implements cog.Builder<common.TextDimensionConfig> {
    protected readonly internal: common.TextDimensionConfig;

    constructor() {
        this.internal = common.defaultTextDimensionConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.TextDimensionConfig {
        return this.internal;
    }

    mode(mode: common.TextDimensionMode): this {
        this.internal.mode = mode;
        return this;
    }

    // fixed: T -- will be added by each element
    field(field: string): this {
        this.internal.field = field;
        return this;
    }

    fixed(fixed: string): this {
        this.internal.fixed = fixed;
        return this;
    }
}
