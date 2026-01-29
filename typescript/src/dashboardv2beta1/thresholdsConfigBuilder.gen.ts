// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ThresholdsConfigBuilder implements cog.Builder<dashboardv2beta1.ThresholdsConfig> {
    protected readonly internal: dashboardv2beta1.ThresholdsConfig;

    constructor() {
        this.internal = dashboardv2beta1.defaultThresholdsConfig();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ThresholdsConfig {
        return this.internal;
    }

    mode(mode: dashboardv2beta1.ThresholdsMode): this {
        this.internal.mode = mode;
        return this;
    }

    steps(steps: dashboardv2beta1.Threshold[]): this {
        this.internal.steps = steps;
        return this;
    }
}

