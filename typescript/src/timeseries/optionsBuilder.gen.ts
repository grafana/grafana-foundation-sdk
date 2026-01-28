// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as timeseries from '../timeseries';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<timeseries.Options> {
    protected readonly internal: timeseries.Options;

    constructor() {
        this.internal = timeseries.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): timeseries.Options {
        return this.internal;
    }

    timezone(timezone: common.TimeZone[]): this {
        this.internal.timezone = timezone;
        return this;
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

    orientation(orientation: common.VizOrientation): this {
        this.internal.orientation = orientation;
        return this;
    }
}

