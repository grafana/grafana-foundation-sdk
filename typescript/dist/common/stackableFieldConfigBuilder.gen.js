"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.StackableFieldConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class StackableFieldConfigBuilder {
    constructor() {
        this.internal = common.defaultStackableFieldConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    stacking(stacking) {
        const stackingResource = stacking.build();
        this.internal.stacking = stackingResource;
        return this;
    }
}
exports.StackableFieldConfigBuilder = StackableFieldConfigBuilder;
//# sourceMappingURL=stackableFieldConfigBuilder.gen.js.map