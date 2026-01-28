// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as gauge from '../gauge';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<gauge.Options> {
    protected readonly internal: gauge.Options;

    constructor() {
        this.internal = gauge.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): gauge.Options {
        return this.internal;
    }

    showThresholdLabels(showThresholdLabels: boolean): this {
        this.internal.showThresholdLabels = showThresholdLabels;
        return this;
    }

    showThresholdMarkers(showThresholdMarkers: boolean): this {
        this.internal.showThresholdMarkers = showThresholdMarkers;
        return this;
    }

    sizing(sizing: common.BarGaugeSizing): this {
        this.internal.sizing = sizing;
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

