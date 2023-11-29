// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class CumulativeSumBuilder implements cog.Builder<elasticsearch.CumulativeSum> {
    private readonly internal: elasticsearch.CumulativeSum;

    constructor() {
        this.internal = elasticsearch.defaultCumulativeSum();
        this.internal.type = "cumulative_sum";
    }

    build(): elasticsearch.CumulativeSum {
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

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    settings(settings: {
	format?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
