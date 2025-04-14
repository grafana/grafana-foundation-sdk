// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class SelectableValueBuilder implements cog.Builder<azuremonitor.SelectableValue> {
    protected readonly internal: azuremonitor.SelectableValue;

    constructor() {
        this.internal = azuremonitor.defaultSelectableValue();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.SelectableValue {
        return this.internal;
    }

    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

