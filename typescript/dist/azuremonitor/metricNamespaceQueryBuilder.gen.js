"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MetricNamespaceQueryBuilder = void 0;
const tslib_1 = require("tslib");
const azuremonitor = tslib_1.__importStar(require("../azuremonitor"));
class MetricNamespaceQueryBuilder {
    constructor() {
        this.internal = azuremonitor.defaultMetricNamespaceQuery();
        this.internal.kind = "MetricNamespaceQuery";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    rawQuery(rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    subscription(subscription) {
        this.internal.subscription = subscription;
        return this;
    }
    resourceGroup(resourceGroup) {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }
    metricNamespace(metricNamespace) {
        this.internal.metricNamespace = metricNamespace;
        return this;
    }
    resourceName(resourceName) {
        this.internal.resourceName = resourceName;
        return this;
    }
}
exports.MetricNamespaceQueryBuilder = MetricNamespaceQueryBuilder;
//# sourceMappingURL=metricNamespaceQueryBuilder.gen.js.map