"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.FilterBuilder = void 0;
const tslib_1 = require("tslib");
const googlecloudmonitoring = tslib_1.__importStar(require("../googlecloudmonitoring"));
// Query filter representation.
class FilterBuilder {
    constructor() {
        this.internal = googlecloudmonitoring.defaultFilter();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Filter key.
    key(key) {
        this.internal.key = key;
        return this;
    }
    // Filter operator.
    operator(operator) {
        this.internal.operator = operator;
        return this;
    }
    // Filter value.
    value(value) {
        this.internal.value = value;
        return this;
    }
    // Filter condition.
    condition(condition) {
        this.internal.condition = condition;
        return this;
    }
}
exports.FilterBuilder = FilterBuilder;
//# sourceMappingURL=filterBuilder.gen.js.map