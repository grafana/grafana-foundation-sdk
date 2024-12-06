// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MovingAverageLinearModelSettingsBuilder implements cog.Builder<elasticsearch.MovingAverageLinearModelSettings> {
    protected readonly internal: elasticsearch.MovingAverageLinearModelSettings;

    constructor() {
        this.internal = elasticsearch.defaultMovingAverageLinearModelSettings();
        this.internal.model = "linear";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MovingAverageLinearModelSettings {
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
