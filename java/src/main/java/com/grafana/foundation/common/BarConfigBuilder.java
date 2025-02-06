// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class BarConfigBuilder implements com.grafana.foundation.cog.Builder<BarConfig> {
    protected final BarConfig internal;
    
    public BarConfigBuilder() {
        this.internal = new BarConfig();
    }
    public BarConfigBuilder barAlignment(BarAlignment barAlignment) {
        this.internal.barAlignment = barAlignment;
        return this;
    }
    
    public BarConfigBuilder barWidthFactor(Double barWidthFactor) {
        this.internal.barWidthFactor = barWidthFactor;
        return this;
    }
    
    public BarConfigBuilder barMaxWidth(Double barMaxWidth) {
        this.internal.barMaxWidth = barMaxWidth;
        return this;
    }
    public BarConfig build() {
        return this.internal;
    }
}
