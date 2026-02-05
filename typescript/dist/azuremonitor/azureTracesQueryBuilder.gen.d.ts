import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class AzureTracesQueryBuilder implements cog.Builder<azuremonitor.AzureTracesQuery> {
    protected readonly internal: azuremonitor.AzureTracesQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureTracesQuery;
    resultFormat(resultFormat: azuremonitor.ResultFormat): this;
    resources(resources: string[]): this;
    operationId(operationId: string): this;
    traceTypes(traceTypes: string[]): this;
    filters(filters: cog.Builder<azuremonitor.AzureTracesFilter>[]): this;
    query(query: string): this;
}
