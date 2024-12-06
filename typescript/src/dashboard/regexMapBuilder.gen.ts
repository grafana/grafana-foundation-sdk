// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Maps regular expressions to replacement text and a color.
// For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
export class RegexMapBuilder implements cog.Builder<dashboard.RegexMap> {
    protected readonly internal: dashboard.RegexMap;

    constructor() {
        this.internal = dashboard.defaultRegexMap();
        this.internal.type = "regex";
    }

    /**
     * Builds the object.
     */
    build(): dashboard.RegexMap {
        return this.internal;
    }

    // Regular expression to match against and the result to apply when the value matches the regex
    options(options: {
	// Regular expression to match against
	pattern: string;
	// Config to apply when the value matches the regex
	result: dashboard.ValueMappingResult;
}): this {
        this.internal.options = options;
        return this;
    }
}
