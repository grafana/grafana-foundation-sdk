// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;


public class HeatmapLegendBuilder implements com.grafana.foundation.cog.Builder<HeatmapLegend> {
    protected final HeatmapLegend internal;
    
    public HeatmapLegendBuilder() {
        this.internal = new HeatmapLegend();
    }
    public HeatmapLegendBuilder show(Boolean show) {
    this.internal.show = show;
        return this;
    }
    public HeatmapLegend build() {
        return this.internal;
    }
}
