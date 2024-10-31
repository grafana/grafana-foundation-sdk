// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class OptionsWithLegendBuilder implements cog.Builder<common.OptionsWithLegend> {
    protected readonly internal: common.OptionsWithLegend;

    constructor() {
        this.internal = common.defaultOptionsWithLegend();
    }

    /**
     * Builds the object.
     */
    build(): common.OptionsWithLegend {
        return this.internal;
    }

    legend(legend: cog.Builder<common.VizLegendOptions>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }
}
