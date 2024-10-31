// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as elasticsearch from '../elasticsearch';

export class MovingAverageHoltWintersModelSettingsBuilder implements cog.Builder<elasticsearch.MovingAverageHoltWintersModelSettings> {
    protected readonly internal: elasticsearch.MovingAverageHoltWintersModelSettings;

    constructor() {
        this.internal = elasticsearch.defaultMovingAverageHoltWintersModelSettings();
        this.internal.model = "holt_winters";
    }

    /**
     * Builds the object.
     */
    build(): elasticsearch.MovingAverageHoltWintersModelSettings {
        return this.internal;
    }

    settings(settings: {
	alpha?: string;
	beta?: string;
	gamma?: string;
	period?: string;
	pad?: boolean;
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
