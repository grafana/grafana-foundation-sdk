import * as cog from '../cog';
import * as heatmap from '../heatmap';
export declare class FilterValueRangeBuilder implements cog.Builder<heatmap.FilterValueRange> {
    protected readonly internal: heatmap.FilterValueRange;
    constructor();
    /**
     * Builds the object.
     */
    build(): heatmap.FilterValueRange;
    le(le: number): this;
    ge(ge: number): this;
}
