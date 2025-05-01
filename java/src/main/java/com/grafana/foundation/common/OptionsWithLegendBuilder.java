// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class OptionsWithLegendBuilder implements com.grafana.foundation.cog.Builder<OptionsWithLegend> {
    protected final OptionsWithLegend internal;
    
    public OptionsWithLegendBuilder() {
        this.internal = new OptionsWithLegend();
    }
    public OptionsWithLegendBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
    VizLegendOptions legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }
    public OptionsWithLegend build() {
        return this.internal;
    }
}
