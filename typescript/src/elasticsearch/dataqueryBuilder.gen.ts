// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';
import * as dashboard from '../dashboard';

export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: elasticsearch.dataquery;

    constructor() {
        this.internal = elasticsearch.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.dataquery {
        return this.internal;
    }

    // Alias pattern
    alias(alias: string): this {
        this.internal.alias = alias;
        return this;
    }

    // Lucene query
    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    // Name of time field
    timeField(timeField: string): this {
        this.internal.timeField = timeField;
        return this;
    }

    // List of bucket aggregations
    bucketAggs(bucketAggs: cog.Builder<elasticsearch.BucketAggregation>[]): this {
        const bucketAggsResources = bucketAggs.map(builder1 => builder1.build());
        this.internal.bucketAggs = bucketAggsResources;
        return this;
    }

    // List of metric aggregations
    metrics(metrics: cog.Builder<elasticsearch.MetricAggregation>[]): this {
        const metricsResources = metrics.map(builder1 => builder1.build());
        this.internal.metrics = metricsResources;
        return this;
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

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }
}

