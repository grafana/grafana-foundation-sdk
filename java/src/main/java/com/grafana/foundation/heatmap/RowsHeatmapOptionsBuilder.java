// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.grafana.foundation.common.HeatmapCellLayout;

public class RowsHeatmapOptionsBuilder implements com.grafana.foundation.cog.Builder<RowsHeatmapOptions> {
    protected final RowsHeatmapOptions internal;
    
    public RowsHeatmapOptionsBuilder() {
        this.internal = new RowsHeatmapOptions();
    }
    public RowsHeatmapOptionsBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public RowsHeatmapOptionsBuilder layout(HeatmapCellLayout layout) {
        this.internal.layout = layout;
        return this;
    }
    public RowsHeatmapOptions build() {
        return this.internal;
    }
}
