// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;


public class ControlsOptionsBuilder implements com.grafana.foundation.cog.Builder<ControlsOptions> {
    protected final ControlsOptions internal;
    
    public ControlsOptionsBuilder() {
        this.internal = new ControlsOptions();
    }
    public ControlsOptionsBuilder showZoom(Boolean showZoom) {
        this.internal.showZoom = showZoom;
        return this;
    }
    
    public ControlsOptionsBuilder mouseWheelZoom(Boolean mouseWheelZoom) {
        this.internal.mouseWheelZoom = mouseWheelZoom;
        return this;
    }
    
    public ControlsOptionsBuilder showAttribution(Boolean showAttribution) {
        this.internal.showAttribution = showAttribution;
        return this;
    }
    
    public ControlsOptionsBuilder showScale(Boolean showScale) {
        this.internal.showScale = showScale;
        return this;
    }
    
    public ControlsOptionsBuilder showDebug(Boolean showDebug) {
        this.internal.showDebug = showDebug;
        return this;
    }
    
    public ControlsOptionsBuilder showMeasure(Boolean showMeasure) {
        this.internal.showMeasure = showMeasure;
        return this;
    }
    public ControlsOptions build() {
        return this.internal;
    }
}
