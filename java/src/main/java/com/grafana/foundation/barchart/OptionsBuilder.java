// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.barchart;

import com.grafana.foundation.common.VizOrientation;
import com.grafana.foundation.common.StackingMode;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder xField(String xField) {
        this.internal.xField = xField;
        return this;
    }
    
    public OptionsBuilder colorByField(String colorByField) {
        this.internal.colorByField = colorByField;
        return this;
    }
    
    public OptionsBuilder orientation(VizOrientation orientation) {
        this.internal.orientation = orientation;
        return this;
    }
    
    public OptionsBuilder barRadius(Double barRadius) {
        if (!(barRadius >= 0)) {
            throw new IllegalArgumentException("barRadius must be >= 0");
        }
        if (!(barRadius <= 0.5)) {
            throw new IllegalArgumentException("barRadius must be <= 0.5");
        }
        this.internal.barRadius = barRadius;
        return this;
    }
    
    public OptionsBuilder xTickLabelRotation(Integer xTickLabelRotation) {
        if (!(xTickLabelRotation >= -90)) {
            throw new IllegalArgumentException("xTickLabelRotation must be >= -90");
        }
        if (!(xTickLabelRotation <= 90)) {
            throw new IllegalArgumentException("xTickLabelRotation must be <= 90");
        }
        this.internal.xTickLabelRotation = xTickLabelRotation;
        return this;
    }
    
    public OptionsBuilder xTickLabelMaxLength(Integer xTickLabelMaxLength) {
        if (!(xTickLabelMaxLength >= 0)) {
            throw new IllegalArgumentException("xTickLabelMaxLength must be >= 0");
        }
        this.internal.xTickLabelMaxLength = xTickLabelMaxLength;
        return this;
    }
    
    public OptionsBuilder xTickLabelSpacing(Integer xTickLabelSpacing) {
        this.internal.xTickLabelSpacing = xTickLabelSpacing;
        return this;
    }
    
    public OptionsBuilder stacking(StackingMode stacking) {
        this.internal.stacking = stacking;
        return this;
    }
    
    public OptionsBuilder showValue(VisibilityMode showValue) {
        this.internal.showValue = showValue;
        return this;
    }
    
    public OptionsBuilder barWidth(Double barWidth) {
        if (!(barWidth >= 0)) {
            throw new IllegalArgumentException("barWidth must be >= 0");
        }
        if (!(barWidth <= 1)) {
            throw new IllegalArgumentException("barWidth must be <= 1");
        }
        this.internal.barWidth = barWidth;
        return this;
    }
    
    public OptionsBuilder groupWidth(Double groupWidth) {
        if (!(groupWidth >= 0)) {
            throw new IllegalArgumentException("groupWidth must be >= 0");
        }
        if (!(groupWidth <= 1)) {
            throw new IllegalArgumentException("groupWidth must be <= 1");
        }
        this.internal.groupWidth = groupWidth;
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
    
    public OptionsBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
    VizTextDisplayOptions textResource = text.build();
        this.internal.text = textResource;
        return this;
    }
    
    public OptionsBuilder fullHighlight(Boolean fullHighlight) {
        this.internal.fullHighlight = fullHighlight;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
