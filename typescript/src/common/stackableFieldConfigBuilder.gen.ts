// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class StackableFieldConfigBuilder implements cog.Builder<common.StackableFieldConfig> {
    protected readonly internal: common.StackableFieldConfig;

    constructor() {
        this.internal = common.defaultStackableFieldConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.StackableFieldConfig {
        return this.internal;
    }

    stacking(stacking: cog.Builder<common.StackingConfig>): this {
        const stackingResource = stacking.build();
        this.internal.stacking = stackingResource;
        return this;
    }
}
