// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class AzureTracesFilterBuilder implements cog.Builder<azuremonitor.AzureTracesFilter> {
    protected readonly internal: azuremonitor.AzureTracesFilter;

    constructor() {
        this.internal = azuremonitor.defaultAzureTracesFilter();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureTracesFilter {
        return this.internal;
    }

    // Property name, auto-populated based on available traces.
    property(property: string): this {
        this.internal.property = property;
        return this;
    }

    // Comparison operator to use. Either equals or not equals.
    operation(operation: string): this {
        this.internal.operation = operation;
        return this;
    }

    // Values to filter by.
    filters(filters: string[]): this {
        this.internal.filters = filters;
        return this;
    }
}
