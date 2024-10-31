// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

export class TimePickerBuilder implements cog.Builder<dashboard.TimePicker> {
    protected readonly internal: dashboard.TimePicker;

    constructor() {
        this.internal = dashboard.defaultTimePicker();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.TimePicker {
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

    // Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
    collapse(collapse: boolean): this {
        this.internal.collapse = collapse;
        return this;
    }

    // Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
    enable(enable: boolean): this {
        this.internal.enable = enable;
        return this;
    }

    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    timeOptions(timeOptions: string[]): this {
        this.internal.time_options = timeOptions;
        return this;
    }
}
