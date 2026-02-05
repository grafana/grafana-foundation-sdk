"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConnectionCoordinatesBuilder = void 0;
const tslib_1 = require("tslib");
const canvas = tslib_1.__importStar(require("../canvas"));
class ConnectionCoordinatesBuilder {
    constructor() {
        this.internal = canvas.defaultConnectionCoordinates();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    x(x) {
        this.internal.x = x;
        return this;
    }
    y(y) {
        this.internal.y = y;
        return this;
    }
}
exports.ConnectionCoordinatesBuilder = ConnectionCoordinatesBuilder;
//# sourceMappingURL=connectionCoordinatesBuilder.gen.js.map