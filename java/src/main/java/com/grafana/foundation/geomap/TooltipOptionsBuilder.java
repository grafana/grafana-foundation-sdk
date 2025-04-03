// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;


public class TooltipOptionsBuilder implements com.grafana.foundation.cog.Builder<TooltipOptions> {
    protected final TooltipOptions internal;
    
    public TooltipOptionsBuilder() {
        this.internal = new TooltipOptions();
    }
    public TooltipOptionsBuilder mode(TooltipMode mode) {
        this.internal.mode = mode;
        return this;
    }
    public TooltipOptions build() {
        return this.internal;
    }
}
