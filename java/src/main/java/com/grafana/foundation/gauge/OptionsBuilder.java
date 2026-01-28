// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.gauge;

import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder showThresholdLabels(Boolean showThresholdLabels) {
        this.internal.showThresholdLabels = showThresholdLabels;
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
    
    public OptionsBuilder showThresholdMarkers(Boolean showThresholdMarkers) {
        this.internal.showThresholdMarkers = showThresholdMarkers;
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
