// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class FillConfigBuilder implements cog.Builder<common.FillConfig> {
    protected readonly internal: common.FillConfig;

    constructor() {
        this.internal = common.defaultFillConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.FillConfig {
        return this.internal;
    }

    fillColor(fillColor: string): this {
        this.internal.fillColor = fillColor;
        return this;
    }

    fillOpacity(fillOpacity: number): this {
        this.internal.fillOpacity = fillOpacity;
        return this;
    }

    fillBelowTo(fillBelowTo: string): this {
        this.internal.fillBelowTo = fillBelowTo;
        return this;
    }
}
