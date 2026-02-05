import * as cog from '../cog';
import * as candlestick from '../candlestick';
export declare class CandlestickColorsBuilder implements cog.Builder<candlestick.CandlestickColors> {
    protected readonly internal: candlestick.CandlestickColors;
    constructor();
    /**
     * Builds the object.
     */
    build(): candlestick.CandlestickColors;
    up(up: string): this;
    down(down: string): this;
    flat(flat: string): this;
}
