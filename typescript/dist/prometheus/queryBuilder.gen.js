"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
const prometheus = tslib_1.__importStar(require("../prometheus"));
class QueryBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "prometheus";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    version(version) {
        this.internal.version = version;
        return this;
    }
    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref) {
        this.internal.datasource = ref;
        return this;
    }
    // The actual expression/query that will be evaluated by Prometheus
    expr(expr) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.expr = expr;
        return this;
    }
    // Returns only the latest value that Prometheus has scraped for the requested time series
    instant() {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.instant = true;
        return this;
    }
    // Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    range() {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.range = true;
        return this;
    }
    // Execute an additional query to identify interesting raw samples relevant for the given expr
    exemplar(exemplar) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.exemplar = exemplar;
        return this;
    }
    // Specifies which editor is being used to prepare the query. It can be "code" or "builder"
    editorMode(editorMode) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.editorMode = editorMode;
        return this;
    }
    // Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
    format(format) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.format = format;
        return this;
    }
    // Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    legendFormat(legendFormat) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.legendFormat = legendFormat;
        return this;
    }
    // @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
    // See https://github.com/grafana/grafana/issues/48081
    intervalFactor(intervalFactor) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.intervalFactor = intervalFactor;
        return this;
    }
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }
    // An additional lower limit for the step parameter of the Prometheus query and for the
    // `$__interval` and `$__rate_interval` variables.
    interval(interval) {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.interval = interval;
        return this;
    }
    rangeAndInstant() {
        if (!this.internal.spec) {
            this.internal.spec = prometheus.defaultDataquery();
        }
        this.internal.spec.range = true;
        this.internal.spec.instant = true;
        return this;
    }
}
exports.QueryBuilder = QueryBuilder;
//# sourceMappingURL=queryBuilder.gen.js.map