import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as tempo from '../tempo';
export declare class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind;
    version(version: string): this;
    datasource(ref: {
        name?: string;
    }): this;
    refId(refId: string): this;
    hide(hide: boolean): this;
    queryType(queryType: string): this;
    query(query: string): this;
    search(search: string): this;
    serviceName(serviceName: string): this;
    spanName(spanName: string): this;
    minDuration(minDuration: string): this;
    maxDuration(maxDuration: string): this;
    serviceMapQuery(serviceMapQuery: string | string[]): this;
    serviceMapIncludeNamespace(serviceMapIncludeNamespace: boolean): this;
    limit(limit: number): this;
    spss(spss: number): this;
    filters(filters: cog.Builder<tempo.TraceqlFilter>[]): this;
    groupBy(groupBy: cog.Builder<tempo.TraceqlFilter>[]): this;
    tableType(tableType: tempo.SearchTableType): this;
    step(step: string): this;
    exemplars(exemplars: number): this;
    metricsQueryType(metricsQueryType: tempo.MetricsQueryType): this;
}
