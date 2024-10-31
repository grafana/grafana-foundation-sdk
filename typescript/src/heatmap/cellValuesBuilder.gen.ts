// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';

// Controls cell value options
export class CellValuesBuilder implements cog.Builder<heatmap.CellValues> {
    protected readonly internal: heatmap.CellValues;

    constructor() {
        this.internal = heatmap.defaultCellValues();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.CellValues {
        return this.internal;
    }

    // Controls the cell value unit
    unit(unit: string): this {
        this.internal.unit = unit;
        return this;
    }

    // Controls the number of decimals for cell values
    decimals(decimals: number): this {
        this.internal.decimals = decimals;
        return this;
    }
}
