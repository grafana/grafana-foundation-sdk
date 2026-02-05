// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as tempo from '../tempo';

export class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "tempo";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref: {
	name?: string;
}): this {
        this.internal.datasource = ref;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }

    // TraceQL query or trace ID
    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.query = query;
        return this;
    }

    // @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
    search(search: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.search = search;
        return this;
    }

    // @deprecated Query traces by service name
    serviceName(serviceName: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.serviceName = serviceName;
        return this;
    }

    // @deprecated Query traces by span name
    spanName(spanName: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.spanName = spanName;
        return this;
    }

    // @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
    minDuration(minDuration: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.minDuration = minDuration;
        return this;
    }

    // @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
    maxDuration(maxDuration: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.maxDuration = maxDuration;
        return this;
    }

    // Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
    serviceMapQuery(serviceMapQuery: string | string[]): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.serviceMapQuery = serviceMapQuery;
        return this;
    }

    // Use service.namespace in addition to service.name to uniquely identify a service.
    serviceMapIncludeNamespace(serviceMapIncludeNamespace: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        return this;
    }

    // Defines the maximum number of traces that are returned from Tempo
    limit(limit: number): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.limit = limit;
        return this;
    }

    // Defines the maximum number of spans per spanset that are returned from Tempo
    spss(spss: number): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.spss = spss;
        return this;
    }

    filters(filters: cog.Builder<tempo.TraceqlFilter>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        const filtersResources = filters.map(builder1 => builder1.build());
        this.internal.spec.filters = filtersResources;
        return this;
    }

    // Filters that are used to query the metrics summary
    groupBy(groupBy: cog.Builder<tempo.TraceqlFilter>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        const groupByResources = groupBy.map(builder1 => builder1.build());
        this.internal.spec.groupBy = groupByResources;
        return this;
    }

    // The type of the table that is used to display the search results
    tableType(tableType: tempo.SearchTableType): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.tableType = tableType;
        return this;
    }

    // For metric queries, the step size to use
    step(step: string): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.step = step;
        return this;
    }

    // For metric queries, how many exemplars to request, 0 means no exemplars
    exemplars(exemplars: number): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.exemplars = exemplars;
        return this;
    }

    // For metric queries, whether to run instant or range queries
    metricsQueryType(metricsQueryType: tempo.MetricsQueryType): this {
        if (!this.internal.spec) {
            this.internal.spec = tempo.defaultTempoQuery();
        }
        this.internal.spec.metricsQueryType = metricsQueryType;
        return this;
    }
}

