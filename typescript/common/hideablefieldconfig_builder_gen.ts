// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class HideableFieldConfigBuilder implements cog.OptionsBuilder<common.HideableFieldConfig> {
    private readonly internal: common.HideableFieldConfig;

    constructor() {
        this.internal = common.defaultHideableFieldConfig();
    }

    build(): common.HideableFieldConfig {
        return this.internal;
    }

    hideFrom(hideFrom: cog.OptionsBuilder<common.HideSeriesConfig>): this {
        const hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }
}
