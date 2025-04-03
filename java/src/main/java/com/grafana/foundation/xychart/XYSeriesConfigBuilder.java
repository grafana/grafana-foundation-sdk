// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


public class XYSeriesConfigBuilder implements com.grafana.foundation.cog.Builder<XYSeriesConfig> {
    protected final XYSeriesConfig internal;
    
    public XYSeriesConfigBuilder() {
        this.internal = new XYSeriesConfig();
    }
    public XYSeriesConfigBuilder name(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigName> name) {
        this.internal.name = name.build();
        return this;
    }
    
    public XYSeriesConfigBuilder frame(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigFrame> frame) {
        this.internal.frame = frame.build();
        return this;
    }
    
    public XYSeriesConfigBuilder x(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigX> x) {
        this.internal.x = x.build();
        return this;
    }
    
    public XYSeriesConfigBuilder y(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigY> y) {
        this.internal.y = y.build();
        return this;
    }
    
    public XYSeriesConfigBuilder color(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigColor> color) {
        this.internal.color = color.build();
        return this;
    }
    
    public XYSeriesConfigBuilder size(com.grafana.foundation.cog.Builder<XychartXYSeriesConfigSize> size) {
        this.internal.size = size.build();
        return this;
    }
    public XYSeriesConfig build() {
        return this.internal;
    }
}
