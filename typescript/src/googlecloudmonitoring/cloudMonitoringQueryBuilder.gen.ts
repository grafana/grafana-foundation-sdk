// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';
import * as dashboard from '../dashboard';

export class CloudMonitoringQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: googlecloudmonitoring.CloudMonitoringQuery;

    constructor() {
        this.internal = googlecloudmonitoring.defaultCloudMonitoringQuery();
    }

    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.CloudMonitoringQuery {
        return this.internal;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    aliasBy(aliasBy: string): this {
        this.internal.aliasBy = aliasBy;
        return this;
    }

    // GCM query type.
    // queryType: #QueryType
    // Time Series List sub-query properties.
    timeSeriesList(timeSeriesList: cog.Builder<googlecloudmonitoring.TimeSeriesList>): this {
        const timeSeriesListResource = timeSeriesList.build();
        this.internal.timeSeriesList = timeSeriesListResource;
        return this;
    }

    // Time Series sub-query properties.
    timeSeriesQuery(timeSeriesQuery: cog.Builder<googlecloudmonitoring.TimeSeriesQuery>): this {
        const timeSeriesQueryResource = timeSeriesQuery.build();
        this.internal.timeSeriesQuery = timeSeriesQueryResource;
        return this;
    }

    // SLO sub-query properties.
    sloQuery(sloQuery: cog.Builder<googlecloudmonitoring.SLOQuery>): this {
        const sloQueryResource = sloQuery.build();
        this.internal.sloQuery = sloQueryResource;
        return this;
    }

    // PromQL sub-query properties.
    promQLQuery(promQLQuery: cog.Builder<googlecloudmonitoring.PromQLQuery>): this {
        const promQLQueryResource = promQLQuery.build();
        this.internal.promQLQuery = promQLQueryResource;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // Time interval in milliseconds.
    intervalMs(intervalMs: number): this {
        this.internal.intervalMs = intervalMs;
        return this;
    }
}
