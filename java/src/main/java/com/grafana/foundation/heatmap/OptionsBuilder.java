// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.grafana.foundation.common.HeatmapCalculationOptions;
import com.grafana.foundation.common.VisibilityMode;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder calculate(Boolean calculate) {
        this.internal.calculate = calculate;
        return this;
    }
    
    public OptionsBuilder calculation(com.grafana.foundation.cog.Builder<HeatmapCalculationOptions> calculation) {
    HeatmapCalculationOptions calculationResource = calculation.build();
        this.internal.calculation = calculationResource;
        return this;
    }
    
    public OptionsBuilder color(com.grafana.foundation.cog.Builder<HeatmapColorOptions> color) {
    HeatmapColorOptions colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }
    
    public OptionsBuilder filterValues(com.grafana.foundation.cog.Builder<FilterValueRange> filterValues) {
    FilterValueRange filterValuesResource = filterValues.build();
        this.internal.filterValues = filterValuesResource;
        return this;
    }
    
    public OptionsBuilder rowsFrame(com.grafana.foundation.cog.Builder<RowsHeatmapOptions> rowsFrame) {
    RowsHeatmapOptions rowsFrameResource = rowsFrame.build();
        this.internal.rowsFrame = rowsFrameResource;
        return this;
    }
    
    public OptionsBuilder showValue(VisibilityMode showValue) {
        this.internal.showValue = showValue;
        return this;
    }
    
    public OptionsBuilder cellGap(Integer cellGap) {
        if (!(cellGap <= 25)) {
            throw new IllegalArgumentException("cellGap must be <= 25");
        }
        this.internal.cellGap = cellGap;
        return this;
    }
    
    public OptionsBuilder cellRadius(Float cellRadius) {
        this.internal.cellRadius = cellRadius;
        return this;
    }
    
    public OptionsBuilder cellValues(com.grafana.foundation.cog.Builder<CellValues> cellValues) {
    CellValues cellValuesResource = cellValues.build();
        this.internal.cellValues = cellValuesResource;
        return this;
    }
    
    public OptionsBuilder yAxis(com.grafana.foundation.cog.Builder<YAxisConfig> yAxis) {
    YAxisConfig yAxisResource = yAxis.build();
        this.internal.yAxis = yAxisResource;
        return this;
    }
    
    public OptionsBuilder legend(com.grafana.foundation.cog.Builder<HeatmapLegend> legend) {
    HeatmapLegend legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }
    
    public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<HeatmapTooltip> tooltip) {
    HeatmapTooltip tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }
    
    public OptionsBuilder exemplars(com.grafana.foundation.cog.Builder<ExemplarConfig> exemplars) {
    ExemplarConfig exemplarsResource = exemplars.build();
        this.internal.exemplars = exemplarsResource;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
