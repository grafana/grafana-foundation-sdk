// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as candlestick from '../candlestick';

export class CandlestickColorsBuilder implements cog.Builder<candlestick.CandlestickColors> {
    protected readonly internal: candlestick.CandlestickColors;

    constructor() {
        this.internal = candlestick.defaultCandlestickColors();
    }

    /**
     * Builds the object.
     */
    build(): candlestick.CandlestickColors {
        return this.internal;
    }

    up(up: string): this {
        this.internal.up = up;
        return this;
    }

    down(down: string): this {
        this.internal.down = down;
        return this;
    }

    flat(flat: string): this {
        this.internal.flat = flat;
        return this;
    }
}
