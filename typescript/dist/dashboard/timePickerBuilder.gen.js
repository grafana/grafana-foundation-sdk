"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimePickerBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// Time picker configuration
// It defines the default config for the time picker and the refresh picker for the specific dashboard.
class TimePickerBuilder {
    constructor() {
        this.internal = dashboard.defaultTimePickerConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Whether timepicker is visible or not.
    hidden(hidden) {
        this.internal.hidden = hidden;
        return this;
    }
    // Interval options available in the refresh picker dropdown.
    refreshIntervals(refreshIntervals) {
        this.internal.refresh_intervals = refreshIntervals;
        return this;
    }
    // Quick ranges for time picker.
    quickRanges(quickRanges) {
        const quickRangesResources = quickRanges.map(builder1 => builder1.build());
        this.internal.quick_ranges = quickRangesResources;
        return this;
    }
    // Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    nowDelay(nowDelay) {
        this.internal.nowDelay = nowDelay;
        return this;
    }
}
exports.TimePickerBuilder = TimePickerBuilder;
//# sourceMappingURL=timePickerBuilder.gen.js.map