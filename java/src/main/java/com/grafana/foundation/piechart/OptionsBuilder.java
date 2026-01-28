// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.piechart;

import java.util.List;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder pieType(PieChartType pieType) {
        this.internal.pieType = pieType;
        return this;
    }
    
    public OptionsBuilder displayLabels(List<PieChartLabels> displayLabels) {
        this.internal.displayLabels = displayLabels;
        return this;
    }
    
    public OptionsBuilder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
    VizTooltipOptions tooltipResource = tooltip.build();
        this.internal.tooltip = tooltipResource;
        return this;
    }
    
    public OptionsBuilder reduceOptions(com.grafana.foundation.cog.Builder<ReduceDataOptions> reduceOptions) {
    ReduceDataOptions reduceOptionsResource = reduceOptions.build();
        this.internal.reduceOptions = reduceOptionsResource;
        return this;
    }
    
    public OptionsBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
    VizTextDisplayOptions textResource = text.build();
        this.internal.text = textResource;
        return this;
    }
    
    public OptionsBuilder legend(com.grafana.foundation.cog.Builder<PieChartLegendOptions> legend) {
    PieChartLegendOptions legendResource = legend.build();
        this.internal.legend = legendResource;
        return this;
    }
    
    public OptionsBuilder orientation(VizOrientation orientation) {
        this.internal.orientation = orientation;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
