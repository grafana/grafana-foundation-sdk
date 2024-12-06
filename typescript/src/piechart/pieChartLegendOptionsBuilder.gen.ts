// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as piechart from '../piechart';
import * as common from '../common';

export class PieChartLegendOptionsBuilder implements cog.Builder<piechart.PieChartLegendOptions> {
    protected readonly internal: piechart.PieChartLegendOptions;

    constructor() {
        this.internal = piechart.defaultPieChartLegendOptions();
    }

    /**
     * Builds the object.
     */
    build(): piechart.PieChartLegendOptions {
        return this.internal;
    }

    values(values: piechart.PieChartLegendValues[]): this {
        this.internal.values = values;
        return this;
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
