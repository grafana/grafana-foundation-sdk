// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class FillConfigBuilder implements com.grafana.foundation.cog.Builder<FillConfig> {
    protected final FillConfig internal;
    
    public FillConfigBuilder() {
        this.internal = new FillConfig();
    }
    public FillConfigBuilder fillColor(String fillColor) {
        this.internal.fillColor = fillColor;
        return this;
    }
    
    public FillConfigBuilder fillOpacity(Double fillOpacity) {
        this.internal.fillOpacity = fillOpacity;
        return this;
    }
    
    public FillConfigBuilder fillBelowTo(String fillBelowTo) {
        this.internal.fillBelowTo = fillBelowTo;
        return this;
    }
    public FillConfig build() {
        return this.internal;
    }
}
