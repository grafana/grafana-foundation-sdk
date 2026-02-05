import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
export declare class QueryEditorArrayExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorArrayExpression> {
    protected readonly internal: cloudwatch.QueryEditorArrayExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorArrayExpression;
    type(type: "and" | "or"): this;
    expressions(expressions: cog.Builder<cloudwatch.QueryEditorExpression>[]): this;
}
