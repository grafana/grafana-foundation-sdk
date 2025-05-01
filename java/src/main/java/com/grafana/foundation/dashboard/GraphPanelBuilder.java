// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class GraphPanelBuilder implements com.grafana.foundation.cog.Builder<GraphPanel> {
    protected final GraphPanel internal;
    
    public GraphPanelBuilder() {
        this.internal = new GraphPanel();
        this.internal.type = "graph";
    }
    public GraphPanelBuilder legend(com.grafana.foundation.cog.Builder<DashboardGraphPanelLegend> legend) {
    DashboardGraphPanelLegend legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }
    public GraphPanel build() {
        return this.internal;
    }
}
