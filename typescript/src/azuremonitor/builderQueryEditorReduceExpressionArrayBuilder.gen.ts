// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorReduceExpressionArrayBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorReduceExpressionArray> {
    protected readonly internal: azuremonitor.BuilderQueryEditorReduceExpressionArray;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorReduceExpressionArray();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorReduceExpressionArray {
        return this.internal;
    }

    expressions(expressions: cog.Builder<azuremonitor.BuilderQueryEditorReduceExpression>[]): this {
        const expressionsResources = expressions.map(builder1 => builder1.build());
        this.internal.expressions = expressionsResources;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

