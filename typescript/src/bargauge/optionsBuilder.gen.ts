// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as bargauge from '../bargauge';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<bargauge.Options> {
    protected readonly internal: bargauge.Options;

    constructor() {
        this.internal = bargauge.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): bargauge.Options {
        return this.internal;
    }

    displayMode(displayMode: common.BarGaugeDisplayMode): this {
        this.internal.displayMode = displayMode;
        return this;
    }

    valueMode(valueMode: common.BarGaugeValueMode): this {
        this.internal.valueMode = valueMode;
        return this;
    }

    showUnfilled(showUnfilled: boolean): this {
        this.internal.showUnfilled = showUnfilled;
        return this;
    }

    minVizWidth(minVizWidth: number): this {
        this.internal.minVizWidth = minVizWidth;
        return this;
    }

    reduceOptions(reduceOptions: cog.Builder<common.ReduceDataOptions>): this {
        const reduceOptionsResource = reduceOptions.build();
        this.internal.reduceOptions = reduceOptionsResource;
        return this;
    }

    text(text: cog.Builder<common.VizTextDisplayOptions>): this {
        const textResource = text.build();
        this.internal.text = textResource;
        return this;
    }

    minVizHeight(minVizHeight: number): this {
        this.internal.minVizHeight = minVizHeight;
        return this;
    }

    orientation(orientation: common.VizOrientation): this {
        this.internal.orientation = orientation;
        return this;
    }
}

