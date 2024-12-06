// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as heatmap from '../heatmap';

// Controls the value filter range
export class FilterValueRangeBuilder implements cog.Builder<heatmap.FilterValueRange> {
    protected readonly internal: heatmap.FilterValueRange;

    constructor() {
        this.internal = heatmap.defaultFilterValueRange();
    }

    /**
     * Builds the object.
     */
    build(): heatmap.FilterValueRange {
        return this.internal;
    }

    // Sets the filter range to values less than or equal to the given value
    le(le: number): this {
        this.internal.le = le;
        return this;
    }

    // Sets the filter range to values greater than or equal to the given value
    ge(ge: number): this {
        this.internal.ge = ge;
        return this;
    }
}
