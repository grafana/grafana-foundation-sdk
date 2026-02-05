"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TempoQueryBuilder = void 0;
const tslib_1 = require("tslib");
const tempo = tslib_1.__importStar(require("../tempo"));
class TempoQueryBuilder {
    constructor() {
        this.internal = tempo.defaultTempoQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId) {
        this.internal.refId = refId;
        return this;
    }
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    // TraceQL query or trace ID
    query(query) {
        this.internal.query = query;
        return this;
    }
    // @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
    search(search) {
        this.internal.search = search;
        return this;
    }
    // @deprecated Query traces by service name
    serviceName(serviceName) {
        this.internal.serviceName = serviceName;
        return this;
    }
    // @deprecated Query traces by span name
    spanName(spanName) {
        this.internal.spanName = spanName;
        return this;
    }
    // @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
    minDuration(minDuration) {
        this.internal.minDuration = minDuration;
        return this;
    }
    // @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
    maxDuration(maxDuration) {
        this.internal.maxDuration = maxDuration;
        return this;
    }
    // Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
    serviceMapQuery(serviceMapQuery) {
        this.internal.serviceMapQuery = serviceMapQuery;
        return this;
    }
    // Use service.namespace in addition to service.name to uniquely identify a service.
    serviceMapIncludeNamespace(serviceMapIncludeNamespace) {
        this.internal.serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        return this;
    }
    // Defines the maximum number of traces that are returned from Tempo
    limit(limit) {
        this.internal.limit = limit;
        return this;
    }
    // Defines the maximum number of spans per spanset that are returned from Tempo
    spss(spss) {
        this.internal.spss = spss;
        return this;
    }
    filters(filters) {
        const filtersResources = filters.map(builder1 => builder1.build());
        this.internal.filters = filtersResources;
        return this;
    }
    // Filters that are used to query the metrics summary
    groupBy(groupBy) {
        const groupByResources = groupBy.map(builder1 => builder1.build());
        this.internal.groupBy = groupByResources;
        return this;
    }
    // The type of the table that is used to display the search results
    tableType(tableType) {
        this.internal.tableType = tableType;
        return this;
    }
    // For metric queries, the step size to use
    step(step) {
        this.internal.step = step;
        return this;
    }
    // For metric queries, how many exemplars to request, 0 means no exemplars
    exemplars(exemplars) {
        this.internal.exemplars = exemplars;
        return this;
    }
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    // For metric queries, whether to run instant or range queries
    metricsQueryType(metricsQueryType) {
        this.internal.metricsQueryType = metricsQueryType;
        return this;
    }
}
exports.TempoQueryBuilder = TempoQueryBuilder;
//# sourceMappingURL=tempoQueryBuilder.gen.js.map