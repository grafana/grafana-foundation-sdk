// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorWhereExpressionBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorWhereExpression> {
    protected readonly internal: azuremonitor.BuilderQueryEditorWhereExpression;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorWhereExpression();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorWhereExpression {
        return this.internal;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }

    expressions(expressions: cog.Builder<azuremonitor.BuilderQueryEditorWhereExpressionItems>[]): this {
        const expressionsResources = expressions.map(builder1 => builder1.build());
        this.internal.expressions = expressionsResources;
        return this;
    }
}

