// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class HideSeriesConfigBuilder implements com.grafana.foundation.cog.Builder<HideSeriesConfig> {
    protected final HideSeriesConfig internal;
    
    public HideSeriesConfigBuilder() {
        this.internal = new HideSeriesConfig();
    }
    public HideSeriesConfigBuilder tooltip(Boolean tooltip) {
        this.internal.tooltip = tooltip;
        return this;
    }
    
    public HideSeriesConfigBuilder legend(Boolean legend) {
        this.internal.legend = legend;
        return this;
    }
    
    public HideSeriesConfigBuilder viz(Boolean viz) {
        this.internal.viz = viz;
        return this;
    }
    public HideSeriesConfig build() {
        return this.internal;
    }
}
