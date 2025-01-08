// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


public class XychartXYSeriesConfigColorBuilder implements com.grafana.foundation.cog.Builder<XychartXYSeriesConfigColor> {
    protected final XychartXYSeriesConfigColor internal;
    
    public XychartXYSeriesConfigColorBuilder() {
        this.internal = new XychartXYSeriesConfigColor();
    }
    public XychartXYSeriesConfigColorBuilder matcher(com.grafana.foundation.cog.Builder<MatcherConfig> matcher) {
    this.internal.matcher = matcher.build();
        return this;
    }
    public XychartXYSeriesConfigColor build() {
        return this.internal;
    }
}
