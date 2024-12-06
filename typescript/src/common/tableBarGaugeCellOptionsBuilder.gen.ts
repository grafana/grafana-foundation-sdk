// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// Gauge cell options
export class TableBarGaugeCellOptionsBuilder implements cog.Builder<common.TableBarGaugeCellOptions> {
    protected readonly internal: common.TableBarGaugeCellOptions;

    constructor() {
        this.internal = common.defaultTableBarGaugeCellOptions();
        this.internal.type = "gauge";
    }

    /**
     * Builds the object.
     */
    build(): common.TableBarGaugeCellOptions {
        return this.internal;
    }

    mode(mode: common.BarGaugeDisplayMode): this {
        this.internal.mode = mode;
        return this;
    }

    valueDisplayMode(valueDisplayMode: common.BarGaugeValueMode): this {
        this.internal.valueDisplayMode = valueDisplayMode;
        return this;
    }
}
