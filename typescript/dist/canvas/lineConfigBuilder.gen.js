"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.LineConfigBuilder = void 0;
const tslib_1 = require("tslib");
const canvas = tslib_1.__importStar(require("../canvas"));
class LineConfigBuilder {
    constructor() {
        this.internal = canvas.defaultLineConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    color(color) {
        const colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }
    width(width) {
        this.internal.width = width;
        return this;
    }
    radius(radius) {
        this.internal.radius = radius;
        return this;
    }
}
exports.LineConfigBuilder = LineConfigBuilder;
//# sourceMappingURL=lineConfigBuilder.gen.js.map