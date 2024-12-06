// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorFunctionExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorFunctionExpression> {
    protected readonly internal: cloudwatch.QueryEditorFunctionExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorFunctionExpression();
        this.internal.type = "function";
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorFunctionExpression {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    parameters(parameters: cog.Builder<cloudwatch.QueryEditorFunctionParameterExpression>[]): this {
        const parametersResources = parameters.map(builder1 => builder1.build());
        this.internal.parameters = parametersResources;
        return this;
    }
}
