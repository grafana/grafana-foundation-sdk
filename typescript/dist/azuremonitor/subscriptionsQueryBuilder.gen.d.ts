import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class SubscriptionsQueryBuilder implements cog.Builder<azuremonitor.SubscriptionsQuery> {
    protected readonly internal: azuremonitor.SubscriptionsQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.SubscriptionsQuery;
    rawQuery(rawQuery: string): this;
}
