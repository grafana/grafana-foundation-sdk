// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as prometheus from '../prometheus';
import * as common from '../common';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: prometheus.Dataquery;

    constructor() {
        this.internal = prometheus.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): prometheus.Dataquery {
        return this.internal;
    }

    // Additional Ad-hoc filters that take precedence over Scope on conflict.
    adhocFilters(adhocFilters: cog.Builder<prometheus.AdhocFilters>[]): this {
        const adhocFiltersResources = adhocFilters.map(builder1 => builder1.build());
        this.internal.adhocFilters = adhocFiltersResources;
        return this;
    }

    // The datasource
    datasource(datasource: common.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // what we should show in the editor
    // Possible enum values:
    //  - `"builder"` 
    //  - `"code"` 
    editorMode(editorMode: prometheus.QueryEditorMode): this {
        this.internal.editorMode = editorMode;
        return this;
    }

    // Execute an additional query to identify interesting raw samples relevant for the given expr
    exemplar(exemplar: boolean): this {
        this.internal.exemplar = exemplar;
        return this;
    }

    // The actual expression/query that will be evaluated by Prometheus
    expr(expr: string): this {
        this.internal.expr = expr;
        return this;
    }

    // The response format
    // Possible enum values:
    //  - `"time_series"` 
    //  - `"table"` 
    //  - `"heatmap"` 
    format(format: prometheus.PromQueryFormat): this {
        this.internal.format = format;
        return this;
    }

    // Group By parameters to apply to aggregate expressions in the query
    groupByKeys(groupByKeys: string[]): this {
        this.internal.groupByKeys = groupByKeys;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Returns only the latest value that Prometheus has scraped for the requested time series
    instant(): this {
        this.internal.instant = true;
        this.internal.range = false;
        return this;
    }

    // Used to specify how many times to divide max data points by. We use max data points under query options
    // See https://github.com/grafana/grafana/issues/48081
    // Deprecated: use interval
    intervalFactor(intervalFactor: number): this {
        this.internal.intervalFactor = intervalFactor;
        return this;
    }

    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    intervalMs(intervalMs: number): this {
        this.internal.intervalMs = intervalMs;
        return this;
    }

    // Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    legendFormat(legendFormat: string): this {
        this.internal.legendFormat = legendFormat;
        return this;
    }

    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    maxDataPoints(maxDataPoints: number): this {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }

    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    range(): this {
        this.internal.range = true;
        this.internal.instant = false;
        return this;
    }

    // RefID is the unique identifier of the query, set by the frontend call.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // Optionally define expected query result behavior
    resultAssertions(resultAssertions: cog.Builder<prometheus.ResultAssertions>): this {
        const resultAssertionsResource = resultAssertions.build();
        this.internal.resultAssertions = resultAssertionsResource;
        return this;
    }

    // A set of filters applied to apply to the query
    scopes(scopes: cog.Builder<prometheus.Scopes>[]): this {
        const scopesResources = scopes.map(builder1 => builder1.build());
        this.internal.scopes = scopesResources;
        return this;
    }

    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    timeRange(timeRange: cog.Builder<prometheus.TimeRange>): this {
        const timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }

    // An additional lower limit for the step parameter of the Prometheus query and for the
    // `$__interval` and `$__rate_interval` variables.
    interval(interval: string): this {
        this.internal.interval = interval;
        return this;
    }

    rangeAndInstant(): this {
        this.internal.range = true;
        this.internal.instant = true;
        return this;
    }
}

