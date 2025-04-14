// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorOrderByExpressionBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorOrderByExpression> {
    protected readonly internal: azuremonitor.BuilderQueryEditorOrderByExpression;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorOrderByExpression();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorOrderByExpression {
        return this.internal;
    }

    property(property: cog.Builder<azuremonitor.BuilderQueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }

    order(order: azuremonitor.BuilderQueryEditorOrderByOptions): this {
        this.internal.order = order;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

