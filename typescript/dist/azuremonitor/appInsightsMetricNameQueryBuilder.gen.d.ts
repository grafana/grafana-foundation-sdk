import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class AppInsightsMetricNameQueryBuilder implements cog.Builder<azuremonitor.AppInsightsMetricNameQuery> {
    protected readonly internal: azuremonitor.AppInsightsMetricNameQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.AppInsightsMetricNameQuery;
    rawQuery(rawQuery: string): this;
}
