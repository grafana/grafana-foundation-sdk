// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class VizTooltipOptionsBuilder implements com.grafana.foundation.cog.Builder<VizTooltipOptions> {
    protected final VizTooltipOptions internal;
    
    public VizTooltipOptionsBuilder() {
        this.internal = new VizTooltipOptions();
    }
    public VizTooltipOptionsBuilder mode(TooltipDisplayMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public VizTooltipOptionsBuilder sort(SortOrder sort) {
    this.internal.sort = sort;
        return this;
    }
    
    public VizTooltipOptionsBuilder maxWidth(Double maxWidth) {
    this.internal.maxWidth = maxWidth;
        return this;
    }
    
    public VizTooltipOptionsBuilder maxHeight(Double maxHeight) {
    this.internal.maxHeight = maxHeight;
        return this;
    }
    public VizTooltipOptions build() {
        return this.internal;
    }
}
