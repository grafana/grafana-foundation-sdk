// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorReduceExpressionBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorReduceExpression> {
    protected readonly internal: azuremonitor.BuilderQueryEditorReduceExpression;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorReduceExpression();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorReduceExpression {
        return this.internal;
    }

    property(property: cog.Builder<azuremonitor.BuilderQueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }

    reduce(reduce: cog.Builder<azuremonitor.BuilderQueryEditorProperty>): this {
        const reduceResource = reduce.build();
        this.internal.reduce = reduceResource;
        return this;
    }

    parameters(parameters: cog.Builder<azuremonitor.BuilderQueryEditorFunctionParameterExpression>[]): this {
        const parametersResources = parameters.map(builder1 => builder1.build());
        this.internal.parameters = parametersResources;
        return this;
    }

    focus(focus: boolean): this {
        this.internal.focus = focus;
        return this;
    }
}

