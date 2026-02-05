import * as cog from '../cog';
import * as common from '../common';
export declare class VizLegendOptionsBuilder implements cog.Builder<common.VizLegendOptions> {
    protected readonly internal: common.VizLegendOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.VizLegendOptions;
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
