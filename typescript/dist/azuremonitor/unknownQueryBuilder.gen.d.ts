import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class UnknownQueryBuilder implements cog.Builder<azuremonitor.UnknownQuery> {
    protected readonly internal: azuremonitor.UnknownQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.UnknownQuery;
    rawQuery(rawQuery: string): this;
}
