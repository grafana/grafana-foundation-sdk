import * as cog from '../cog';
import * as bigquery from '../bigquery';
export declare class QueryEditorGroupByExpressionBuilder implements cog.Builder<bigquery.QueryEditorGroupByExpression> {
    protected readonly internal: bigquery.QueryEditorGroupByExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorGroupByExpression;
    property(property: cog.Builder<bigquery.QueryEditorProperty>): this;
}
