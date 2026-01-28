// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';
import * as common from '../common';

export class OptionsBuilder implements cog.Builder<heatmap.Options> {
    protected readonly internal: heatmap.Options;

    constructor() {
        this.internal = heatmap.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.Options {
        return this.internal;
    }

    // Controls if the heatmap should be calculated from data
    calculate(calculate: boolean): this {
        this.internal.calculate = calculate;
        return this;
    }

    // Calculation options for the heatmap
    calculation(calculation: cog.Builder<common.HeatmapCalculationOptions>): this {
        const calculationResource = calculation.build();
        this.internal.calculation = calculationResource;
        return this;
    }

    // Controls the color options
    color(color: cog.Builder<heatmap.HeatmapColorOptions>): this {
        const colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }

    // Filters values between a given range
    filterValues(filterValues: cog.Builder<heatmap.FilterValueRange>): this {
        const filterValuesResource = filterValues.build();
        this.internal.filterValues = filterValuesResource;
        return this;
    }

    // Controls tick alignment and value name when not calculating from data
    rowsFrame(rowsFrame: cog.Builder<heatmap.RowsHeatmapOptions>): this {
        const rowsFrameResource = rowsFrame.build();
        this.internal.rowsFrame = rowsFrameResource;
        return this;
    }

    // | *{
    // 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls the display of the value in the cell
    showValue(showValue: common.VisibilityMode): this {
        this.internal.showValue = showValue;
        return this;
    }

    // Controls gap between cells
    cellGap(cellGap: number): this {
        if (!(cellGap <= 25)) {
            throw new Error("cellGap must be <= 25");
        }
        this.internal.cellGap = cellGap;
        return this;
    }

    // Controls cell radius
    cellRadius(cellRadius: number): this {
        this.internal.cellRadius = cellRadius;
        return this;
    }

    // Controls cell value unit
    cellValues(cellValues: cog.Builder<heatmap.CellValues>): this {
        const cellValuesResource = cellValues.build();
        this.internal.cellValues = cellValuesResource;
        return this;
    }

    // Controls yAxis placement
    yAxis(yAxis: cog.Builder<heatmap.YAxisConfig>): this {
        const yAxisResource = yAxis.build();
        this.internal.yAxis = yAxisResource;
        return this;
    }

    // | *{
    // 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls legend options
    legend(legend: cog.Builder<heatmap.HeatmapLegend>): this {
        const legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }

    // Controls tooltip options
    tooltip(tooltip: cog.Builder<heatmap.HeatmapTooltip>): this {
        const tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }

    // Controls exemplar options
    exemplars(exemplars: cog.Builder<heatmap.ExemplarConfig>): this {
        const exemplarsResource = exemplars.build();
        this.internal.exemplars = exemplarsResource;
        return this;
    }
}

