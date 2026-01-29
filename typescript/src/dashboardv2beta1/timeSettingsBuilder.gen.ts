// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Time configuration
// It defines the default time config for the time picker, the refresh picker for the specific dashboard.
export class TimeSettingsBuilder implements cog.Builder<dashboardv2beta1.TimeSettingsSpec> {
    protected readonly internal: dashboardv2beta1.TimeSettingsSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultTimeSettingsSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TimeSettingsSpec {
        return this.internal;
    }

    // Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    timezone(timezone: string): this {
        this.internal.timezone = timezone;
        return this;
    }

    // Start time range for dashboard.
    // Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    from(from: string): this {
        this.internal.from = from;
        return this;
    }

    // End time range for dashboard.
    // Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    to(to: string): this {
        this.internal.to = to;
        return this;
    }

    // Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    // v1: refresh
    autoRefresh(autoRefresh: string): this {
        this.internal.autoRefresh = autoRefresh;
        return this;
    }

    // Interval options available in the refresh picker dropdown.
    // v1: timepicker.refresh_intervals
    autoRefreshIntervals(autoRefreshIntervals: string[]): this {
        this.internal.autoRefreshIntervals = autoRefreshIntervals;
        return this;
    }

    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    // v1: timepicker.quick_ranges , not exposed in the UI
    quickRanges(quickRanges: cog.Builder<dashboardv2beta1.TimeRangeOption>[]): this {
        const quickRangesResources = quickRanges.map(builder1 => builder1.build());
        this.internal.quickRanges = quickRangesResources;
        return this;
    }

    // Whether timepicker is visible or not.
    // v1: timepicker.hidden
    hideTimepicker(hideTimepicker: boolean): this {
        this.internal.hideTimepicker = hideTimepicker;
        return this;
    }

    // Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    weekStart(weekStart: "saturday" | "monday" | "sunday"): this {
        this.internal.weekStart = weekStart;
        return this;
    }

    // The month that the fiscal year starts on. 0 = January, 11 = December
    fiscalYearStartMonth(fiscalYearStartMonth: number): this {
        this.internal.fiscalYearStartMonth = fiscalYearStartMonth;
        return this;
    }

    // Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    // v1: timepicker.nowDelay
    nowDelay(nowDelay: string): this {
        this.internal.nowDelay = nowDelay;
        return this;
    }
}

