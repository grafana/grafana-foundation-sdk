"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ScaleDistributionConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class ScaleDistributionConfigBuilder {
    constructor() {
        this.internal = common.defaultScaleDistributionConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    type(type) {
        this.internal.type = type;
        return this;
    }
    log(log) {
        this.internal.log = log;
        return this;
    }
    linearThreshold(linearThreshold) {
        this.internal.linearThreshold = linearThreshold;
        return this;
    }
}
exports.ScaleDistributionConfigBuilder = ScaleDistributionConfigBuilder;
//# sourceMappingURL=scaleDistributionConfigBuilder.gen.js.map