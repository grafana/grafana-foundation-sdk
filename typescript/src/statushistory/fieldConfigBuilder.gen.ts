// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as statushistory from '../statushistory';
import * as common from '../common';

export class FieldConfigBuilder implements cog.Builder<statushistory.FieldConfig> {
    protected readonly internal: statushistory.FieldConfig;

    constructor() {
        this.internal = statushistory.defaultFieldConfig();
    }

    /**
     * Builds the object.
     */
    build(): statushistory.FieldConfig {
        return this.internal;
    }

    lineWidth(lineWidth: number): this {
        if (!(lineWidth <= 10)) {
            throw new Error("lineWidth must be <= 10");
        }
        this.internal.lineWidth = lineWidth;
        return this;
    }

    hideFrom(hideFrom: cog.Builder<common.HideSeriesConfig>): this {
        const hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }

    fillOpacity(fillOpacity: number): this {
        if (!(fillOpacity <= 100)) {
            throw new Error("fillOpacity must be <= 100");
        }
        this.internal.fillOpacity = fillOpacity;
        return this;
    }
}

