"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ArcOptionBuilder = void 0;
const tslib_1 = require("tslib");
const nodegraph = tslib_1.__importStar(require("../nodegraph"));
class ArcOptionBuilder {
    constructor() {
        this.internal = nodegraph.defaultArcOption();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Field from which to get the value. Values should be less than 1, representing fraction of a circle.
    field(field) {
        this.internal.field = field;
        return this;
    }
    // The color of the arc.
    color(color) {
        this.internal.color = color;
        return this;
    }
}
exports.ArcOptionBuilder = ArcOptionBuilder;
//# sourceMappingURL=arcOptionBuilder.gen.js.map