"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryEditorArrayExpressionBuilder = void 0;
const tslib_1 = require("tslib");
const cloudwatch = tslib_1.__importStar(require("../cloudwatch"));
class QueryEditorArrayExpressionBuilder {
    constructor() {
        this.internal = cloudwatch.defaultQueryEditorArrayExpression();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    type(type) {
        this.internal.type = type;
        return this;
    }
    expressions(expressions) {
        const expressionsResources = expressions.map(builder1 => builder1.build());
        this.internal.expressions = expressionsResources;
        return this;
    }
}
exports.QueryEditorArrayExpressionBuilder = QueryEditorArrayExpressionBuilder;
//# sourceMappingURL=queryEditorArrayExpressionBuilder.gen.js.map