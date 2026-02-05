"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.PulseWaveQueryBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class PulseWaveQueryBuilder {
    constructor() {
        this.internal = testdata.defaultPulseWaveQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    offCount(offCount) {
        this.internal.offCount = offCount;
        return this;
    }
    offValue(offValue) {
        this.internal.offValue = offValue;
        return this;
    }
    onCount(onCount) {
        this.internal.onCount = onCount;
        return this;
    }
    onValue(onValue) {
        this.internal.onValue = onValue;
        return this;
    }
    timeStep(timeStep) {
        this.internal.timeStep = timeStep;
        return this;
    }
}
exports.PulseWaveQueryBuilder = PulseWaveQueryBuilder;
//# sourceMappingURL=pulseWaveQueryBuilder.gen.js.map