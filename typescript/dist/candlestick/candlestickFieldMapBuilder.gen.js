"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CandlestickFieldMapBuilder = void 0;
const tslib_1 = require("tslib");
const candlestick = tslib_1.__importStar(require("../candlestick"));
class CandlestickFieldMapBuilder {
    constructor() {
        this.internal = candlestick.defaultCandlestickFieldMap();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Corresponds to the starting value of the given period
    open(open) {
        this.internal.open = open;
        return this;
    }
    // Corresponds to the highest value of the given period
    high(high) {
        this.internal.high = high;
        return this;
    }
    // Corresponds to the lowest value of the given period
    low(low) {
        this.internal.low = low;
        return this;
    }
    // Corresponds to the final (end) value of the given period
    close(close) {
        this.internal.close = close;
        return this;
    }
    // Corresponds to the sample count in the given period. (e.g. number of trades)
    volume(volume) {
        this.internal.volume = volume;
        return this;
    }
}
exports.CandlestickFieldMapBuilder = CandlestickFieldMapBuilder;
//# sourceMappingURL=candlestickFieldMapBuilder.gen.js.map