"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.SimulationQueryBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class SimulationQueryBuilder {
    constructor() {
        this.internal = testdata.defaultSimulationQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    config(config) {
        this.internal.config = config;
        return this;
    }
    key(key) {
        const keyResource = key.build();
        this.internal.key = keyResource;
        return this;
    }
    last(last) {
        this.internal.last = last;
        return this;
    }
    stream(stream) {
        this.internal.stream = stream;
        return this;
    }
}
exports.SimulationQueryBuilder = SimulationQueryBuilder;
//# sourceMappingURL=simulationQueryBuilder.gen.js.map