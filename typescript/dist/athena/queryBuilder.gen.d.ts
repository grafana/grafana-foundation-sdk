import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as athena from '../athena';
export declare class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind;
    version(version: string): this;
    datasource(ref: {
        name?: string;
    }): this;
    format(format: athena.FormatOptions): this;
    connectionArgs(connectionArgs: cog.Builder<athena.ConnectionArgs>): this;
    table(table: string): this;
    column(column: string): this;
    queryID(queryID: string): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    rawSQL(rawSQL: string): this;
}
