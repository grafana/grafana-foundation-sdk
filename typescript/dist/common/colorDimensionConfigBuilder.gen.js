"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ColorDimensionConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
class ColorDimensionConfigBuilder {
    constructor() {
        this.internal = common.defaultColorDimensionConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // color value
    fixed(fixed) {
        this.internal.fixed = fixed;
        return this;
    }
    // fixed: T -- will be added by each element
    field(field) {
        this.internal.field = field;
        return this;
    }
}
exports.ColorDimensionConfigBuilder = ColorDimensionConfigBuilder;
//# sourceMappingURL=colorDimensionConfigBuilder.gen.js.map