"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryEditorFunctionExpressionBuilder = void 0;
const tslib_1 = require("tslib");
const cloudwatch = tslib_1.__importStar(require("../cloudwatch"));
class QueryEditorFunctionExpressionBuilder {
    constructor() {
        this.internal = cloudwatch.defaultQueryEditorFunctionExpression();
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
    parameters(parameters) {
        const parametersResources = parameters.map(builder1 => builder1.build());
        this.internal.parameters = parametersResources;
        return this;
    }
}
exports.QueryEditorFunctionExpressionBuilder = QueryEditorFunctionExpressionBuilder;
//# sourceMappingURL=queryEditorFunctionExpressionBuilder.gen.js.map