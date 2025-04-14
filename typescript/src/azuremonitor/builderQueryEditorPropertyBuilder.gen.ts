// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BuilderQueryEditorPropertyBuilder implements cog.Builder<azuremonitor.BuilderQueryEditorProperty> {
    protected readonly internal: azuremonitor.BuilderQueryEditorProperty;

    constructor() {
        this.internal = azuremonitor.defaultBuilderQueryEditorProperty();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.BuilderQueryEditorProperty {
        return this.internal;
    }

    type(type: azuremonitor.BuilderQueryEditorPropertyType): this {
        this.internal.type = type;
        return this;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }
}

