// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class BaseMovingAverageModelSettingsBuilder implements cog.Builder<elasticsearch.BaseMovingAverageModelSettings> {
    protected readonly internal: elasticsearch.BaseMovingAverageModelSettings;

    constructor() {
        this.internal = elasticsearch.defaultBaseMovingAverageModelSettings();
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.BaseMovingAverageModelSettings {
        return this.internal;
    }

    model(model: elasticsearch.MovingAverageModel): this {
        this.internal.model = model;
        return this;
    }

    window(window: string): this {
        this.internal.window = window;
        return this;
    }

    predict(predict: string): this {
        this.internal.predict = predict;
        return this;
    }
}
