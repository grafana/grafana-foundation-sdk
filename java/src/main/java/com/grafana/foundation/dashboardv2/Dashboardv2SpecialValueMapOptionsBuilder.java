// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;


public class Dashboardv2SpecialValueMapOptionsBuilder implements com.grafana.foundation.cog.Builder<Dashboardv2SpecialValueMapOptions> {
    protected final Dashboardv2SpecialValueMapOptions internal;
    
    public Dashboardv2SpecialValueMapOptionsBuilder() {
        this.internal = new Dashboardv2SpecialValueMapOptions();
    }
    public Dashboardv2SpecialValueMapOptionsBuilder match(SpecialValueMatch match) {
        this.internal.match = match;
        return this;
    }
    
    public Dashboardv2SpecialValueMapOptionsBuilder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result) {
    ValueMappingResult resultResource = result.build();
        this.internal.result = resultResource;
        return this;
    }
    public Dashboardv2SpecialValueMapOptions build() {
        return this.internal;
    }
}
