"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AzureMonitorQueryBuilder = void 0;
const tslib_1 = require("tslib");
const azuremonitor = tslib_1.__importStar(require("../azuremonitor"));
class AzureMonitorQueryBuilder {
    constructor() {
        this.internal = azuremonitor.defaultAzureMonitorQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId) {
        this.internal.refId = refId;
        return this;
    }
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    // Azure subscription containing the resource(s) to be queried.
    // Also used for template variable queries
    subscription(subscription) {
        this.internal.subscription = subscription;
        return this;
    }
    // Subscriptions to be queried via Azure Resource Graph.
    subscriptions(subscriptions) {
        this.internal.subscriptions = subscriptions;
        return this;
    }
    // Azure Monitor Metrics sub-query properties.
    azureMonitor(azureMonitor) {
        const azureMonitorResource = azureMonitor.build();
        this.internal.azureMonitor = azureMonitorResource;
        return this;
    }
    // Azure Monitor Logs sub-query properties.
    azureLogAnalytics(azureLogAnalytics) {
        const azureLogAnalyticsResource = azureLogAnalytics.build();
        this.internal.azureLogAnalytics = azureLogAnalyticsResource;
        return this;
    }
    // Azure Resource Graph sub-query properties.
    azureResourceGraph(azureResourceGraph) {
        const azureResourceGraphResource = azureResourceGraph.build();
        this.internal.azureResourceGraph = azureResourceGraphResource;
        return this;
    }
    // Application Insights Traces sub-query properties.
    azureTraces(azureTraces) {
        const azureTracesResource = azureTraces.build();
        this.internal.azureTraces = azureTracesResource;
        return this;
    }
    // @deprecated Legacy template variable support.
    grafanaTemplateVariableFn(grafanaTemplateVariableFn) {
        const grafanaTemplateVariableFnResource = grafanaTemplateVariableFn.build();
        this.internal.grafanaTemplateVariableFn = grafanaTemplateVariableFnResource;
        return this;
    }
    // Resource group used in template variable queries
    resourceGroup(resourceGroup) {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }
    // Namespace used in template variable queries
    namespace(namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    // Resource used in template variable queries
    resource(resource) {
        this.internal.resource = resource;
        return this;
    }
    // Region used in template variable queries
    region(region) {
        this.internal.region = region;
        return this;
    }
    // Custom namespace used in template variable queries
    customNamespace(customNamespace) {
        this.internal.customNamespace = customNamespace;
        return this;
    }
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    // Used only for exemplar queries from Prometheus
    query(query) {
        this.internal.query = query;
        return this;
    }
}
exports.AzureMonitorQueryBuilder = AzureMonitorQueryBuilder;
//# sourceMappingURL=azureMonitorQueryBuilder.gen.js.map