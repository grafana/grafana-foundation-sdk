// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;


public class HeatmapTooltipBuilder implements com.grafana.foundation.cog.Builder<HeatmapTooltip> {
    protected final HeatmapTooltip internal;
    
    public HeatmapTooltipBuilder() {
        this.internal = new HeatmapTooltip();
    }
    public HeatmapTooltipBuilder show(Boolean show) {
    this.internal.show = show;
        return this;
    }
    
    public HeatmapTooltipBuilder yHistogram(Boolean yHistogram) {
    this.internal.yHistogram = yHistogram;
        return this;
    }
    
    public HeatmapTooltipBuilder showColorScale(Boolean showColorScale) {
    this.internal.showColorScale = showColorScale;
        return this;
    }
    public HeatmapTooltip build() {
        return this.internal;
    }
}
