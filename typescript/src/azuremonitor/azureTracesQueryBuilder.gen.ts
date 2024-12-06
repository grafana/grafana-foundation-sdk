// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

// Application Insights Traces sub-query properties
export class AzureTracesQueryBuilder implements cog.Builder<azuremonitor.AzureTracesQuery> {
    protected readonly internal: azuremonitor.AzureTracesQuery;

    constructor() {
        this.internal = azuremonitor.defaultAzureTracesQuery();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureTracesQuery {
        return this.internal;
    }

    // Specifies the format results should be returned as.
    resultFormat(resultFormat: azuremonitor.ResultFormat): this {
        this.internal.resultFormat = resultFormat;
        return this;
    }

    // Array of resource URIs to be queried.
    resources(resources: string[]): this {
        this.internal.resources = resources;
        return this;
    }

    // Operation ID. Used only for Traces queries.
    operationId(operationId: string): this {
        this.internal.operationId = operationId;
        return this;
    }

    // Types of events to filter by.
    traceTypes(traceTypes: string[]): this {
        this.internal.traceTypes = traceTypes;
        return this;
    }

    // Filters for property values.
    filters(filters: cog.Builder<azuremonitor.AzureTracesFilter>[]): this {
        const filtersResources = filters.map(builder1 => builder1.build());
        this.internal.filters = filtersResources;
        return this;
    }

    // KQL query to be executed.
    query(query: string): this {
        this.internal.query = query;
        return this;
    }
}
