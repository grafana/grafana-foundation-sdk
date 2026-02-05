"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.RecordRuleBuilder = void 0;
const tslib_1 = require("tslib");
const alerting = tslib_1.__importStar(require("../alerting"));
class RecordRuleBuilder {
    constructor() {
        this.internal = alerting.defaultRecordRule();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Which expression node should be used as the input for the recorded metric.
    from(from) {
        this.internal.from = from;
        return this;
    }
    // Name of the recorded metric.
    metric(metric) {
        this.internal.metric = metric;
        return this;
    }
}
exports.RecordRuleBuilder = RecordRuleBuilder;
//# sourceMappingURL=recordRuleBuilder.gen.js.map