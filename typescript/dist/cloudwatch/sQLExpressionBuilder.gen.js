"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.SQLExpressionBuilder = void 0;
const tslib_1 = require("tslib");
const cloudwatch = tslib_1.__importStar(require("../cloudwatch"));
class SQLExpressionBuilder {
    constructor() {
        this.internal = cloudwatch.defaultSQLExpression();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // SELECT part of the SQL expression
    select(select) {
        const selectResource = select.build();
        this.internal.select = selectResource;
        return this;
    }
    // FROM part of the SQL expression
    from(from) {
        const fromResource = from.build();
        this.internal.from = fromResource;
        return this;
    }
    // WHERE part of the SQL expression
    where(where) {
        const whereResource = where.build();
        this.internal.where = whereResource;
        return this;
    }
    // GROUP BY part of the SQL expression
    groupBy(groupBy) {
        const groupByResource = groupBy.build();
        this.internal.groupBy = groupByResource;
        return this;
    }
    // ORDER BY part of the SQL expression
    orderBy(orderBy) {
        const orderByResource = orderBy.build();
        this.internal.orderBy = orderByResource;
        return this;
    }
    // The sort order of the SQL expression, `ASC` or `DESC`
    orderByDirection(orderByDirection) {
        this.internal.orderByDirection = orderByDirection;
        return this;
    }
    // LIMIT part of the SQL expression
    limit(limit) {
        this.internal.limit = limit;
        return this;
    }
}
exports.SQLExpressionBuilder = SQLExpressionBuilder;
//# sourceMappingURL=sQLExpressionBuilder.gen.js.map