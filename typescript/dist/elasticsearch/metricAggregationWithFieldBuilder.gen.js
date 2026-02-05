"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MetricAggregationWithFieldBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class MetricAggregationWithFieldBuilder {
    constructor() {
        this.internal = elasticsearch.defaultMetricAggregationWithField();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    field(field) {
        this.internal.field = field;
        return this;
    }
    type(type) {
        this.internal.type = type;
        return this;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
}
exports.MetricAggregationWithFieldBuilder = MetricAggregationWithFieldBuilder;
//# sourceMappingURL=metricAggregationWithFieldBuilder.gen.js.map