// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class BaseMetricAggregationBuilder implements cog.Builder<elasticsearch.BaseMetricAggregation> {
    protected readonly internal: elasticsearch.BaseMetricAggregation;

    constructor() {
        this.internal = elasticsearch.defaultBaseMetricAggregation();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.BaseMetricAggregation {
        return this.internal;
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
