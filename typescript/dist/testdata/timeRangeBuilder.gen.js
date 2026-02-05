"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeRangeBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class TimeRangeBuilder {
    constructor() {
        this.internal = testdata.defaultTimeRange();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // From is the start time of the query.
    from(from) {
        this.internal.from = from;
        return this;
    }
    // To is the end time of the query.
    to(to) {
        this.internal.to = to;
        return this;
    }
}
exports.TimeRangeBuilder = TimeRangeBuilder;
//# sourceMappingURL=timeRangeBuilder.gen.js.map