// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class VizTooltipOptionsBuilder implements cog.Builder<common.VizTooltipOptions> {
    private readonly internal: common.VizTooltipOptions;

    constructor() {
        this.internal = common.defaultVizTooltipOptions();
    }

    build(): common.VizTooltipOptions {
        return this.internal;
    }

    mode(mode: common.TooltipDisplayMode): this {
        this.internal.mode = mode;
        return this;
    }

    sort(sort: common.SortOrder): this {
        this.internal.sort = sort;
        return this;
    }

    maxWidth(maxWidth: number): this {
        this.internal.maxWidth = maxWidth;
        return this;
    }

    maxHeight(maxHeight: number): this {
        this.internal.maxHeight = maxHeight;
        return this;
    }
}
