// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Links to a resource (image/svg path)
export class ResourceDimensionConfigBuilder implements cog.Builder<common.ResourceDimensionConfig> {
    protected readonly internal: common.ResourceDimensionConfig;

    constructor() {
        this.internal = common.defaultResourceDimensionConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.ResourceDimensionConfig {
        return this.internal;
    }

    mode(mode: common.ResourceDimensionMode): this {
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
