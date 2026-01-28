// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import java.util.List;
import java.util.LinkedList;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder mapping(SeriesMapping mapping) {
        this.internal.mapping = mapping;
        return this;
    }
    
    public OptionsBuilder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
    VizLegendOptions legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }
    
    public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
    VizTooltipOptions tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }
    
    public OptionsBuilder series(List<com.grafana.foundation.cog.Builder<XYSeriesConfig>> series) {
        List<XYSeriesConfig> seriesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<XYSeriesConfig> r1 : series) {
                XYSeriesConfig seriesDepth1 = r1.build();
                seriesResources.add(seriesDepth1); 
        }
        this.internal.series = seriesResources;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
