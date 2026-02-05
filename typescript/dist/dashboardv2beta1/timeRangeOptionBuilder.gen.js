"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeRangeOptionBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class TimeRangeOptionBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultTimeRangeOption();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    display(display) {
        this.internal.display = display;
        return this;
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
exports.TimeRangeOptionBuilder = TimeRangeOptionBuilder;
//# sourceMappingURL=timeRangeOptionBuilder.gen.js.map