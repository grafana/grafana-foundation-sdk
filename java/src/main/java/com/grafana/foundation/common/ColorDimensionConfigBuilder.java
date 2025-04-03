// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class ColorDimensionConfigBuilder implements com.grafana.foundation.cog.Builder<ColorDimensionConfig> {
    protected final ColorDimensionConfig internal;
    
    public ColorDimensionConfigBuilder() {
        this.internal = new ColorDimensionConfig();
    }
    public ColorDimensionConfigBuilder fixed(String fixed) {
        this.internal.fixed = fixed;
        return this;
    }
    
    public ColorDimensionConfigBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    public ColorDimensionConfig build() {
        return this.internal;
    }
}
