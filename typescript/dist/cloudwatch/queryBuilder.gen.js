"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class QueryBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "cloudwatch";
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
    cloudWatchMetricsQuery(cloudWatchMetricsQuery) {
        const cloudWatchMetricsQueryResource = cloudWatchMetricsQuery.build();
        this.internal.spec = cloudWatchMetricsQueryResource;
        return this;
    }
    cloudWatchLogsQuery(cloudWatchLogsQuery) {
        const cloudWatchLogsQueryResource = cloudWatchLogsQuery.build();
        this.internal.spec = cloudWatchLogsQueryResource;
        return this;
    }
    cloudWatchAnnotationQuery(cloudWatchAnnotationQuery) {
        const cloudWatchAnnotationQueryResource = cloudWatchAnnotationQuery.build();
        this.internal.spec = cloudWatchAnnotationQueryResource;
        return this;
    }
}
exports.QueryBuilder = QueryBuilder;
//# sourceMappingURL=queryBuilder.gen.js.map