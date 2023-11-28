// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

// #MovingAverage's settings are overridden in types.ts
export class MovingAverageBuilder implements cog.OptionsBuilder<elasticsearch.MovingAverage> {
    private readonly internal: elasticsearch.MovingAverage;

    constructor() {
        this.internal = elasticsearch.defaultMovingAverage();
        this.internal.type = "moving_avg";
    }

    build(): elasticsearch.MovingAverage {
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

    settings(settings: any): this {
        this.internal.settings = settings;
        return this;
    }

    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }
}
