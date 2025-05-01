// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


public class XYSeriesConfigBuilder implements com.grafana.foundation.cog.Builder<XYSeriesConfig> {
    protected final XYSeriesConfig internal;
    
    public XYSeriesConfigBuilder() {
        this.internal = new XYSeriesConfig();
    }
    public XYSeriesConfigBuilder name(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigName> name) {
    XychartXYSeriesConfigName nameResource = name.build();
        this.internal.name = nameResource;
        return this;
    }
    
    public XYSeriesConfigBuilder frame(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigFrame> frame) {
    XychartXYSeriesConfigFrame frameResource = frame.build();
        this.internal.frame = frameResource;
        return this;
    }
    
    public XYSeriesConfigBuilder x(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigX> x) {
    XychartXYSeriesConfigX xResource = x.build();
        this.internal.x = xResource;
        return this;
    }
    
    public XYSeriesConfigBuilder y(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigY> y) {
    XychartXYSeriesConfigY yResource = y.build();
        this.internal.y = yResource;
        return this;
    }
    
    public XYSeriesConfigBuilder color(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigColor> color) {
    XychartXYSeriesConfigColor colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }
    
    public XYSeriesConfigBuilder size(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigSize> size) {
    XychartXYSeriesConfigSize sizeResource = size.build();
        this.internal.size = sizeResource;
        return this;
    }
    public XYSeriesConfig build() {
        return this.internal;
    }
}
