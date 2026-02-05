// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as googlecloudmonitoring from '../googlecloudmonitoring';

export class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "cloud-monitoring";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref: {
	name?: string;
}): this {
        this.internal.datasource = ref;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }

    // Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    aliasBy(aliasBy: string): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        this.internal.spec.aliasBy = aliasBy;
        return this;
    }

    // GCM query type.
    // queryType: #QueryType
    // Time Series List sub-query properties.
    timeSeriesList(timeSeriesList: cog.Builder<googlecloudmonitoring.TimeSeriesList>): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        const timeSeriesListResource = timeSeriesList.build();
        this.internal.spec.timeSeriesList = timeSeriesListResource;
        return this;
    }

    // Time Series sub-query properties.
    timeSeriesQuery(timeSeriesQuery: cog.Builder<googlecloudmonitoring.TimeSeriesQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        const timeSeriesQueryResource = timeSeriesQuery.build();
        this.internal.spec.timeSeriesQuery = timeSeriesQueryResource;
        return this;
    }

    // SLO sub-query properties.
    sloQuery(sloQuery: cog.Builder<googlecloudmonitoring.SLOQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        const sloQueryResource = sloQuery.build();
        this.internal.spec.sloQuery = sloQueryResource;
        return this;
    }

    // PromQL sub-query properties.
    promQLQuery(promQLQuery: cog.Builder<googlecloudmonitoring.PromQLQuery>): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        const promQLQueryResource = promQLQuery.build();
        this.internal.spec.promQLQuery = promQLQueryResource;
        return this;
    }

    // Time interval in milliseconds.
    intervalMs(intervalMs: number): this {
        if (!this.internal.spec) {
            this.internal.spec = googlecloudmonitoring.defaultCloudMonitoringQuery();
        }
        this.internal.spec.intervalMs = intervalMs;
        return this;
    }
}

