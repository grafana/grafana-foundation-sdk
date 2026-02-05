"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ValueMappingResultBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Result used as replacement with text and color when the value matches
class ValueMappingResultBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultValueMappingResult();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Text to display when the value matches
    text(text) {
        this.internal.text = text;
        return this;
    }
    // Text to use when the value matches
    color(color) {
        this.internal.color = color;
        return this;
    }
    // Icon to display when the value matches. Only specific visualizations.
    icon(icon) {
        this.internal.icon = icon;
        return this;
    }
    // Position in the mapping array. Only used internally.
    index(index) {
        this.internal.index = index;
        return this;
    }
}
exports.ValueMappingResultBuilder = ValueMappingResultBuilder;
//# sourceMappingURL=valueMappingResultBuilder.gen.js.map