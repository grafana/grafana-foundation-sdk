// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statushistory;

import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import java.util.List;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder rowHeight(Float rowHeight) {
        if (!(rowHeight >= 0)) {
            throw new IllegalArgumentException("rowHeight must be >= 0");
        }
        if (!(rowHeight <= 1)) {
            throw new IllegalArgumentException("rowHeight must be <= 1");
        }
        this.internal.rowHeight = rowHeight;
        return this;
    }
    
    public OptionsBuilder showValue(VisibilityMode showValue) {
        this.internal.showValue = showValue;
        return this;
    }
    
    public OptionsBuilder colWidth(Double colWidth) {
        if (!(colWidth <= 1)) {
            throw new IllegalArgumentException("colWidth must be <= 1");
        }
        this.internal.colWidth = colWidth;
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
    
    public OptionsBuilder perPage(Double perPage) {
        if (!(perPage >= 1)) {
            throw new IllegalArgumentException("perPage must be >= 1");
        }
        this.internal.perPage = perPage;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
