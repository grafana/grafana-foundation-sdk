// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;


public class XychartXYSeriesConfigSizeBuilder implements com.grafana.foundation.cog.Builder<XychartXYSeriesConfigSize> {
    protected final XychartXYSeriesConfigSize internal;
    
    public XychartXYSeriesConfigSizeBuilder() {
        this.internal = new XychartXYSeriesConfigSize();
    }
    public XychartXYSeriesConfigSizeBuilder matcher(com.grafana.foundation.cog.Builder<MatcherConfig> matcher) {
        this.internal.matcher = matcher.build();
        return this;
    }
    public XychartXYSeriesConfigSize build() {
        return this.internal;
    }
}
