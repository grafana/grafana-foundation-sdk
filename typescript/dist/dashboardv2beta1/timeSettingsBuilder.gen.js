"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeSettingsBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Time configuration
// It defines the default time config for the time picker, the refresh picker for the specific dashboard.
class TimeSettingsBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultTimeSettingsSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    timezone(timezone) {
        this.internal.timezone = timezone;
        return this;
    }
    // Start time range for dashboard.
    // Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    from(from) {
        this.internal.from = from;
        return this;
    }
    // End time range for dashboard.
    // Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    to(to) {
        this.internal.to = to;
        return this;
    }
    // Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    // v1: refresh
    autoRefresh(autoRefresh) {
        this.internal.autoRefresh = autoRefresh;
        return this;
    }
    // Interval options available in the refresh picker dropdown.
    // v1: timepicker.refresh_intervals
    autoRefreshIntervals(autoRefreshIntervals) {
        this.internal.autoRefreshIntervals = autoRefreshIntervals;
        return this;
    }
    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    // v1: timepicker.quick_ranges , not exposed in the UI
    quickRanges(quickRanges) {
        const quickRangesResources = quickRanges.map(builder1 => builder1.build());
        this.internal.quickRanges = quickRangesResources;
        return this;
    }
    // Whether timepicker is visible or not.
    // v1: timepicker.hidden
    hideTimepicker(hideTimepicker) {
        this.internal.hideTimepicker = hideTimepicker;
        return this;
    }
    // Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    weekStart(weekStart) {
        this.internal.weekStart = weekStart;
        return this;
    }
    // The month that the fiscal year starts on. 0 = January, 11 = December
    fiscalYearStartMonth(fiscalYearStartMonth) {
        this.internal.fiscalYearStartMonth = fiscalYearStartMonth;
        return this;
    }
    // Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    // v1: timepicker.nowDelay
    nowDelay(nowDelay) {
        this.internal.nowDelay = nowDelay;
        return this;
    }
}
exports.TimeSettingsBuilder = TimeSettingsBuilder;
//# sourceMappingURL=timeSettingsBuilder.gen.js.map