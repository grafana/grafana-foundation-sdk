// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Time picker configuration
// It defines the default config for the time picker and the refresh picker for the specific dashboard.
export class TimePickerBuilder implements cog.Builder<dashboard.TimePickerConfig> {
    protected readonly internal: dashboard.TimePickerConfig;

    constructor() {
        this.internal = dashboard.defaultTimePickerConfig();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.TimePickerConfig {
        return this.internal;
    }

    // Whether timepicker is visible or not.
    hidden(hidden: boolean): this {
        this.internal.hidden = hidden;
        return this;
    }

    // Interval options available in the refresh picker dropdown.
    refreshIntervals(refreshIntervals: string[]): this {
        this.internal.refresh_intervals = refreshIntervals;
        return this;
    }

    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    timeOptions(timeOptions: string[]): this {
        this.internal.time_options = timeOptions;
        return this;
    }
}
