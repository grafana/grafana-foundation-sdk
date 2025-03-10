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

    // Quick ranges for time picker.
    quickRanges(quickRanges: cog.Builder<dashboard.TimeOption>[]): this {
        const quickRangesResources = quickRanges.map(builder1 => builder1.build());
        this.internal.quick_ranges = quickRangesResources;
        return this;
    }

    // Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    nowDelay(nowDelay: string): this {
        this.internal.nowDelay = nowDelay;
        return this;
    }
}
