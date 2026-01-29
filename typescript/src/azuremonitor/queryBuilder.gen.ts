// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as azuremonitor from '../azuremonitor';
import * as common from '../common';

export class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-azure-monitor-datasource";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref: {
	name?: string;
}): this {
        this.internal.datasource = ref;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }

    // Azure subscription containing the resource(s) to be queried.
    // Also used for template variable queries
    subscription(subscription: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.subscription = subscription;
        return this;
    }

    // Subscriptions to be queried via Azure Resource Graph.
    subscriptions(subscriptions: string[]): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.subscriptions = subscriptions;
        return this;
    }

    // Azure Monitor Metrics sub-query properties.
    azureMonitor(azureMonitor: cog.Builder<azuremonitor.AzureMetricQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        const azureMonitorResource = azureMonitor.build();
        this.internal.spec.azureMonitor = azureMonitorResource;
        return this;
    }

    // Azure Monitor Logs sub-query properties.
    azureLogAnalytics(azureLogAnalytics: cog.Builder<azuremonitor.AzureLogsQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        const azureLogAnalyticsResource = azureLogAnalytics.build();
        this.internal.spec.azureLogAnalytics = azureLogAnalyticsResource;
        return this;
    }

    // Azure Resource Graph sub-query properties.
    azureResourceGraph(azureResourceGraph: cog.Builder<azuremonitor.AzureResourceGraphQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        const azureResourceGraphResource = azureResourceGraph.build();
        this.internal.spec.azureResourceGraph = azureResourceGraphResource;
        return this;
    }

    // Application Insights Traces sub-query properties.
    azureTraces(azureTraces: cog.Builder<azuremonitor.AzureTracesQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        const azureTracesResource = azureTraces.build();
        this.internal.spec.azureTraces = azureTracesResource;
        return this;
    }

    // @deprecated Legacy template variable support.
    grafanaTemplateVariableFn(grafanaTemplateVariableFn: cog.Builder<azuremonitor.GrafanaTemplateVariableQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        const grafanaTemplateVariableFnResource = grafanaTemplateVariableFn.build();
        this.internal.spec.grafanaTemplateVariableFn = grafanaTemplateVariableFnResource;
        return this;
    }

    // Resource group used in template variable queries
    resourceGroup(resourceGroup: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.resourceGroup = resourceGroup;
        return this;
    }

    // Namespace used in template variable queries
    namespace(namespace: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.namespace = namespace;
        return this;
    }

    // Resource used in template variable queries
    resource(resource: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.resource = resource;
        return this;
    }

    // Region used in template variable queries
    region(region: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.region = region;
        return this;
    }

    // Custom namespace used in template variable queries
    customNamespace(customNamespace: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.customNamespace = customNamespace;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    oldDatasource(datasource: common.DataSourceRef): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.datasource = datasource;
        return this;
    }

    // Used only for exemplar queries from Prometheus
    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = azuremonitor.defaultAzureMonitorQuery();
        }
        this.internal.spec.query = query;
        return this;
    }
}

