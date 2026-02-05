// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class ScalarDimensionConfigBuilder implements com.grafana.foundation.cog.Builder<ScalarDimensionConfig> {
    protected final ScalarDimensionConfig internal;
    
    public ScalarDimensionConfigBuilder() {
        this.internal = new ScalarDimensionConfig();
    }
    public ScalarDimensionConfigBuilder min(Double min) {
        this.internal.min = min;
        return this;
    }
    
    public ScalarDimensionConfigBuilder max(Double max) {
        this.internal.max = max;
        return this;
    }
    
    public ScalarDimensionConfigBuilder fixed(Double fixed) {
        this.internal.fixed = fixed;
        return this;
    }
    
    public ScalarDimensionConfigBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public ScalarDimensionConfigBuilder mode(ScalarDimensionMode mode) {
        this.internal.mode = mode;
        return this;
    }
    public ScalarDimensionConfig build() {
        return this.internal;
    }
}
