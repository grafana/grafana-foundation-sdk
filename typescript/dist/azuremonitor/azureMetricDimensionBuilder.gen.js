"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AzureMetricDimensionBuilder = void 0;
const tslib_1 = require("tslib");
const azuremonitor = tslib_1.__importStar(require("../azuremonitor"));
class AzureMetricDimensionBuilder {
    constructor() {
        this.internal = azuremonitor.defaultAzureMetricDimension();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Name of Dimension to be filtered on.
    dimension(dimension) {
        this.internal.dimension = dimension;
        return this;
    }
    // String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.
    operator(operator) {
        this.internal.operator = operator;
        return this;
    }
    // Values to match with the filter.
    filters(filters) {
        this.internal.filters = filters;
        return this;
    }
    // @deprecated filter is deprecated in favour of filters to support multiselect.
    filter(filter) {
        this.internal.filter = filter;
        return this;
    }
}
exports.AzureMetricDimensionBuilder = AzureMetricDimensionBuilder;
//# sourceMappingURL=azureMetricDimensionBuilder.gen.js.map