import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as bigquery from '../bigquery';
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
}
