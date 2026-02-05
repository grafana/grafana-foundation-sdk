import * as cog from '../cog';
import * as athena from '../athena';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: athena.Dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): athena.Dataquery;
    format(format: athena.FormatOptions): this;
    connectionArgs(connectionArgs: cog.Builder<athena.ConnectionArgs>): this;
    table(table: string): this;
    column(column: string): this;
    queryID(queryID: string): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    rawSQL(rawSQL: string): this;
    datasource(datasource: common.DataSourceRef): this;
}
