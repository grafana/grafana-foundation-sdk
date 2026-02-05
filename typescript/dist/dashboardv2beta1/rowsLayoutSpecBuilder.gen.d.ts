import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class RowsLayoutSpecBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutSpec> {
    protected readonly internal: dashboardv2beta1.RowsLayoutSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowsLayoutSpec;
    rows(rows: cog.Builder<dashboardv2beta1.RowsLayoutRowKind>[]): this;
}
