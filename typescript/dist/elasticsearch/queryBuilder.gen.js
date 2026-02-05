"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class QueryBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "elasticsearch";
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
    // Alias pattern
    alias(alias) {
        if (!this.internal.spec) {
            this.internal.spec = elasticsearch.defaultDataquery();
        }
        this.internal.spec.alias = alias;
        return this;
    }
    // Lucene query
    query(query) {
        if (!this.internal.spec) {
            this.internal.spec = elasticsearch.defaultDataquery();
        }
        this.internal.spec.query = query;
        return this;
    }
    // Name of time field
    timeField(timeField) {
        if (!this.internal.spec) {
            this.internal.spec = elasticsearch.defaultDataquery();
        }
        this.internal.spec.timeField = timeField;
        return this;
    }
    // List of bucket aggregations
    bucketAggs(bucketAggs) {
        if (!this.internal.spec) {
            this.internal.spec = elasticsearch.defaultDataquery();
        }
        const bucketAggsResources = bucketAggs.map(builder1 => builder1.build());
        this.internal.spec.bucketAggs = bucketAggsResources;
        return this;
    }
    // List of metric aggregations
    metrics(metrics) {
        if (!this.internal.spec) {
            this.internal.spec = elasticsearch.defaultDataquery();
        }
        const metricsResources = metrics.map(builder1 => builder1.build());
        this.internal.spec.metrics = metricsResources;
        return this;
    }
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId) {
        if (!this.internal.spec) {
            this.internal.spec = elasticsearch.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = elasticsearch.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType) {
        if (!this.internal.spec) {
            this.internal.spec = elasticsearch.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }
}
exports.QueryBuilder = QueryBuilder;
//# sourceMappingURL=queryBuilder.gen.js.map