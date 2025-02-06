// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import java.util.List;

public class VizLegendOptionsBuilder implements com.grafana.foundation.cog.Builder<VizLegendOptions> {
    protected final VizLegendOptions internal;
    
    public VizLegendOptionsBuilder() {
        this.internal = new VizLegendOptions();
    }
    public VizLegendOptionsBuilder displayMode(LegendDisplayMode displayMode) {
        this.internal.displayMode = displayMode;
        return this;
    }
    
    public VizLegendOptionsBuilder placement(LegendPlacement placement) {
        this.internal.placement = placement;
        return this;
    }
    
    public VizLegendOptionsBuilder showLegend(Boolean showLegend) {
        this.internal.showLegend = showLegend;
        return this;
    }
    
    public VizLegendOptionsBuilder asTable(Boolean asTable) {
        this.internal.asTable = asTable;
        return this;
    }
    
    public VizLegendOptionsBuilder isVisible(Boolean isVisible) {
        this.internal.isVisible = isVisible;
        return this;
    }
    
    public VizLegendOptionsBuilder sortBy(String sortBy) {
        this.internal.sortBy = sortBy;
        return this;
    }
    
    public VizLegendOptionsBuilder sortDesc(Boolean sortDesc) {
        this.internal.sortDesc = sortDesc;
        return this;
    }
    
    public VizLegendOptionsBuilder width(Double width) {
        this.internal.width = width;
        return this;
    }
    
    public VizLegendOptionsBuilder calcs(List<String> calcs) {
        this.internal.calcs = calcs;
        return this;
    }
    public VizLegendOptions build() {
        return this.internal;
    }
}
