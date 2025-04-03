// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as xychart from '../xychart';

// NOTE: (copied from dashboard_kind.cue, since not exported)
// Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
// It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
export class MatcherConfigBuilder implements cog.Builder<xychart.MatcherConfig> {
    protected readonly internal: xychart.MatcherConfig;

    constructor() {
        this.internal = xychart.defaultMatcherConfig();
    }

    /**
     * Builds the object.
     */
    build(): xychart.MatcherConfig {
        return this.internal;
    }

    // The matcher id. This is used to find the matcher implementation from registry.
    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    // The matcher options. This is specific to the matcher implementation.
    options(options: any): this {
        this.internal.options = options;
        return this;
    }
}

