// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ThresholdsConfigBuilder implements cog.Builder<dashboardv2.ThresholdsConfig> {
    protected readonly internal: dashboardv2.ThresholdsConfig;

    constructor() {
        this.internal = dashboardv2.defaultThresholdsConfig();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.ThresholdsConfig {
        return this.internal;
    }

    mode(mode: dashboardv2.ThresholdsMode): this {
        this.internal.mode = mode;
        return this;
    }

    steps(steps: dashboardv2.Threshold[]): this {
        this.internal.steps = steps;
        return this;
    }
}

