"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.HideableFieldConfigBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// TODO docs
class HideableFieldConfigBuilder {
    constructor() {
        this.internal = common.defaultHideableFieldConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    hideFrom(hideFrom) {
        const hideFromResource = hideFrom.build();
        this.internal.hideFrom = hideFromResource;
        return this;
    }
}
exports.HideableFieldConfigBuilder = HideableFieldConfigBuilder;
//# sourceMappingURL=hideableFieldConfigBuilder.gen.js.map