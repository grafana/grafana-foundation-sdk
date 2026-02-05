"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AzureTracesQueryBuilder = void 0;
const tslib_1 = require("tslib");
const azuremonitor = tslib_1.__importStar(require("../azuremonitor"));
// Application Insights Traces sub-query properties
class AzureTracesQueryBuilder {
    constructor() {
        this.internal = azuremonitor.defaultAzureTracesQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Specifies the format results should be returned as.
    resultFormat(resultFormat) {
        this.internal.resultFormat = resultFormat;
        return this;
    }
    // Array of resource URIs to be queried.
    resources(resources) {
        this.internal.resources = resources;
        return this;
    }
    // Operation ID. Used only for Traces queries.
    operationId(operationId) {
        this.internal.operationId = operationId;
        return this;
    }
    // Types of events to filter by.
    traceTypes(traceTypes) {
        this.internal.traceTypes = traceTypes;
        return this;
    }
    // Filters for property values.
    filters(filters) {
        const filtersResources = filters.map(builder1 => builder1.build());
        this.internal.filters = filtersResources;
        return this;
    }
    // KQL query to be executed.
    query(query) {
        this.internal.query = query;
        return this;
    }
}
exports.AzureTracesQueryBuilder = AzureTracesQueryBuilder;
//# sourceMappingURL=azureTracesQueryBuilder.gen.js.map