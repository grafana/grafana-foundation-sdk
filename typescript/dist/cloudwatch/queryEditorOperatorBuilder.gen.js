"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryEditorOperatorBuilder = void 0;
const tslib_1 = require("tslib");
const cloudwatch = tslib_1.__importStar(require("../cloudwatch"));
// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
class QueryEditorOperatorBuilder {
    constructor() {
        this.internal = cloudwatch.defaultQueryEditorOperator();
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
    value(value) {
        this.internal.value = value;
        return this;
    }
}
exports.QueryEditorOperatorBuilder = QueryEditorOperatorBuilder;
//# sourceMappingURL=queryEditorOperatorBuilder.gen.js.map