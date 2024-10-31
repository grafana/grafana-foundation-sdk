// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorFunctionParameterExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorFunctionParameterExpression> {
    protected readonly internal: cloudwatch.QueryEditorFunctionParameterExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorFunctionParameterExpression();
        this.internal.type = "functionParameter";
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorFunctionParameterExpression {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}
