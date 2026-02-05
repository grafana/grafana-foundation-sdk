"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeRangeBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
// Redefining this to avoid an import cycle
class TimeRangeBuilder {
    constructor() {
        this.internal = alerting.defaultTimeRange();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    from(from) {
        this.internal.from = from;
        return this;
    }
    to(to) {
        this.internal.to = to;
        return this;
    }
}
exports.TimeRangeBuilder = TimeRangeBuilder;
//# sourceMappingURL=timeRangeBuilder.gen.js.map