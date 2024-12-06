// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class AzureMetricDimensionBuilder implements cog.Builder<azuremonitor.AzureMetricDimension> {
    protected readonly internal: azuremonitor.AzureMetricDimension;

    constructor() {
        this.internal = azuremonitor.defaultAzureMetricDimension();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureMetricDimension {
        return this.internal;
    }

    // Name of Dimension to be filtered on.
    dimension(dimension: string): this {
        this.internal.dimension = dimension;
        return this;
    }

    // String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
    operator(operator: string): this {
        this.internal.operator = operator;
        return this;
    }

    // Values to match with the filter.
    filters(filters: string[]): this {
        this.internal.filters = filters;
        return this;
    }

    // @deprecated filter is deprecated in favour of filters to support multiselect.
    filter(filter: string): this {
        this.internal.filter = filter;
        return this;
    }
}
