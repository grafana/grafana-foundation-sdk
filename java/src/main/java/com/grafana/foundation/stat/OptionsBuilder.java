// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.stat;

import com.grafana.foundation.common.BigValueGraphMode;
import com.grafana.foundation.common.BigValueColorMode;
import com.grafana.foundation.common.BigValueJustifyMode;
import com.grafana.foundation.common.BigValueTextMode;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.PercentChangeColorMode;
import com.grafana.foundation.common.VizOrientation;

public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder graphMode(BigValueGraphMode graphMode) {
        this.internal.graphMode = graphMode;
        return this;
    }
    
    public OptionsBuilder colorMode(BigValueColorMode colorMode) {
        this.internal.colorMode = colorMode;
        return this;
    }
    
    public OptionsBuilder justifyMode(BigValueJustifyMode justifyMode) {
        this.internal.justifyMode = justifyMode;
        return this;
    }
    
    public OptionsBuilder textMode(BigValueTextMode textMode) {
        this.internal.textMode = textMode;
        return this;
    }
    
    public OptionsBuilder wideLayout(Boolean wideLayout) {
        this.internal.wideLayout = wideLayout;
        return this;
    }
    
    public OptionsBuilder showPercentChange(Boolean showPercentChange) {
        this.internal.showPercentChange = showPercentChange;
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
    
    public OptionsBuilder percentChangeColorMode(PercentChangeColorMode percentChangeColorMode) {
        this.internal.percentChangeColorMode = percentChangeColorMode;
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
