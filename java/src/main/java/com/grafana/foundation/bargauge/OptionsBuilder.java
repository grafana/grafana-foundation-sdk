// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bargauge;

import com.grafana.foundation.common.BarGaugeDisplayMode;
import com.grafana.foundation.common.BarGaugeValueMode;
import com.grafana.foundation.common.BarGaugeNamePlacement;
import com.grafana.foundation.common.BarGaugeSizing;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder displayMode(BarGaugeDisplayMode displayMode) {
        this.internal.displayMode = displayMode;
        return this;
    }
    
    public OptionsBuilder valueMode(BarGaugeValueMode valueMode) {
        this.internal.valueMode = valueMode;
        return this;
    }
    
    public OptionsBuilder namePlacement(BarGaugeNamePlacement namePlacement) {
        this.internal.namePlacement = namePlacement;
        return this;
    }
    
    public OptionsBuilder showUnfilled(Boolean showUnfilled) {
        this.internal.showUnfilled = showUnfilled;
        return this;
    }
    
    public OptionsBuilder sizing(BarGaugeSizing sizing) {
        this.internal.sizing = sizing;
        return this;
    }
    
    public OptionsBuilder minVizWidth(Integer minVizWidth) {
        this.internal.minVizWidth = minVizWidth;
        return this;
    }
    
    public OptionsBuilder minVizHeight(Integer minVizHeight) {
        this.internal.minVizHeight = minVizHeight;
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
    
    public OptionsBuilder maxVizHeight(Integer maxVizHeight) {
        this.internal.maxVizHeight = maxVizHeight;
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
