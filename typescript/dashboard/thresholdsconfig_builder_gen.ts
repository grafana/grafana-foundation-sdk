// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Thresholds configuration for the panel
export class ThresholdsConfigBuilder implements cog.OptionsBuilder<dashboard.ThresholdsConfig> {
    private readonly internal: dashboard.ThresholdsConfig;

    constructor() {
        this.internal = dashboard.defaultThresholdsConfig();
    }

    build(): dashboard.ThresholdsConfig {
        return this.internal;
    }

    // Thresholds mode.
    mode(mode: dashboard.ThresholdsMode): this {
        this.internal.mode = mode;
        return this;
    }

    // Must be sorted by 'value', first value is always -Infinity
    steps(steps: dashboard.Threshold[]): this {
        this.internal.steps = steps;
        return this;
    }
}
