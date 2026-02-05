"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TextDimensionConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
class TextDimensionConfigBuilder {
    constructor() {
        this.internal = common.defaultTextDimensionConfig();
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
    // fixed: T -- will be added by each element
    field(field) {
        this.internal.field = field;
        return this;
    }
    fixed(fixed) {
        this.internal.fixed = fixed;
        return this;
    }
}
exports.TextDimensionConfigBuilder = TextDimensionConfigBuilder;
//# sourceMappingURL=textDimensionConfigBuilder.gen.js.map