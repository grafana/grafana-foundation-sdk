// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class ResourceDimensionConfigBuilder implements com.grafana.foundation.cog.Builder<ResourceDimensionConfig> {
    protected final ResourceDimensionConfig internal;
    
    public ResourceDimensionConfigBuilder() {
        this.internal = new ResourceDimensionConfig();
    }
    public ResourceDimensionConfigBuilder mode(ResourceDimensionMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public ResourceDimensionConfigBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public ResourceDimensionConfigBuilder fixed(String fixed) {
        this.internal.fixed = fixed;
        return this;
    }
    public ResourceDimensionConfig build() {
        return this.internal;
    }
}
