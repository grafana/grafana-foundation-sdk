// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as trend from '../trend';
import * as common from '../common';

// Identical to timeseries... except it does not have timezone settings
export class OptionsBuilder implements cog.Builder<trend.Options> {
    protected readonly internal: trend.Options;

    constructor() {
        this.internal = trend.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): trend.Options {
        return this.internal;
    }

    legend(legend: cog.Builder<common.VizLegendOptions>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }

    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this {
        const tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }

    // Name of the x field to use (defaults to first number)
    xField(xField: string): this {
        this.internal.xField = xField;
        return this;
    }
}

