import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
export declare class QueryEditorOperatorExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorOperatorExpression> {
    protected readonly internal: cloudwatch.QueryEditorOperatorExpression;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorOperatorExpression;
    property(property: cog.Builder<cloudwatch.QueryEditorProperty>): this;
    operator(operator: cog.Builder<cloudwatch.QueryEditorOperator>): this;
}
