"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.XYSeriesConfigBuilder = void 0;
const tslib_1 = require("tslib");
const xychart = tslib_1.__importStar(require("../xychart"));
class XYSeriesConfigBuilder {
    constructor() {
        this.internal = xychart.defaultXYSeriesConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
    frame(frame) {
        this.internal.frame = frame;
        return this;
    }
    x(x) {
        this.internal.x = x;
        return this;
    }
    y(y) {
        this.internal.y = y;
        return this;
    }
    color(color) {
        this.internal.color = color;
        return this;
    }
    size(size) {
        this.internal.size = size;
        return this;
    }
}
exports.XYSeriesConfigBuilder = XYSeriesConfigBuilder;
//# sourceMappingURL=xYSeriesConfigBuilder.gen.js.map