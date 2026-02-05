"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.StackingConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class StackingConfigBuilder {
    constructor() {
        this.internal = common.defaultStackingConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    group(group) {
        this.internal.group = group;
        return this;
    }
}
exports.StackingConfigBuilder = StackingConfigBuilder;
//# sourceMappingURL=stackingConfigBuilder.gen.js.map