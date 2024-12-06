// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class BarConfigBuilder implements cog.Builder<common.BarConfig> {
    protected readonly internal: common.BarConfig;

    constructor() {
        this.internal = common.defaultBarConfig();
    }

    /**
     * Builds the object.
     */
    build(): common.BarConfig {
        return this.internal;
    }

    barAlignment(barAlignment: common.BarAlignment): this {
        this.internal.barAlignment = barAlignment;
        return this;
    }

    barWidthFactor(barWidthFactor: number): this {
        this.internal.barWidthFactor = barWidthFactor;
        return this;
    }

    barMaxWidth(barMaxWidth: number): this {
        this.internal.barMaxWidth = barMaxWidth;
        return this;
    }
}
