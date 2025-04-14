// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorFunctionParameterExpressionBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorFunctionParameterExpression> {
    protected readonly internal: azuremonitor.BuilderQueryEditorFunctionParameterExpression;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorFunctionParameterExpression();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorFunctionParameterExpression {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    fieldType(fieldType: azuremonitor.BuilderQueryEditorPropertyType): this {
        this.internal.fieldType = fieldType;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

