// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as prometheus from '../prometheus';
import * as dashboard from '../dashboard';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: prometheus.dataquery;

    constructor() {
        this.internal = prometheus.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): prometheus.dataquery {
        return this.internal;
    }

    // The actual expression/query that will be evaluated by Prometheus
    expr(expr: string): this {
        this.internal.expr = expr;
        return this;
    }

    // Returns only the latest value that Prometheus has scraped for the requested time series
    instant(): this {
        this.internal.instant = true;
        this.internal.range = false;
        return this;
    }

    // Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    range(): this {
        this.internal.range = true;
        this.internal.instant = false;
        return this;
    }

    // Execute an additional query to identify interesting raw samples relevant for the given expr
    exemplar(exemplar: boolean): this {
        this.internal.exemplar = exemplar;
        return this;
    }

    // Specifies which editor is being used to prepare the query. It can be "code" or "builder"
    editorMode(editorMode: prometheus.QueryEditorMode): this {
        this.internal.editorMode = editorMode;
        return this;
    }

    // Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
    format(format: prometheus.PromQueryFormat): this {
        this.internal.format = format;
        return this;
    }

    // Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    legendFormat(legendFormat: string): this {
        this.internal.legendFormat = legendFormat;
        return this;
    }

    // @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
    // See https://github.com/grafana/grafana/issues/48081
    intervalFactor(intervalFactor: number): this {
        this.internal.intervalFactor = intervalFactor;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
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
