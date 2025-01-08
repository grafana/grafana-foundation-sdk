// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class TextDimensionConfigBuilder implements com.grafana.foundation.cog.Builder<TextDimensionConfig> {
    protected final TextDimensionConfig internal;
    
    public TextDimensionConfigBuilder() {
        this.internal = new TextDimensionConfig();
    }
    public TextDimensionConfigBuilder mode(TextDimensionMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public TextDimensionConfigBuilder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public TextDimensionConfigBuilder fixed(String fixed) {
    this.internal.fixed = fixed;
        return this;
    }
    public TextDimensionConfig build() {
        return this.internal;
    }
}
