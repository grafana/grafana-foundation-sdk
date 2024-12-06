// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Maps text values to a color or different display text and color.
// For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
export class ValueMapBuilder implements cog.Builder<dashboard.ValueMap> {
    protected readonly internal: dashboard.ValueMap;

    constructor() {
        this.internal = dashboard.defaultValueMap();
        this.internal.type = "value";
    }

    /**
     * Builds the object.
     */
    build(): dashboard.ValueMap {
        return this.internal;
    }

    // Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
    options(options: Record<string, dashboard.ValueMappingResult>): this {
        this.internal.options = options;
        return this;
    }
}
