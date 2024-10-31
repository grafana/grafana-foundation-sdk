// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MetricAggregationWithFieldBuilder implements cog.Builder<elasticsearch.MetricAggregationWithField> {
    protected readonly internal: elasticsearch.MetricAggregationWithField;

    constructor() {
        this.internal = elasticsearch.defaultMetricAggregationWithField();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MetricAggregationWithField {
        return this.internal;
    }

    field(field: string): this {
        this.internal.field = field;
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
