import * as cog from '../cog';
import * as candlestick from '../candlestick';
export declare class CandlestickFieldMapBuilder implements cog.Builder<candlestick.CandlestickFieldMap> {
    protected readonly internal: candlestick.CandlestickFieldMap;
    constructor();
    /**
     * Builds the object.
     */
    build(): candlestick.CandlestickFieldMap;
    open(open: string): this;
    high(high: string): this;
    low(low: string): this;
    close(close: string): this;
    volume(volume: string): this;
}
