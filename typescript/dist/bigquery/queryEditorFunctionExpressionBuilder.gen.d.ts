import * as cog from '../cog';
import * as bigquery from '../bigquery';
export declare class QueryEditorFunctionExpressionBuilder implements cog.Builder<bigquery.QueryEditorFunctionExpression> {
    protected readonly internal: bigquery.QueryEditorFunctionExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): bigquery.QueryEditorFunctionExpression;
    name(name: string): this;
    parameters(parameters: cog.Builder<bigquery.QueryEditorFunctionParameterExpression>[]): this;
}
