// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorColumnsExpressionBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorColumnsExpression> {
    protected readonly internal: azuremonitor.BuilderQueryEditorColumnsExpression;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorColumnsExpression();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorColumnsExpression {
        return this.internal;
    }

    columns(columns: string[]): this {
        this.internal.columns = columns;
        return this;
    }

    type(type: azuremonitor.BuilderQueryEditorExpressionType): this {
        this.internal.type = type;
        return this;
    }
}

