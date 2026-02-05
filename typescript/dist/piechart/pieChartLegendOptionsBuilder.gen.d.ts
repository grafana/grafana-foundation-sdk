import * as cog from '../cog';
import * as piechart from '../piechart';
import * as common from '../common';
export declare class PieChartLegendOptionsBuilder implements cog.Builder<piechart.PieChartLegendOptions> {
    protected readonly internal: piechart.PieChartLegendOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): piechart.PieChartLegendOptions;
    values(values: piechart.PieChartLegendValues[]): this;
    displayMode(displayMode: common.LegendDisplayMode): this;
    placement(placement: common.LegendPlacement): this;
    showLegend(showLegend: boolean): this;
    asTable(asTable: boolean): this;
    isVisible(isVisible: boolean): this;
    sortBy(sortBy: string): this;
    sortDesc(sortDesc: boolean): this;
    width(width: number): this;
    calcs(calcs: string[]): this;
}
