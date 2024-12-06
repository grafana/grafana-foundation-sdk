// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MetricAggregationWithMissingSupportBuilder implements cog.Builder<elasticsearch.MetricAggregationWithMissingSupport> {
    protected readonly internal: elasticsearch.MetricAggregationWithMissingSupport;

    constructor() {
        this.internal = elasticsearch.defaultMetricAggregationWithMissingSupport();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MetricAggregationWithMissingSupport {
        return this.internal;
    }

    settings(settings: {
	missing?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    type(type: elasticsearch.MetricAggregationType): this {
        this.internal.type = type;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
