"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
const loki = tslib_1.__importStar(require("../loki"));
class QueryBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "loki";
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
    // The LogQL query.
    expr(expr) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.expr = expr;
        return this;
    }
    // Used to override the name of the series.
    legendFormat(legendFormat) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.legendFormat = legendFormat;
        return this;
    }
    // Used to limit the number of log rows returned.
    maxLines(maxLines) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.maxLines = maxLines;
        return this;
    }
    // @deprecated, now use step.
    resolution(resolution) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.resolution = resolution;
        return this;
    }
    editorMode(editorMode) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.editorMode = editorMode;
        return this;
    }
    // @deprecated, now use queryType.
    range(range) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.range = range;
        return this;
    }
    // @deprecated, now use queryType.
    instant(instant) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.instant = instant;
        return this;
    }
    // Used to set step value for range queries.
    step(step) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.step = step;
        return this;
    }
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }
    direction(direction) {
        if (!this.internal.spec) {
            this.internal.spec = loki.defaultDataquery();
        }
        this.internal.spec.direction = direction;
        return this;
    }
}
exports.QueryBuilder = QueryBuilder;
//# sourceMappingURL=queryBuilder.gen.js.map