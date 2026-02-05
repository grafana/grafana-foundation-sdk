import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
export declare class QueryEditorFunctionParameterExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorFunctionParameterExpression> {
    protected readonly internal: cloudwatch.QueryEditorFunctionParameterExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorFunctionParameterExpression;
    name(name: string): this;
}
