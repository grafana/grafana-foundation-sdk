"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AzureMetricQueryBuilder = void 0;
const tslib_1 = require("tslib");
const azuremonitor = tslib_1.__importStar(require("../azuremonitor"));
class AzureMetricQueryBuilder {
    constructor() {
        this.internal = azuremonitor.defaultAzureMetricQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Array of resource URIs to be queried.
    resources(resources) {
        const resourcesResources = resources.map(builder1 => builder1.build());
        this.internal.resources = resourcesResources;
        return this;
    }
    // metricNamespace is used as the resource type (or resource namespace).
    // It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
    // Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
    metricNamespace(metricNamespace) {
        this.internal.metricNamespace = metricNamespace;
        return this;
    }
    // Used as the value for the metricNamespace property when it's different from the resource namespace.
    customNamespace(customNamespace) {
        this.internal.customNamespace = customNamespace;
        return this;
    }
    // The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
    metricName(metricName) {
        this.internal.metricName = metricName;
        return this;
    }
    // The Azure region containing the resource(s).
    region(region) {
        this.internal.region = region;
        return this;
    }
    // The granularity of data points to be queried. Defaults to auto.
    timeGrain(timeGrain) {
        this.internal.timeGrain = timeGrain;
        return this;
    }
    // The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
    aggregation(aggregation) {
        this.internal.aggregation = aggregation;
        return this;
    }
    // Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
    dimensionFilters(dimensionFilters) {
        const dimensionFiltersResources = dimensionFilters.map(builder1 => builder1.build());
        this.internal.dimensionFilters = dimensionFiltersResources;
        return this;
    }
    // Maximum number of records to return. Defaults to 10.
    top(top) {
        this.internal.top = top;
        return this;
    }
    // Time grains that are supported by the metric.
    allowedTimeGrainsMs(allowedTimeGrainsMs) {
        this.internal.allowedTimeGrainsMs = allowedTimeGrainsMs;
        return this;
    }
    // Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
    alias(alias) {
        this.internal.alias = alias;
        return this;
    }
    // @deprecated
    timeGrainUnit(timeGrainUnit) {
        this.internal.timeGrainUnit = timeGrainUnit;
        return this;
    }
    // @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    dimension(dimension) {
        this.internal.dimension = dimension;
        return this;
    }
    // @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    dimensionFilter(dimensionFilter) {
        this.internal.dimensionFilter = dimensionFilter;
        return this;
    }
    // @deprecated Use metricNamespace instead
    metricDefinition(metricDefinition) {
        this.internal.metricDefinition = metricDefinition;
        return this;
    }
    // @deprecated Use resourceGroup, resourceName and metricNamespace instead
    resourceUri(resourceUri) {
        this.internal.resourceUri = resourceUri;
        return this;
    }
    // @deprecated Use resources instead
    resourceGroup(resourceGroup) {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }
    // @deprecated Use resources instead
    resourceName(resourceName) {
        this.internal.resourceName = resourceName;
        return this;
    }
}
exports.AzureMetricQueryBuilder = AzureMetricQueryBuilder;
//# sourceMappingURL=azureMetricQueryBuilder.gen.js.map