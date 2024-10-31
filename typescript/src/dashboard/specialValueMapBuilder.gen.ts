// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
// See SpecialValueMatch to see the list of special values.
// For example, you can configure a special value mapping so that null values appear as N/A.
export class SpecialValueMapBuilder implements cog.Builder<dashboard.SpecialValueMap> {
    protected readonly internal: dashboard.SpecialValueMap;

    constructor() {
        this.internal = dashboard.defaultSpecialValueMap();
        this.internal.type = "special";
    }

    /**
     * Builds the object.
     */
    build(): dashboard.SpecialValueMap {
        return this.internal;
    }

    options(options: {
	// Special value to match against
	match: dashboard.SpecialValueMatch;
	// Config to apply when the value matches the special value
	result: dashboard.ValueMappingResult;
}): this {
        this.internal.options = options;
        return this;
    }
}
