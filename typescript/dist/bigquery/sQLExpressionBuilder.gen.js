"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.SQLExpressionBuilder = void 0;
const tslib_1 = require("tslib");
const bigquery = tslib_1.__importStar(require("../bigquery"));
class SQLExpressionBuilder {
    constructor() {
        this.internal = bigquery.defaultSQLExpression();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    columns(columns) {
        const columnsResources = columns.map(builder1 => builder1.build());
        this.internal.columns = columnsResources;
        return this;
    }
    from(from) {
        this.internal.from = from;
        return this;
    }
    // whereJsonTree?: _
    whereString(whereString) {
        this.internal.whereString = whereString;
        return this;
    }
    groupBy(groupBy) {
        const groupByResources = groupBy.map(builder1 => builder1.build());
        this.internal.groupBy = groupByResources;
        return this;
    }
    orderBy(orderBy) {
        const orderByResource = orderBy.build();
        this.internal.orderBy = orderByResource;
        return this;
    }
    orderByDirection(orderByDirection) {
        this.internal.orderByDirection = orderByDirection;
        return this;
    }
    limit(limit) {
        this.internal.limit = limit;
        return this;
    }
    offset(offset) {
        this.internal.offset = offset;
        return this;
    }
}
exports.SQLExpressionBuilder = SQLExpressionBuilder;
//# sourceMappingURL=sQLExpressionBuilder.gen.js.map