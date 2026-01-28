// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.debug;


public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder mode(DebugMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public OptionsBuilder counters(com.grafana.foundation.cog.Builder<UpdateConfig> counters) {
    UpdateConfig countersResource = counters.build();
        this.internal.counters = countersResource;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
