// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


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
