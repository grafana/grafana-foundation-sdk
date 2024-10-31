// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MovingAverageEWMAModelSettingsBuilder implements cog.Builder<elasticsearch.MovingAverageEWMAModelSettings> {
    protected readonly internal: elasticsearch.MovingAverageEWMAModelSettings;

    constructor() {
        this.internal = elasticsearch.defaultMovingAverageEWMAModelSettings();
        this.internal.model = "ewma";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MovingAverageEWMAModelSettings {
        return this.internal;
    }

    settings(settings: {
	alpha?: string;
}): this {
        this.internal.settings = settings;
        return this;
    }

    window(window: string): this {
        this.internal.window = window;
        return this;
    }

    minimize(minimize: boolean): this {
        this.internal.minimize = minimize;
        return this;
    }

    predict(predict: string): this {
        this.internal.predict = predict;
        return this;
    }
}
