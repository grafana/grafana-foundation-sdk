"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.LogGroupBuilder = void 0;
const tslib_1 = require("tslib");
const cloudwatch = tslib_1.__importStar(require("../cloudwatch"));
class LogGroupBuilder {
    constructor() {
        this.internal = cloudwatch.defaultLogGroup();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // ARN of the log group
    arn(arn) {
        this.internal.arn = arn;
        return this;
    }
    // Name of the log group
    name(name) {
        this.internal.name = name;
        return this;
    }
    // AccountId of the log group
    accountId(accountId) {
        this.internal.accountId = accountId;
        return this;
    }
    // Label of the log group
    accountLabel(accountLabel) {
        this.internal.accountLabel = accountLabel;
        return this;
    }
}
exports.LogGroupBuilder = LogGroupBuilder;
//# sourceMappingURL=logGroupBuilder.gen.js.map