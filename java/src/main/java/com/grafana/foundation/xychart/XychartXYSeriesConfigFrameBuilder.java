// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


public class XychartXYSeriesConfigFrameBuilder implements com.grafana.foundation.cog.Builder<XychartXYSeriesConfigFrame> {
    protected final XychartXYSeriesConfigFrame internal;
    
    public XychartXYSeriesConfigFrameBuilder() {
        this.internal = new XychartXYSeriesConfigFrame();
    }
    public XychartXYSeriesConfigFrameBuilder matcher(com.grafana.foundation.cog.Builder<MatcherConfig> matcher) {
    MatcherConfig matcherResource = matcher.build();
        this.internal.matcher = matcherResource;
        return this;
    }
    public XychartXYSeriesConfigFrame build() {
        return this.internal;
    }
}
