// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class VizLegendOptionsBuilder implements cog.Builder<common.VizLegendOptions> {
    protected readonly internal: common.VizLegendOptions;

    constructor() {
        this.internal = common.defaultVizLegendOptions();
    }

    /**
     * Builds the object.
     */
    build(): common.VizLegendOptions {
        return this.internal;
    }

    displayMode(displayMode: common.LegendDisplayMode): this {
        this.internal.displayMode = displayMode;
        return this;
    }

    placement(placement: common.LegendPlacement): this {
        this.internal.placement = placement;
        return this;
    }

    showLegend(showLegend: boolean): this {
        this.internal.showLegend = showLegend;
        return this;
    }

    asTable(asTable: boolean): this {
        this.internal.asTable = asTable;
        return this;
    }

    isVisible(isVisible: boolean): this {
        this.internal.isVisible = isVisible;
        return this;
    }

    sortBy(sortBy: string): this {
        this.internal.sortBy = sortBy;
        return this;
    }

    sortDesc(sortDesc: boolean): this {
        this.internal.sortDesc = sortDesc;
        return this;
    }

    width(width: number): this {
        this.internal.width = width;
        return this;
    }

    calcs(calcs: string[]): this {
        this.internal.calcs = calcs;
        return this;
    }
}
