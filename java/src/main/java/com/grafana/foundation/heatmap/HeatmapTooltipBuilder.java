// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.grafana.foundation.common.TooltipDisplayMode;

public class HeatmapTooltipBuilder implements com.grafana.foundation.cog.Builder<HeatmapTooltip> {
    protected final HeatmapTooltip internal;
    
    public HeatmapTooltipBuilder() {
        this.internal = new HeatmapTooltip();
    }
    public HeatmapTooltipBuilder mode(TooltipDisplayMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public HeatmapTooltipBuilder maxHeight(Double maxHeight) {
        this.internal.maxHeight = maxHeight;
        return this;
    }
    
    public HeatmapTooltipBuilder maxWidth(Double maxWidth) {
        this.internal.maxWidth = maxWidth;
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
