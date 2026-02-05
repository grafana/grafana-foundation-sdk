"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MetricFindValueBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Define the MetricFindValue type
class MetricFindValueBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultMetricFindValue();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    text(text) {
        this.internal.text = text;
        return this;
    }
    value(value) {
        this.internal.value = value;
        return this;
    }
    group(group) {
        this.internal.group = group;
        return this;
    }
    expandable(expandable) {
        this.internal.expandable = expandable;
        return this;
    }
}
exports.MetricFindValueBuilder = MetricFindValueBuilder;
//# sourceMappingURL=metricFindValueBuilder.gen.js.map