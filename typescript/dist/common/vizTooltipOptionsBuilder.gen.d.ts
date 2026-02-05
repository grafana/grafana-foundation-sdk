import * as cog from '../cog';
import * as common from '../common';
export declare class VizTooltipOptionsBuilder implements cog.Builder<common.VizTooltipOptions> {
    protected readonly internal: common.VizTooltipOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.VizTooltipOptions;
    mode(mode: common.TooltipDisplayMode): this;
    sort(sort: common.SortOrder): this;
    maxWidth(maxWidth: number): this;
    maxHeight(maxHeight: number): this;
    hideZeros(hideZeros: boolean): this;
}
