// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class StackingConfigBuilder implements com.grafana.foundation.cog.Builder<StackingConfig> {
    protected final StackingConfig internal;
    
    public StackingConfigBuilder() {
        this.internal = new StackingConfig();
    }
    public StackingConfigBuilder mode(StackingMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public StackingConfigBuilder group(String group) {
    this.internal.group = group;
        return this;
    }
    public StackingConfig build() {
        return this.internal;
    }
}
