// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class QueryEditorArrayExpressionBuilder implements cog.Builder<cloudwatch.QueryEditorArrayExpression> {
    protected readonly internal: cloudwatch.QueryEditorArrayExpression;

    constructor() {
        this.internal = cloudwatch.defaultQueryEditorArrayExpression();
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.QueryEditorArrayExpression {
        return this.internal;
    }

    type(type: "and" | "or"): this {
        this.internal.type = type;
        return this;
    }

    expressions(expressions: cog.Builder<cloudwatch.QueryEditorExpression>[]): this {
        const expressionsResources = expressions.map(builder1 => builder1.build());
        this.internal.expressions = expressionsResources;
        return this;
    }
}
