// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorOrderByExpressionArrayBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorOrderByExpressionArray> {
    protected readonly internal: azuremonitor.BuilderQueryEditorOrderByExpressionArray;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorOrderByExpressionArray();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorOrderByExpressionArray {
        return this.internal;
    }

    expressions(expressions: cog.Builder<azuremonitor.BuilderQueryEditorOrderByExpression>[]): this {
        const expressionsResources = expressions.map(builder1 => builder1.build());
        this.internal.expressions = expressionsResources;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

