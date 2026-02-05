import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
export declare class QueryEditorOperatorBuilder implements cog.Builder<cloudwatch.QueryEditorOperator> {
    protected readonly internal: cloudwatch.QueryEditorOperator;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorOperator;
    name(name: string): this;
    value(value: cloudwatch.QueryEditorOperatorType | cloudwatch.QueryEditorOperatorType[]): this;
}
