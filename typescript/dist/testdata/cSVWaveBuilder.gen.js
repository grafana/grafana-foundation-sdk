"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CSVWaveBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class CSVWaveBuilder {
    constructor() {
        this.internal = testdata.defaultCSVWave();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    labels(labels) {
        this.internal.labels = labels;
        return this;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
    timeStep(timeStep) {
        this.internal.timeStep = timeStep;
        return this;
    }
    valuesCSV(valuesCSV) {
        this.internal.valuesCSV = valuesCSV;
        return this;
    }
}
exports.CSVWaveBuilder = CSVWaveBuilder;
//# sourceMappingURL=cSVWaveBuilder.gen.js.map