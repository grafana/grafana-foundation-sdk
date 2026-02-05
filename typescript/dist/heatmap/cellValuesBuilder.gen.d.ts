import * as cog from '../cog';
import * as heatmap from '../heatmap';
export declare class CellValuesBuilder implements cog.Builder<heatmap.CellValues> {
    protected readonly internal: heatmap.CellValues;
    constructor();
    /**
     * Builds the object.
     */
    build(): heatmap.CellValues;
    unit(unit: string): this;
    decimals(decimals: number): this;
}
