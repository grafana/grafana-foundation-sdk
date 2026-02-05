"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeIntervalBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class TimeIntervalBuilder {
    constructor() {
        this.internal = alerting.defaultTimeInterval();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    times(times) {
        const timesResources = times.map(builder1 => builder1.build());
        this.internal.times = timesResources;
        return this;
    }
    weekdays(weekdays) {
        const weekdaysResources = weekdays.map(builder1 => builder1.build());
        this.internal.weekdays = weekdaysResources;
        return this;
    }
    daysOfMonth(daysOfMonth) {
        const daysOfMonthResources = daysOfMonth.map(builder1 => builder1.build());
        this.internal.days_of_month = daysOfMonthResources;
        return this;
    }
    months(months) {
        const monthsResources = months.map(builder1 => builder1.build());
        this.internal.months = monthsResources;
        return this;
    }
    years(years) {
        const yearsResources = years.map(builder1 => builder1.build());
        this.internal.years = yearsResources;
        return this;
    }
    location(location) {
        this.internal.location = location;
        return this;
    }
}
exports.TimeIntervalBuilder = TimeIntervalBuilder;
//# sourceMappingURL=timeIntervalBuilder.gen.js.map