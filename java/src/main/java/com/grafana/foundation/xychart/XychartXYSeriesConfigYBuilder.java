// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


public class XychartXYSeriesConfigYBuilder implements com.grafana.foundation.cog.Builder<XychartXYSeriesConfigY> {
    protected final XychartXYSeriesConfigY internal;
    
    public XychartXYSeriesConfigYBuilder() {
        this.internal = new XychartXYSeriesConfigY();
    }
    public XychartXYSeriesConfigYBuilder matcher(com.grafana.foundation.cog.Builder<MatcherConfig> matcher) {
    MatcherConfig matcherResource = matcher.build();
        this.internal.matcher = matcherResource;
        return this;
    }
    public XychartXYSeriesConfigY build() {
        return this.internal;
    }
}
