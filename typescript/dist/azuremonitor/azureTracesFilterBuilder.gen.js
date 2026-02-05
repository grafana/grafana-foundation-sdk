"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AzureTracesFilterBuilder = void 0;
const tslib_1 = require("tslib");
const azuremonitor = tslib_1.__importStar(require("../azuremonitor"));
class AzureTracesFilterBuilder {
    constructor() {
        this.internal = azuremonitor.defaultAzureTracesFilter();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Property name, auto-populated based on available traces.
    property(property) {
        this.internal.property = property;
        return this;
    }
    // Comparison operator to use. Either equals or not equals.
    operation(operation) {
        this.internal.operation = operation;
        return this;
    }
    // Values to filter by.
    filters(filters) {
        this.internal.filters = filters;
        return this;
    }
}
exports.AzureTracesFilterBuilder = AzureTracesFilterBuilder;
//# sourceMappingURL=azureTracesFilterBuilder.gen.js.map