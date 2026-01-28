// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as candlestick from '../candlestick';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<candlestick.Options> {
    protected readonly internal: candlestick.Options;

    constructor() {
        this.internal = candlestick.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): candlestick.Options {
        return this.internal;
    }

    // Sets which dimensions are used for the visualization
    mode(mode: candlestick.VizDisplayMode): this {
        this.internal.mode = mode;
        return this;
    }

    // Sets the style of the candlesticks
    candleStyle(candleStyle: candlestick.CandleStyle): this {
        this.internal.candleStyle = candleStyle;
        return this;
    }

    // Sets the color strategy for the candlesticks
    colorStrategy(colorStrategy: candlestick.ColorStrategy): this {
        this.internal.colorStrategy = colorStrategy;
        return this;
    }

    // Map fields to appropriate dimension
    fields(fields: cog.Builder<candlestick.CandlestickFieldMap>): this {
        const fieldsResource = fields.build();
        this.internal.fields = fieldsResource;
        return this;
    }

    // Set which colors are used when the price movement is up or down
    colors(colors: cog.Builder<candlestick.CandlestickColors>): this {
        const colorsResource = colors.build();
        this.internal.colors = colorsResource;
        return this;
    }

    legend(legend: cog.Builder<common.VizLegendOptions>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }

    // When enabled, all fields will be sent to the graph
    includeAllFields(includeAllFields: boolean): this {
        this.internal.includeAllFields = includeAllFields;
        return this;
    }
}

