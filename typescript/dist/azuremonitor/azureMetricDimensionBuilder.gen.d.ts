import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class AzureMetricDimensionBuilder implements cog.Builder<azuremonitor.AzureMetricDimension> {
    protected readonly internal: azuremonitor.AzureMetricDimension;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureMetricDimension;
    dimension(dimension: string): this;
    operator(operator: string): this;
    filters(filters: string[]): this;
    filter(filter: string): this;
}
