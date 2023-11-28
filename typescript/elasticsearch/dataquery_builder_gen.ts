// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class DataqueryBuilder implements cog.OptionsBuilder<cog.Dataquery> {
    private readonly internal: elasticsearch.dataquery;

    constructor() {
        this.internal = elasticsearch.defaultDataquery();
    }

    build(): elasticsearch.dataquery {
        return this.internal;
    }

    alias(alias: string): this {
        this.internal.alias = alias;
        return this;
    }

    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    timeField(timeField: string): this {
        this.internal.timeField = timeField;
        return this;
    }

    bucketAggs(bucketAggs: elasticsearch.BucketAggregation[]): this {
        this.internal.bucketAggs = bucketAggs;
        return this;
    }

    metrics(metrics: elasticsearch.MetricAggregation[]): this {
        this.internal.metrics = metrics;
        return this;
    }

    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    datasource(datasource: any): this {
        this.internal.datasource = datasource;
        return this;
    }
}
