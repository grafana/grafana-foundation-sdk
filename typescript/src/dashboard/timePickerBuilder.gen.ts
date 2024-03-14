// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Time picker configuration
// It defines the default config for the time picker and the refresh picker for the specific dashboard.
export class TimePickerBuilder implements cog.Builder<dashboard.TimePickerConfig> {
    private readonly internal: dashboard.TimePickerConfig;

    constructor() {
        this.internal = dashboard.defaultTimePickerConfig();
    }

    build(): dashboard.TimePickerConfig {
        return this.internal;
    }

    // Whether timepicker is visible or not.
    hidden(hidden: boolean): this {
        this.internal.hidden = hidden;
        return this;
    }

    // Interval options available in the refresh picker dropdown.
    refresh_intervals(refresh_intervals: string[]): this {
        this.internal.refresh_intervals = refresh_intervals;
        return this;
    }

    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    time_options(time_options: string[]): this {
        this.internal.time_options = time_options;
        return this;
    }
}
