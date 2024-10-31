// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
import * as dashboard from '../dashboard';

export class AzureMonitorQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: azuremonitor.AzureMonitorQuery;

    constructor() {
        this.internal = azuremonitor.defaultAzureMonitorQuery();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureMonitorQuery {
        return this.internal;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // Azure subscription containing the resource(s) to be queried.
    subscription(subscription: string): this {
        this.internal.subscription = subscription;
        return this;
    }

    // Subscriptions to be queried via Azure Resource Graph.
    subscriptions(subscriptions: string[]): this {
        this.internal.subscriptions = subscriptions;
        return this;
    }

    // Azure Monitor Metrics sub-query properties.
    azureMonitor(azureMonitor: cog.Builder<azuremonitor.AzureMetricQuery>): this {
        const azureMonitorResource = azureMonitor.build();
        this.internal.azureMonitor = azureMonitorResource;
        return this;
    }

    // Azure Monitor Logs sub-query properties.
    azureLogAnalytics(azureLogAnalytics: cog.Builder<azuremonitor.AzureLogsQuery>): this {
        const azureLogAnalyticsResource = azureLogAnalytics.build();
        this.internal.azureLogAnalytics = azureLogAnalyticsResource;
        return this;
    }

    // Azure Resource Graph sub-query properties.
    azureResourceGraph(azureResourceGraph: cog.Builder<azuremonitor.AzureResourceGraphQuery>): this {
        const azureResourceGraphResource = azureResourceGraph.build();
        this.internal.azureResourceGraph = azureResourceGraphResource;
        return this;
    }

    // Application Insights Traces sub-query properties.
    azureTraces(azureTraces: cog.Builder<azuremonitor.AzureTracesQuery>): this {
        const azureTracesResource = azureTraces.build();
        this.internal.azureTraces = azureTracesResource;
        return this;
    }

    // @deprecated Legacy template variable support.
    grafanaTemplateVariableFn(grafanaTemplateVariableFn: azuremonitor.GrafanaTemplateVariableQuery): this {
        this.internal.grafanaTemplateVariableFn = grafanaTemplateVariableFn;
        return this;
    }

    // Template variables params. These exist for backwards compatiblity with legacy template variables.
    resourceGroup(resourceGroup: string): this {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }

    namespace(namespace: string): this {
        this.internal.namespace = namespace;
        return this;
    }

    resource(resource: string): this {
        this.internal.resource = resource;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // Azure Monitor query type.
    // queryType: #AzureQueryType
    region(region: string): this {
        this.internal.region = region;
        return this;
    }
}
