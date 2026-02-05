import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
import * as common from '../common';
export declare class AzureMonitorQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: azuremonitor.AzureMonitorQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureMonitorQuery;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    subscription(subscription: string): this;
    subscriptions(subscriptions: string[]): this;
    azureMonitor(azureMonitor: cog.Builder<azuremonitor.AzureMetricQuery>): this;
    azureLogAnalytics(azureLogAnalytics: cog.Builder<azuremonitor.AzureLogsQuery>): this;
    azureResourceGraph(azureResourceGraph: cog.Builder<azuremonitor.AzureResourceGraphQuery>): this;
    azureTraces(azureTraces: cog.Builder<azuremonitor.AzureTracesQuery>): this;
    grafanaTemplateVariableFn(grafanaTemplateVariableFn: cog.Builder<azuremonitor.GrafanaTemplateVariableQuery>): this;
    resourceGroup(resourceGroup: string): this;
    namespace(namespace: string): this;
    resource(resource: string): this;
    region(region: string): this;
    customNamespace(customNamespace: string): this;
    datasource(datasource: common.DataSourceRef): this;
    query(query: string): this;
}
