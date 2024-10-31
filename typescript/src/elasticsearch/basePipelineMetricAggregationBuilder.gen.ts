// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class BasePipelineMetricAggregationBuilder implements cog.Builder<elasticsearch.BasePipelineMetricAggregation> {
    protected readonly internal: elasticsearch.BasePipelineMetricAggregation;

    constructor() {
        this.internal = elasticsearch.defaultBasePipelineMetricAggregation();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.BasePipelineMetricAggregation {
        return this.internal;
    }

    pipelineAgg(pipelineAgg: string): this {
        this.internal.pipelineAgg = pipelineAgg;
        return this;
    }

    field(field: string): this {
        this.internal.field = field;
        return this;
    }

    type(type: string): this {
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
