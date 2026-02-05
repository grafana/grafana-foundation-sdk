"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ControlsOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const geomap = tslib_1.__importStar(require("../geomap"));
class ControlsOptionsBuilder {
    constructor() {
        this.internal = geomap.defaultControlsOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Zoom (upper left)
    showZoom(showZoom) {
        this.internal.showZoom = showZoom;
        return this;
    }
    // let the mouse wheel zoom
    mouseWheelZoom(mouseWheelZoom) {
        this.internal.mouseWheelZoom = mouseWheelZoom;
        return this;
    }
    // Lower right
    showAttribution(showAttribution) {
        this.internal.showAttribution = showAttribution;
        return this;
    }
    // Scale options
    showScale(showScale) {
        this.internal.showScale = showScale;
        return this;
    }
    // Show debug
    showDebug(showDebug) {
        this.internal.showDebug = showDebug;
        return this;
    }
    // Show measure
    showMeasure(showMeasure) {
        this.internal.showMeasure = showMeasure;
        return this;
    }
}
exports.ControlsOptionsBuilder = ControlsOptionsBuilder;
//# sourceMappingURL=controlsOptionsBuilder.gen.js.map