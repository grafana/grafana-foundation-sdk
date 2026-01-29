// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class Dashboardv2beta1RangeMapOptionsBuilder implements com.grafana.foundation.cog.Builder<Dashboardv2beta1RangeMapOptions> {
    protected final Dashboardv2beta1RangeMapOptions internal;
    
    public Dashboardv2beta1RangeMapOptionsBuilder() {
        this.internal = new Dashboardv2beta1RangeMapOptions();
    }
    public Dashboardv2beta1RangeMapOptionsBuilder from(Double from) {
        this.internal.from = from;
        return this;
    }
    
    public Dashboardv2beta1RangeMapOptionsBuilder to(Double to) {
        this.internal.to = to;
        return this;
    }
    
    public Dashboardv2beta1RangeMapOptionsBuilder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result) {
    ValueMappingResult resultResource = result.build();
        this.internal.result = resultResource;
        return this;
    }
    public Dashboardv2beta1RangeMapOptions build() {
        return this.internal;
    }
}
