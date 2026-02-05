import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class RowsLayoutBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutKind> {
    protected readonly internal: dashboardv2beta1.RowsLayoutKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowsLayoutKind;
    rows(rows: cog.Builder<dashboardv2beta1.RowsLayoutRowKind>[]): this;
    row(row: cog.Builder<dashboardv2beta1.RowsLayoutRowKind>): this;
}
