"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.USAQueryBuilder = void 0;
const tslib_1 = require("tslib");
const testdata = tslib_1.__importStar(require("../testdata"));
class USAQueryBuilder {
    constructor() {
        this.internal = testdata.defaultUSAQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    fields(fields) {
        this.internal.fields = fields;
        return this;
    }
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    period(period) {
        this.internal.period = period;
        return this;
    }
    states(states) {
        this.internal.states = states;
        return this;
    }
}
exports.USAQueryBuilder = USAQueryBuilder;
//# sourceMappingURL=uSAQueryBuilder.gen.js.map