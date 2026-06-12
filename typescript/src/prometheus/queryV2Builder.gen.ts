// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';
import * as prometheus from '../prometheus';

export class QueryV2Builder implements cog.Builder<dashboardv2.DataQueryKind> {
    protected readonly internal: dashboardv2.DataQueryKind;

    constructor() {
        this.internal = dashboardv2.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "prometheus";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    labels(labels: Record<string, string>): this {
        this.internal.labels = labels;
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

    // Additional Ad-hoc filters that take precedence over Scope on conflict.
    adhocFilters(adhocFilters: cog.Builder<prometheus.AdhocFilters>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        const adhocFiltersResources = adhocFilters.map(builder1 => builder1.build());
        this.internal.spec.adhocFilters = adhocFiltersResources;
        return this;
    }

    // what we should show in the editor
    // Possible enum values:
    //  - `"builder"` 
    //  - `"code"` 
    editorMode(editorMode: prometheus.QueryEditorMode): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.editorMode = editorMode;
        return this;
    }

    // Execute an additional query to identify interesting raw samples relevant for the given expr
    exemplar(exemplar: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.exemplar = exemplar;
        return this;
    }

    // The actual expression/query that will be evaluated by Prometheus
    expr(expr: string): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.expr = expr;
        return this;
    }

    // The response format
    // Possible enum values:
    //  - `"time_series"` 
    //  - `"table"` 
    //  - `"heatmap"` 
    format(format: prometheus.PromQueryFormat): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.format = format;
        return this;
    }

    // Group By parameters to apply to aggregate expressions in the query
    groupByKeys(groupByKeys: string[]): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.groupByKeys = groupByKeys;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    // Returns only the latest value that Prometheus has scraped for the requested time series
    instant(instant: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.instant = instant;
        return this;
    }

    // Used to specify how many times to divide max data points by. We use max data points under query options
    // See https://github.com/grafana/grafana/issues/48081
    // Deprecated: use interval
    intervalFactor(intervalFactor: number): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.intervalFactor = intervalFactor;
        return this;
    }

    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    intervalMs(intervalMs: number): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.intervalMs = intervalMs;
        return this;
    }

    // Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    legendFormat(legendFormat: string): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.legendFormat = legendFormat;
        return this;
    }

    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    maxDataPoints(maxDataPoints: number): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.maxDataPoints = maxDataPoints;
        return this;
    }

    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    queryType(queryType: string): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }

    // Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    range(range: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.range = range;
        return this;
    }

    // RefID is the unique identifier of the query, set by the frontend call.
    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    // Optionally define expected query result behavior
    resultAssertions(resultAssertions: cog.Builder<prometheus.ResultAssertions>): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        const resultAssertionsResource = resultAssertions.build();
        this.internal.spec.resultAssertions = resultAssertionsResource;
        return this;
    }

    // A set of filters applied to apply to the query
    scopes(scopes: cog.Builder<prometheus.Scopes>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        const scopesResources = scopes.map(builder1 => builder1.build());
        this.internal.spec.scopes = scopesResources;
        return this;
    }

    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    timeRange(timeRange: cog.Builder<prometheus.TimeRange>): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        const timeRangeResource = timeRange.build();
        this.internal.spec.timeRange = timeRangeResource;
        return this;
    }

    // An additional lower limit for the step parameter of the Prometheus query and for the
    // `$__interval` and `$__rate_interval` variables.
    interval(interval: string): this {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.interval = interval;
        return this;
    }
}

