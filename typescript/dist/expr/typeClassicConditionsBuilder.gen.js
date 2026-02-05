"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TypeClassicConditionsBuilder = void 0;
const tslib_1 = require("tslib");
const expr = tslib_1.__importStar(require("../expr"));
class TypeClassicConditionsBuilder {
    constructor() {
        this.internal = expr.defaultTypeClassicConditions();
        this.internal.type = "classic_conditions";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    conditions(conditions) {
        this.internal.conditions = conditions;
        return this;
    }
    // The datasource
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    // true if query is disabled (ie should not be returned to the dashboard)
    // NOTE: this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    // Interval is the suggested duration between time points in a time series query.
    // NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    // from the interval required to fill a pixels in the visualization
    intervalMs(intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    // MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    // NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    // from the number of pixels visible in a visualization
    maxDataPoints(maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    // QueryType is an optional identifier for the type of query.
    // It can be used to distinguish different types of queries.
    queryType(queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    // RefID is the unique identifier of the query, set by the frontend call.
    refId(refId) {
        this.internal.refId = refId;
        return this;
    }
    // Optionally define expected query result behavior
    resultAssertions(resultAssertions) {
        this.internal.resultAssertions = resultAssertions;
        return this;
    }
    // TimeRange represents the query range
    // NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    // NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    timeRange(timeRange) {
        this.internal.timeRange = timeRange;
        return this;
    }
}
exports.TypeClassicConditionsBuilder = TypeClassicConditionsBuilder;
//# sourceMappingURL=typeClassicConditionsBuilder.gen.js.map