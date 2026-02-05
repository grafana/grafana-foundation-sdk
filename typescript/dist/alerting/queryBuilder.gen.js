"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class QueryBuilder {
    constructor(refId) {
        this.internal = alerting.defaultQuery();
        this.internal.refId = refId;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
    datasourceUid(datasourceUid) {
        this.internal.datasourceUid = datasourceUid;
        return this;
    }
    // JSON is the raw JSON query and includes the above properties as well as custom properties.
    model(model) {
        const modelResource = model.build();
        this.internal.model = modelResource;
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
    // RelativeTimeRange is the per query start and end time
    // for requests.
    relativeTimeRange(relativeTimeRange) {
        this.internal.relativeTimeRange = relativeTimeRange;
        return this;
    }
}
exports.QueryBuilder = QueryBuilder;
//# sourceMappingURL=queryBuilder.gen.js.map