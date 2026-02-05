"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.PlacementBuilder = void 0;
const tslib_1 = require("tslib");
const canvas = tslib_1.__importStar(require("../canvas"));
class PlacementBuilder {
    constructor() {
        this.internal = canvas.defaultPlacement();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    top(top) {
        this.internal.top = top;
        return this;
    }
    left(left) {
        this.internal.left = left;
        return this;
    }
    right(right) {
        this.internal.right = right;
        return this;
    }
    bottom(bottom) {
        this.internal.bottom = bottom;
        return this;
    }
    width(width) {
        this.internal.width = width;
        return this;
    }
    height(height) {
        this.internal.height = height;
        return this;
    }
    rotation(rotation) {
        this.internal.rotation = rotation;
        return this;
    }
}
exports.PlacementBuilder = PlacementBuilder;
//# sourceMappingURL=placementBuilder.gen.js.map