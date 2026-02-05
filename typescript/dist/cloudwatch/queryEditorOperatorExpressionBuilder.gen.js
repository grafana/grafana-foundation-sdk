"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryEditorOperatorExpressionBuilder = void 0;
const tslib_1 = require("tslib");
const cloudwatch = tslib_1.__importStar(require("../cloudwatch"));
class QueryEditorOperatorExpressionBuilder {
    constructor() {
        this.internal = cloudwatch.defaultQueryEditorOperatorExpression();
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
    // TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
    operator(operator) {
        const operatorResource = operator.build();
        this.internal.operator = operatorResource;
        return this;
    }
}
exports.QueryEditorOperatorExpressionBuilder = QueryEditorOperatorExpressionBuilder;
//# sourceMappingURL=queryEditorOperatorExpressionBuilder.gen.js.map