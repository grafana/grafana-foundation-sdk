// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorPropertyExpressionBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorPropertyExpression> {
    protected readonly internal: azuremonitor.BuilderQueryEditorPropertyExpression;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorPropertyExpression();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorPropertyExpression {
        return this.internal;
    }

    property(property: cog.Builder<azuremonitor.BuilderQueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

