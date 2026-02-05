import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as azuremonitor from '../azuremonitor';
export declare class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind;
    version(version: string): this;
    datasource(ref: {
        name?: string;
    }): this;
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
    query(query: string): this;
}
