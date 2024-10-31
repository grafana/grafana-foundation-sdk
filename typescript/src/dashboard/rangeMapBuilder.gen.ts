// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Maps numerical ranges to a display text and color.
// For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
export class RangeMapBuilder implements cog.Builder<dashboard.RangeMap> {
    protected readonly internal: dashboard.RangeMap;

    constructor() {
        this.internal = dashboard.defaultRangeMap();
        this.internal.type = "range";
    }

    /**
     * Builds the object.
     */
    build(): dashboard.RangeMap {
        return this.internal;
    }

    // Range to match against and the result to apply when the value is within the range
    options(options: {
	// Min value of the range. It can be null which means -Infinity
	from: number | null;
	// Max value of the range. It can be null which means +Infinity
	to: number | null;
	// Config to apply when the value is within the range
	result: dashboard.ValueMappingResult;
}): this {
        this.internal.options = options;
        return this;
    }
}
