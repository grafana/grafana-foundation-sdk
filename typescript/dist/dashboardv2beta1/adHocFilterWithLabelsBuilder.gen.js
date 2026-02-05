"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdHocFilterWithLabelsBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Define the AdHocFilterWithLabels type
class AdHocFilterWithLabelsBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultAdHocFilterWithLabels();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    key(key) {
        this.internal.key = key;
        return this;
    }
    operator(operator) {
        this.internal.operator = operator;
        return this;
    }
    value(value) {
        this.internal.value = value;
        return this;
    }
    values(values) {
        this.internal.values = values;
        return this;
    }
    keyLabel(keyLabel) {
        this.internal.keyLabel = keyLabel;
        return this;
    }
    valueLabels(valueLabels) {
        this.internal.valueLabels = valueLabels;
        return this;
    }
    forceEdit(forceEdit) {
        this.internal.forceEdit = forceEdit;
        return this;
    }
    origin(origin) {
        this.internal.origin = origin;
        return this;
    }
    // @deprecated
    condition(condition) {
        this.internal.condition = condition;
        return this;
    }
}
exports.AdHocFilterWithLabelsBuilder = AdHocFilterWithLabelsBuilder;
//# sourceMappingURL=adHocFilterWithLabelsBuilder.gen.js.map