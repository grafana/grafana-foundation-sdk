// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class VizTextDisplayOptionsBuilder implements com.grafana.foundation.cog.Builder<VizTextDisplayOptions> {
    protected final VizTextDisplayOptions internal;
    
    public VizTextDisplayOptionsBuilder() {
        this.internal = new VizTextDisplayOptions();
    }
    public VizTextDisplayOptionsBuilder titleSize(Double titleSize) {
    this.internal.titleSize = titleSize;
        return this;
    }
    
    public VizTextDisplayOptionsBuilder valueSize(Double valueSize) {
    this.internal.valueSize = valueSize;
        return this;
    }
    public VizTextDisplayOptions build() {
        return this.internal;
    }
}
