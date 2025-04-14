// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorGroupByExpressionArrayBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorGroupByExpressionArray> {
    protected readonly internal: azuremonitor.BuilderQueryEditorGroupByExpressionArray;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorGroupByExpressionArray();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorGroupByExpressionArray {
        return this.internal;
    }

    expressions(expressions: cog.Builder<azuremonitor.BuilderQueryEditorGroupByExpression>[]): this {
        const expressionsResources = expressions.map(builder1 => builder1.build());
        this.internal.expressions = expressionsResources;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

