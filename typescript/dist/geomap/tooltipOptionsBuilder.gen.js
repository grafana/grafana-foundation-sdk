"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TooltipOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const geomap = tslib_1.__importStar(require("../geomap"));
class TooltipOptionsBuilder {
    constructor() {
        this.internal = geomap.defaultTooltipOptions();
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
}
exports.TooltipOptionsBuilder = TooltipOptionsBuilder;
//# sourceMappingURL=tooltipOptionsBuilder.gen.js.map