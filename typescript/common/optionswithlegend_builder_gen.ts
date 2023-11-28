// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class OptionsWithLegendBuilder implements cog.OptionsBuilder<common.OptionsWithLegend> {
    private readonly internal: common.OptionsWithLegend;

    constructor() {
        this.internal = common.defaultOptionsWithLegend();
    }

    build(): common.OptionsWithLegend {
        return this.internal;
    }

    legend(legend: cog.OptionsBuilder<common.VizLegendOptions>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }
}
