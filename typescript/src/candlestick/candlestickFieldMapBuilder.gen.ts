// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as candlestick from '../candlestick';

export class CandlestickFieldMapBuilder implements cog.Builder<candlestick.CandlestickFieldMap> {
    protected readonly internal: candlestick.CandlestickFieldMap;

    constructor() {
        this.internal = candlestick.defaultCandlestickFieldMap();
    }

    /**
     * Builds the object.
     */
    build(): candlestick.CandlestickFieldMap {
        return this.internal;
    }

    // Corresponds to the starting value of the given period
    open(open: string): this {
        this.internal.open = open;
        return this;
    }

    // Corresponds to the highest value of the given period
    high(high: string): this {
        this.internal.high = high;
        return this;
    }

    // Corresponds to the lowest value of the given period
    low(low: string): this {
        this.internal.low = low;
        return this;
    }

    // Corresponds to the final (end) value of the given period
    close(close: string): this {
        this.internal.close = close;
        return this;
    }

    // Corresponds to the sample count in the given period. (e.g. number of trades)
    volume(volume: string): this {
        this.internal.volume = volume;
        return this;
    }
}
