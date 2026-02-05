import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class MetricNamespaceQueryBuilder implements cog.Builder<azuremonitor.MetricNamespaceQuery> {
    protected readonly internal: azuremonitor.MetricNamespaceQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.MetricNamespaceQuery;
    rawQuery(rawQuery: string): this;
    subscription(subscription: string): this;
    resourceGroup(resourceGroup: string): this;
    metricNamespace(metricNamespace: string): this;
    resourceName(resourceName: string): this;
}
