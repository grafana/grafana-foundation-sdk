import * as cog from '../cog';
import * as bigquery from '../bigquery';
export declare class SQLExpressionBuilder implements cog.Builder<bigquery.SQLExpression> {
    protected readonly internal: bigquery.SQLExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): bigquery.SQLExpression;
    columns(columns: cog.Builder<bigquery.QueryEditorFunctionExpression>[]): this;
    from(from: string): this;
    whereString(whereString: string): this;
    groupBy(groupBy: cog.Builder<bigquery.QueryEditorGroupByExpression>[]): this;
    orderBy(orderBy: cog.Builder<bigquery.QueryEditorPropertyExpression>): this;
    orderByDirection(orderByDirection: bigquery.OrderByDirection): this;
    limit(limit: number): this;
    offset(offset: number): this;
}
