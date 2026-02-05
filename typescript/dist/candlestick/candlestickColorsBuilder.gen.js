"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CandlestickColorsBuilder = void 0;
const tslib_1 = require("tslib");
const candlestick = tslib_1.__importStar(require("../candlestick"));
class CandlestickColorsBuilder {
    constructor() {
        this.internal = candlestick.defaultCandlestickColors();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    up(up) {
        this.internal.up = up;
        return this;
    }
    down(down) {
        this.internal.down = down;
        return this;
    }
    flat(flat) {
        this.internal.flat = flat;
        return this;
    }
}
exports.CandlestickColorsBuilder = CandlestickColorsBuilder;
//# sourceMappingURL=candlestickColorsBuilder.gen.js.map