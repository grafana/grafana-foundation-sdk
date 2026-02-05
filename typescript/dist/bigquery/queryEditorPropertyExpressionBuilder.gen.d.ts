import * as cog from '../cog';
import * as bigquery from '../bigquery';
export declare class QueryEditorPropertyExpressionBuilder implements cog.Builder<bigquery.QueryEditorPropertyExpression> {
    protected readonly internal: bigquery.QueryEditorPropertyExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorPropertyExpression;
    property(property: cog.Builder<bigquery.QueryEditorProperty>): this;
}
