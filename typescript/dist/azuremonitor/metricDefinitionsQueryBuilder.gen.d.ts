import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class MetricDefinitionsQueryBuilder implements cog.Builder<azuremonitor.MetricDefinitionsQuery> {
    protected readonly internal: azuremonitor.MetricDefinitionsQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.MetricDefinitionsQuery;
    rawQuery(rawQuery: string): this;
    subscription(subscription: string): this;
    resourceGroup(resourceGroup: string): this;
    metricNamespace(metricNamespace: string): this;
    resourceName(resourceName: string): this;
}
