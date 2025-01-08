// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import java.util.List;

public class XYDimensionConfigBuilder implements com.grafana.foundation.cog.Builder<XYDimensionConfig> {
    protected final XYDimensionConfig internal;
    
    public XYDimensionConfigBuilder() {
        this.internal = new XYDimensionConfig();
    }
    public XYDimensionConfigBuilder frame(Integer frame) {
        if (!(frame >= 0)) {
            throw new IllegalArgumentException("frame must be >= 0");
        }
    this.internal.frame = frame;
        return this;
    }
    
    public XYDimensionConfigBuilder x(String x) {
    this.internal.x = x;
        return this;
    }
    
    public XYDimensionConfigBuilder exclude(List<String> exclude) {
    this.internal.exclude = exclude;
        return this;
    }
    public XYDimensionConfig build() {
        return this.internal;
    }
}
