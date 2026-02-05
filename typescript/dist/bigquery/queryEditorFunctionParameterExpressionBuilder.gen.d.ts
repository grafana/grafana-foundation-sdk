import * as cog from '../cog';
import * as bigquery from '../bigquery';
export declare class QueryEditorFunctionParameterExpressionBuilder implements cog.Builder<bigquery.QueryEditorFunctionParameterExpression> {
    protected readonly internal: bigquery.QueryEditorFunctionParameterExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorFunctionParameterExpression;
    name(name: string): this;
}
