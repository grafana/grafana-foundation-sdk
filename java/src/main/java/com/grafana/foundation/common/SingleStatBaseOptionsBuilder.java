// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class SingleStatBaseOptionsBuilder implements com.grafana.foundation.cog.Builder<SingleStatBaseOptions> {
    protected final SingleStatBaseOptions internal;
    
    public SingleStatBaseOptionsBuilder() {
        this.internal = new SingleStatBaseOptions();
    }
    public SingleStatBaseOptionsBuilder reduceOptions(com.grafana.foundation.cog.Builder<ReduceDataOptions> reduceOptions) {
    ReduceDataOptions reduceOptionsResource = reduceOptions.build();
        this.internal.reduceOptions = reduceOptionsResource;
        return this;
    }
    
    public SingleStatBaseOptionsBuilder text(com.grafana.foundation.cog.Builder<VizTextDisplayOptions> text) {
    VizTextDisplayOptions textResource = text.build();
        this.internal.text = textResource;
        return this;
    }
    
    public SingleStatBaseOptionsBuilder orientation(VizOrientation orientation) {
        this.internal.orientation = orientation;
        return this;
    }
    public SingleStatBaseOptions build() {
        return this.internal;
    }
}
