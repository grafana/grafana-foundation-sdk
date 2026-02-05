"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.WeekdayRangeBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class WeekdayRangeBuilder {
    constructor() {
        this.internal = alerting.defaultWeekdayRange();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    begin(begin) {
        this.internal.begin = begin;
        return this;
    }
    end(end) {
        this.internal.end = end;
        return this;
    }
}
exports.WeekdayRangeBuilder = WeekdayRangeBuilder;
//# sourceMappingURL=weekdayRangeBuilder.gen.js.map