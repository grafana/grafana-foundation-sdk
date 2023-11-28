// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class StackableFieldConfigBuilder implements cog.OptionsBuilder<common.StackableFieldConfig> {
    private readonly internal: common.StackableFieldConfig;

    constructor() {
        this.internal = common.defaultStackableFieldConfig();
    }

    build(): common.StackableFieldConfig {
        return this.internal;
    }

    stacking(stacking: cog.OptionsBuilder<common.StackingConfig>): this {
        const stackingResource = stacking.build();
        this.internal.stacking = stackingResource;
        return this;
    }
}
