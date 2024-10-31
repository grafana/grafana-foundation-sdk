// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as tempo from '../tempo';
import * as dashboard from '../dashboard';

export class TempoQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: tempo.TempoQuery;

    constructor() {
        this.internal = tempo.defaultTempoQuery();
    }

    /**
     * Builds the object.
     */
    build(): tempo.TempoQuery {
        return this.internal;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // TraceQL query or trace ID
    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    // @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
    search(search: string): this {
        this.internal.search = search;
        return this;
    }

    // @deprecated Query traces by service name
    serviceName(serviceName: string): this {
        this.internal.serviceName = serviceName;
        return this;
    }

    // @deprecated Query traces by span name
    spanName(spanName: string): this {
        this.internal.spanName = spanName;
        return this;
    }

    // @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
    minDuration(minDuration: string): this {
        this.internal.minDuration = minDuration;
        return this;
    }

    // @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
    maxDuration(maxDuration: string): this {
        this.internal.maxDuration = maxDuration;
        return this;
    }

    // Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}
    serviceMapQuery(serviceMapQuery: string): this {
        this.internal.serviceMapQuery = serviceMapQuery;
        return this;
    }

    // Use service.namespace in addition to service.name to uniquely identify a service.
    serviceMapIncludeNamespace(serviceMapIncludeNamespace: boolean): this {
        this.internal.serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        return this;
    }

    // Defines the maximum number of traces that are returned from Tempo
    limit(limit: number): this {
        this.internal.limit = limit;
        return this;
    }

    // Defines the maximum number of spans per spanset that are returned from Tempo
    spss(spss: number): this {
        this.internal.spss = spss;
        return this;
    }

    filters(filters: cog.Builder<tempo.TraceqlFilter>[]): this {
        const filtersResources = filters.map(builder1 => builder1.build());
        this.internal.filters = filtersResources;
        return this;
    }

    // Filters that are used to query the metrics summary
    groupBy(groupBy: cog.Builder<tempo.TraceqlFilter>[]): this {
        const groupByResources = groupBy.map(builder1 => builder1.build());
        this.internal.groupBy = groupByResources;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // The type of the table that is used to display the search results
    tableType(tableType: tempo.SearchTableType): this {
        this.internal.tableType = tableType;
        return this;
    }
}
