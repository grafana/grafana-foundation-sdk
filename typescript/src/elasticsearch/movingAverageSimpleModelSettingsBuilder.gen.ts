// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MovingAverageSimpleModelSettingsBuilder implements cog.Builder<elasticsearch.MovingAverageSimpleModelSettings> {
    protected readonly internal: elasticsearch.MovingAverageSimpleModelSettings;

    constructor() {
        this.internal = elasticsearch.defaultMovingAverageSimpleModelSettings();
        this.internal.model = "simple";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MovingAverageSimpleModelSettings {
        return this.internal;
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
