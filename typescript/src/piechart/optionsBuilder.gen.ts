// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as piechart from '../piechart';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<piechart.Options> {
    protected readonly internal: piechart.Options;

    constructor() {
        this.internal = piechart.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): piechart.Options {
        return this.internal;
    }

    pieType(pieType: piechart.PieChartType): this {
        this.internal.pieType = pieType;
        return this;
    }

    displayLabels(displayLabels: piechart.PieChartLabels[]): this {
        this.internal.displayLabels = displayLabels;
        return this;
    }

    tooltip(tooltip: cog.Builder<common.VizTooltipOptions>): this {
        const tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
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

    legend(legend: cog.Builder<piechart.PieChartLegendOptions>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }

    orientation(orientation: common.VizOrientation): this {
        this.internal.orientation = orientation;
        return this;
    }
}

