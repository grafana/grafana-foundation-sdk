"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConstraintBuilder = void 0;
const tslib_1 = require("tslib");
const canvas = tslib_1.__importStar(require("../canvas"));
class ConstraintBuilder {
    constructor() {
        this.internal = canvas.defaultConstraint();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    horizontal(horizontal) {
        this.internal.horizontal = horizontal;
        return this;
    }
    vertical(vertical) {
        this.internal.vertical = vertical;
        return this;
    }
}
exports.ConstraintBuilder = ConstraintBuilder;
//# sourceMappingURL=constraintBuilder.gen.js.map