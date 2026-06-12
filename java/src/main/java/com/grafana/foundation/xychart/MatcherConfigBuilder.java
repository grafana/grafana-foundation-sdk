// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


/**
 * NOTE: (copied from dashboard_kind.cue, since not exported)
 * Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
 * It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
 */
public class MatcherConfigBuilder implements com.grafana.foundation.cog.Builder<MatcherConfig> {
    protected final MatcherConfig internal;
    
    public MatcherConfigBuilder() {
        this.internal = new MatcherConfig();
    }
    public MatcherConfigBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MatcherConfigBuilder options(Object options) {
        this.internal.options = options;
        return this;
    }
    public MatcherConfig build() {
        return this.internal;
    }
}
