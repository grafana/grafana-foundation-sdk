"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryEditorFunctionParameterExpressionBuilder = void 0;
const tslib_1 = require("tslib");
const bigquery = tslib_1.__importStar(require("../bigquery"));
class QueryEditorFunctionParameterExpressionBuilder {
    constructor() {
        this.internal = bigquery.defaultQueryEditorFunctionParameterExpression();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
}
exports.QueryEditorFunctionParameterExpressionBuilder = QueryEditorFunctionParameterExpressionBuilder;
//# sourceMappingURL=queryEditorFunctionParameterExpressionBuilder.gen.js.map