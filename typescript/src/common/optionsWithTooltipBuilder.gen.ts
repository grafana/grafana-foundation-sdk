// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class OptionsWithTooltipBuilder implements cog.Builder<common.OptionsWithTooltip> {
    private readonly internal: common.OptionsWithTooltip;

    constructor() {
        this.internal = common.defaultOptionsWithTooltip();
    }

    build(): common.OptionsWithTooltip {
        return this.internal;
    }

    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this {
        const tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }
}
