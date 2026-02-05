"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MuteTimingBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class MuteTimingBuilder {
    constructor() {
        this.internal = alerting.defaultMuteTiming();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
    timeIntervals(timeIntervals) {
        const timeIntervalsResources = timeIntervals.map(builder1 => builder1.build());
        this.internal.time_intervals = timeIntervalsResources;
        return this;
    }
}
exports.MuteTimingBuilder = MuteTimingBuilder;
//# sourceMappingURL=muteTimingBuilder.gen.js.map