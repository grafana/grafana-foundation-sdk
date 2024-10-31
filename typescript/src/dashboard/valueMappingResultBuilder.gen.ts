// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Result used as replacement with text and color when the value matches
export class ValueMappingResultBuilder implements cog.Builder<dashboard.ValueMappingResult> {
    protected readonly internal: dashboard.ValueMappingResult;

    constructor() {
        this.internal = dashboard.defaultValueMappingResult();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.ValueMappingResult {
        return this.internal;
    }

    // Text to display when the value matches
    text(text: string): this {
        this.internal.text = text;
        return this;
    }

    // Text to use when the value matches
    color(color: string): this {
        this.internal.color = color;
        return this;
    }

    // Icon to display when the value matches. Only specific visualizations.
    icon(icon: string): this {
        this.internal.icon = icon;
        return this;
    }

    // Position in the mapping array. Only used internally.
    index(index: number): this {
        this.internal.index = index;
        return this;
    }
}
