// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorWhereExpressionItemsBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorWhereExpressionItems> {
    protected readonly internal: azuremonitor.BuilderQueryEditorWhereExpressionItems;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorWhereExpressionItems();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorWhereExpressionItems {
        return this.internal;
    }

    property(property: cog.Builder<azuremonitor.BuilderQueryEditorProperty>): this {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }

    operator(operator: cog.Builder<azuremonitor.BuilderQueryEditorOperator>): this {
        const operatorResource = operator.build();
        this.internal.operator = operatorResource;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

