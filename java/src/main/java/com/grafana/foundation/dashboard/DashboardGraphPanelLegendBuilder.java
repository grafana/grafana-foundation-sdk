// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class DashboardGraphPanelLegendBuilder implements com.grafana.foundation.cog.Builder<DashboardGraphPanelLegend> {
    protected final DashboardGraphPanelLegend internal;
    
    public DashboardGraphPanelLegendBuilder() {
        this.internal = new DashboardGraphPanelLegend();
    }
    public DashboardGraphPanelLegendBuilder show(Boolean show) {
        this.internal.show = show;
        return this;
    }
    
    public DashboardGraphPanelLegendBuilder sort(String sort) {
        this.internal.sort = sort;
        return this;
    }
    
    public DashboardGraphPanelLegendBuilder sortDesc(Boolean sortDesc) {
        this.internal.sortDesc = sortDesc;
        return this;
    }
    public DashboardGraphPanelLegend build() {
        return this.internal;
    }
}
