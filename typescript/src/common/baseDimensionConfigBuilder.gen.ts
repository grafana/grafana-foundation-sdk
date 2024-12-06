// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

export class BaseDimensionConfigBuilder implements cog.Builder<common.BaseDimensionConfig> {
    protected readonly internal: common.BaseDimensionConfig;

    constructor() {
        this.internal = common.defaultBaseDimensionConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.BaseDimensionConfig {
        return this.internal;
    }

    // fixed: T -- will be added by each element
    field(field: string): this {
        this.internal.field = field;
        return this;
    }
}
