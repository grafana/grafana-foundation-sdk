// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorWhereExpressionArrayBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorWhereExpressionArray> {
    protected readonly internal: azuremonitor.BuilderQueryEditorWhereExpressionArray;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorWhereExpressionArray();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorWhereExpressionArray {
        return this.internal;
    }

    expressions(expressions: cog.Builder<azuremonitor.BuilderQueryEditorWhereExpression>[]): this {
        const expressionsResources = expressions.map(builder1 => builder1.build());
        this.internal.expressions = expressionsResources;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

