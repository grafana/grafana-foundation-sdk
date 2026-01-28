// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statetimeline;

import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import java.util.List;
import com.grafana.foundation.common.TimelineValueAlignment;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder showValue(VisibilityMode showValue) {
        this.internal.showValue = showValue;
        return this;
    }
    
    public OptionsBuilder rowHeight(Double rowHeight) {
        if (!(rowHeight <= 1)) {
            throw new IllegalArgumentException("rowHeight must be <= 1");
        }
        this.internal.rowHeight = rowHeight;
        return this;
    }
    
    public OptionsBuilder mergeValues(Boolean mergeValues) {
        this.internal.mergeValues = mergeValues;
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
    
    public OptionsBuilder timezone(List<String> timezone) {
        this.internal.timezone = timezone;
        return this;
    }
    
    public OptionsBuilder alignValue(TimelineValueAlignment alignValue) {
        this.internal.alignValue = alignValue;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
