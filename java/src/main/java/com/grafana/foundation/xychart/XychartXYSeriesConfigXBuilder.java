// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


public class XychartXYSeriesConfigXBuilder implements com.grafana.foundation.cog.Builder<XychartXYSeriesConfigX> {
    protected final XychartXYSeriesConfigX internal;
    
    public XychartXYSeriesConfigXBuilder() {
        this.internal = new XychartXYSeriesConfigX();
    }
    public XychartXYSeriesConfigXBuilder matcher(com.grafana.foundation.cog.Builder<MatcherConfig> matcher) {
    this.internal.matcher = matcher.build();
        return this;
    }
    public XychartXYSeriesConfigX build() {
        return this.internal;
    }
}
