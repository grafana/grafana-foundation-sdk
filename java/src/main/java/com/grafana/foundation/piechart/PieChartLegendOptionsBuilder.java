// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.piechart;

import java.util.List;
import com.grafana.foundation.common.LegendDisplayMode;
import com.grafana.foundation.common.LegendPlacement;

public class PieChartLegendOptionsBuilder implements com.grafana.foundation.cog.Builder<PieChartLegendOptions> {
    protected final PieChartLegendOptions internal;
    
    public PieChartLegendOptionsBuilder() {
        this.internal = new PieChartLegendOptions();
    }
    public PieChartLegendOptionsBuilder values(List<PieChartLegendValues> values) {
        this.internal.values = values;
        return this;
    }
    
    public PieChartLegendOptionsBuilder displayMode(LegendDisplayMode displayMode) {
        this.internal.displayMode = displayMode;
        return this;
    }
    
    public PieChartLegendOptionsBuilder placement(LegendPlacement placement) {
        this.internal.placement = placement;
        return this;
    }
    
    public PieChartLegendOptionsBuilder showLegend(Boolean showLegend) {
        this.internal.showLegend = showLegend;
        return this;
    }
    
    public PieChartLegendOptionsBuilder asTable(Boolean asTable) {
        this.internal.asTable = asTable;
        return this;
    }
    
    public PieChartLegendOptionsBuilder isVisible(Boolean isVisible) {
        this.internal.isVisible = isVisible;
        return this;
    }
    
    public PieChartLegendOptionsBuilder sortBy(String sortBy) {
        this.internal.sortBy = sortBy;
        return this;
    }
    
    public PieChartLegendOptionsBuilder sortDesc(Boolean sortDesc) {
        this.internal.sortDesc = sortDesc;
        return this;
    }
    
    public PieChartLegendOptionsBuilder width(Double width) {
        this.internal.width = width;
        return this;
    }
    
    public PieChartLegendOptionsBuilder calcs(List<String> calcs) {
        this.internal.calcs = calcs;
        return this;
    }
    public PieChartLegendOptions build() {
        return this.internal;
    }
}
