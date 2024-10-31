// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
export class QueryEditorOperatorBuilder implements cog.Builder<cloudwatch.QueryEditorOperator> {
    protected readonly internal: cloudwatch.QueryEditorOperator;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorOperator();
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorOperator {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    value(value: cloudwatch.QueryEditorOperatorType | cloudwatch.QueryEditorOperatorType[]): this {
        this.internal.value = value;
        return this;
    }
}
