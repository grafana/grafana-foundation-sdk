import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class AppInsightsGroupByQueryBuilder implements cog.Builder<azuremonitor.AppInsightsGroupByQuery> {
    protected readonly internal: azuremonitor.AppInsightsGroupByQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.AppInsightsGroupByQuery;
    rawQuery(rawQuery: string): this;
    metricName(metricName: string): this;
}
