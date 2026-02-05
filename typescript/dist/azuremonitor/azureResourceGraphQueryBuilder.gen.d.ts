import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class AzureResourceGraphQueryBuilder implements cog.Builder<azuremonitor.AzureResourceGraphQuery> {
    protected readonly internal: azuremonitor.AzureResourceGraphQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureResourceGraphQuery;
    query(query: string): this;
    resultFormat(resultFormat: string): this;
}
