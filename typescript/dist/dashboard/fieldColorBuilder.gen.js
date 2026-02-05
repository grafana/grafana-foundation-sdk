"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.FieldColorBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// Map a field to a color.
class FieldColorBuilder {
    constructor() {
        this.internal = dashboard.defaultFieldColor();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // The main color scheme mode.
    mode(mode) {
        this.internal.mode = mode;
        return this;
    }
    // The fixed color value for fixed or shades color modes.
    fixedColor(fixedColor) {
        this.internal.fixedColor = fixedColor;
        return this;
    }
    // Some visualizations need to know how to assign a series color from by value color schemes.
    seriesBy(seriesBy) {
        this.internal.seriesBy = seriesBy;
        return this;
    }
}
exports.FieldColorBuilder = FieldColorBuilder;
//# sourceMappingURL=fieldColorBuilder.gen.js.map