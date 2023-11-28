// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MovingAverageModelOptionBuilder implements cog.OptionsBuilder<elasticsearch.MovingAverageModelOption> {
    private readonly internal: elasticsearch.MovingAverageModelOption;

    constructor() {
        this.internal = elasticsearch.defaultMovingAverageModelOption();
    }

    build(): elasticsearch.MovingAverageModelOption {
        return this.internal;
    }

    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    value(value: elasticsearch.MovingAverageModel): this {
        this.internal.value = value;
        return this;
    }
}
