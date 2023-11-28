// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorArrayExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorArrayExpression> {
    private readonly internal: cloudwatch.QueryEditorArrayExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorArrayExpression();
    }

    build(): cloudwatch.QueryEditorArrayExpression {
        return this.internal;
    }

    type(type: "and" | "or"): this {
        this.internal.type = type;
        return this;
    }

    expressions(expressions: cloudwatch.QueryEditorExpression[]): this {
        this.internal.expressions = expressions;
        return this;
    }
}
