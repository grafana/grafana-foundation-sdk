"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryEditorGroupByExpressionBuilder = void 0;
const tslib_1 = require("tslib");
const bigquery = tslib_1.__importStar(require("../bigquery"));
class QueryEditorGroupByExpressionBuilder {
    constructor() {
        this.internal = bigquery.defaultQueryEditorGroupByExpression();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    property(property) {
        const propertyResource = property.build();
        this.internal.property = propertyResource;
        return this;
    }
}
exports.QueryEditorGroupByExpressionBuilder = QueryEditorGroupByExpressionBuilder;
//# sourceMappingURL=queryEditorGroupByExpressionBuilder.gen.js.map