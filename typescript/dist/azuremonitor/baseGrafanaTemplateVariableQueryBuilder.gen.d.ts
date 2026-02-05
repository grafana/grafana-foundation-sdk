import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';
export declare class BaseGrafanaTemplateVariableQueryBuilder implements cog.Builder<azuremonitor.BaseGrafanaTemplateVariableQuery> {
    protected readonly internal: azuremonitor.BaseGrafanaTemplateVariableQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): azuremonitor.BaseGrafanaTemplateVariableQuery;
    rawQuery(rawQuery: string): this;
}
