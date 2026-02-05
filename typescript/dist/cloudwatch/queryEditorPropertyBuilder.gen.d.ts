import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
export declare class QueryEditorPropertyBuilder implements cog.Builder<cloudwatch.QueryEditorProperty> {
    protected readonly internal: cloudwatch.QueryEditorProperty;
    constructor();
    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorProperty;
    name(name: string): this;
}
