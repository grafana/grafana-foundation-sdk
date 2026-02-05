import * as cog from '../cog';
import * as bigquery from '../bigquery';
import * as common from '../common';
export declare class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: bigquery.Dataquery;
    constructor();
    /**
     * Builds the object.
     */
    build(): bigquery.Dataquery;
    dataset(dataset: string): this;
    table(table: string): this;
    project(project: string): this;
    format(format: bigquery.QueryFormat): this;
    rawQuery(rawQuery: boolean): this;
    rawSql(rawSql: string): this;
    location(location: string): this;
    partitioned(partitioned: boolean): this;
    partitionedField(partitionedField: string): this;
    convertToUTC(convertToUTC: boolean): this;
    sharded(sharded: boolean): this;
    queryPriority(queryPriority: bigquery.QueryPriority): this;
    timeShift(timeShift: string): this;
    editorMode(editorMode: bigquery.EditorMode): this;
    sql(sql: cog.Builder<bigquery.SQLExpression>): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    datasource(datasource: common.DataSourceRef): this;
}
