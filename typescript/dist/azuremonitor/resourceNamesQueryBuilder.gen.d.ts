import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class ResourceNamesQueryBuilder implements cog.Builder<azuremonitor.ResourceNamesQuery> {
    protected readonly internal: azuremonitor.ResourceNamesQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.ResourceNamesQuery;
    rawQuery(rawQuery: string): this;
    subscription(subscription: string): this;
    resourceGroup(resourceGroup: string): this;
    metricNamespace(metricNamespace: string): this;
}
