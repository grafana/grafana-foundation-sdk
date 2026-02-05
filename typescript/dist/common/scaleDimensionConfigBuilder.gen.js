"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ScaleDimensionConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
class ScaleDimensionConfigBuilder {
    constructor() {
        this.internal = common.defaultScaleDimensionConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    min(min) {
        this.internal.min = min;
        return this;
    }
    max(max) {
        this.internal.max = max;
        return this;
    }
    fixed(fixed) {
        this.internal.fixed = fixed;
        return this;
    }
    // fixed: T -- will be added by each element
    field(field) {
        this.internal.field = field;
        return this;
    }
    // | *"linear"
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
}
exports.ScaleDimensionConfigBuilder = ScaleDimensionConfigBuilder;
//# sourceMappingURL=scaleDimensionConfigBuilder.gen.js.map