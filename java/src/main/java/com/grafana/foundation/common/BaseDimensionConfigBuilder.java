// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class BaseDimensionConfigBuilder implements com.grafana.foundation.cog.Builder<BaseDimensionConfig> {
    protected final BaseDimensionConfig internal;
    
    public BaseDimensionConfigBuilder() {
        this.internal = new BaseDimensionConfig();
    }
    public BaseDimensionConfigBuilder field(String field) {
    this.internal.field = field;
        return this;
    }
    public BaseDimensionConfig build() {
        return this.internal;
    }
}
