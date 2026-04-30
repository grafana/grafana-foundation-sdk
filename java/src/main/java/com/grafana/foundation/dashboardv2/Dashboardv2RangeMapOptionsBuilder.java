// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;


public class Dashboardv2RangeMapOptionsBuilder implements com.grafana.foundation.cog.Builder<Dashboardv2RangeMapOptions> {
    protected final Dashboardv2RangeMapOptions internal;
    
    public Dashboardv2RangeMapOptionsBuilder() {
        this.internal = new Dashboardv2RangeMapOptions();
    }
    public Dashboardv2RangeMapOptionsBuilder from(Double from) {
        this.internal.from = from;
        return this;
    }
    
    public Dashboardv2RangeMapOptionsBuilder to(Double to) {
        this.internal.to = to;
        return this;
    }
    
    public Dashboardv2RangeMapOptionsBuilder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result) {
    ValueMappingResult resultResource = result.build();
        this.internal.result = resultResource;
        return this;
    }
    public Dashboardv2RangeMapOptions build() {
        return this.internal;
    }
}
