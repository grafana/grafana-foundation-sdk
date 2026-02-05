import * as cog from '../cog';
import * as tempo from '../tempo';
import * as common from '../common';
export declare class TempoQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: tempo.TempoQuery;
    constructor();
    /**
     * Builds the object.
     */
    build(): tempo.TempoQuery;
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
    datasource(datasource: common.DataSourceRef): this;
    metricsQueryType(metricsQueryType: tempo.MetricsQueryType): this;
}
