// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorOperatorBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorOperator> {
    protected readonly internal: azuremonitor.BuilderQueryEditorOperator;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorOperator();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorOperator {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    labelValue(labelValue: string): this {
        this.internal.labelValue = labelValue;
        return this;
    }
}

