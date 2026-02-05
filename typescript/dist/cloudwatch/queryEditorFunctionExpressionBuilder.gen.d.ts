import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
export declare class QueryEditorFunctionExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorFunctionExpression> {
    protected readonly internal: cloudwatch.QueryEditorFunctionExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorFunctionExpression;
    name(name: string): this;
    parameters(parameters: cog.Builder<cloudwatch.QueryEditorFunctionParameterExpression>[]): this;
}
