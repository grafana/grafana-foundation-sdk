// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MovingAverageHoltModelSettingsBuilder implements cog.Builder<elasticsearch.MovingAverageHoltModelSettings> {
    protected readonly internal: elasticsearch.MovingAverageHoltModelSettings;

    constructor() {
        this.internal = elasticsearch.defaultMovingAverageHoltModelSettings();
        this.internal.model = "holt";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MovingAverageHoltModelSettings {
        return this.internal;
    }

    settings(settings: {
	alpha?: string;
	beta?: string;
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
