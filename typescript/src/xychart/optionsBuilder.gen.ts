// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as xychart from '../xychart';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<xychart.Options> {
    protected readonly internal: xychart.Options;

    constructor() {
        this.internal = xychart.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): xychart.Options {
        return this.internal;
    }

    seriesMapping(seriesMapping: xychart.SeriesMapping): this {
        this.internal.seriesMapping = seriesMapping;
        return this;
    }

    dims(dims: cog.Builder<xychart.XYDimensionConfig>): this {
        const dimsResource = dims.build();
        this.internal.dims = dimsResource;
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

    series(series: cog.Builder<xychart.ScatterSeriesConfig>[]): this {
        const seriesResources = series.map(builder1 => builder1.build());
        this.internal.series = seriesResources;
        return this;
    }
}

