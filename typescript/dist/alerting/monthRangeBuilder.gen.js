"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MonthRangeBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class MonthRangeBuilder {
    constructor() {
        this.internal = alerting.defaultMonthRange();
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
exports.MonthRangeBuilder = MonthRangeBuilder;
//# sourceMappingURL=monthRangeBuilder.gen.js.map