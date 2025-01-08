// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class ScaleDimensionConfigBuilder implements com.grafana.foundation.cog.Builder<ScaleDimensionConfig> {
    protected final ScaleDimensionConfig internal;
    
    public ScaleDimensionConfigBuilder() {
        this.internal = new ScaleDimensionConfig();
    }
    public ScaleDimensionConfigBuilder min(Double min) {
    this.internal.min = min;
        return this;
    }
    
    public ScaleDimensionConfigBuilder max(Double max) {
    this.internal.max = max;
        return this;
    }
    
    public ScaleDimensionConfigBuilder fixed(Double fixed) {
    this.internal.fixed = fixed;
        return this;
    }
    
    public ScaleDimensionConfigBuilder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public ScaleDimensionConfigBuilder mode(ScaleDimensionMode mode) {
    this.internal.mode = mode;
        return this;
    }
    public ScaleDimensionConfig build() {
        return this.internal;
    }
}
