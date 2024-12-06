// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';
import * as common from '../common';

// Controls frame rows options
export class RowsHeatmapOptionsBuilder implements cog.Builder<heatmap.RowsHeatmapOptions> {
    protected readonly internal: heatmap.RowsHeatmapOptions;

    constructor() {
        this.internal = heatmap.defaultRowsHeatmapOptions();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.RowsHeatmapOptions {
        return this.internal;
    }

    // Sets the name of the cell when not calculating from data
    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    // Controls tick alignment when not calculating from data
    layout(layout: common.HeatmapCellLayout): this {
        this.internal.layout = layout;
        return this;
    }
}
