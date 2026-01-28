// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as piechart from '../piechart';
import * as common from '../common';

export class FieldConfigBuilder implements cog.Builder<piechart.FieldConfig> {
    protected readonly internal: piechart.FieldConfig;

    constructor() {
        this.internal = piechart.defaultFieldConfig();
    }

    /**
     * Builds the object.
     */
    build(): piechart.FieldConfig {
        return this.internal;
    }

    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this {
        const hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }
}

