// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class Dashboardv2beta1SpecialValueMapOptionsBuilder implements com.grafana.foundation.cog.Builder<Dashboardv2beta1SpecialValueMapOptions> {
    protected final Dashboardv2beta1SpecialValueMapOptions internal;
    
    public Dashboardv2beta1SpecialValueMapOptionsBuilder() {
        this.internal = new Dashboardv2beta1SpecialValueMapOptions();
    }
    public Dashboardv2beta1SpecialValueMapOptionsBuilder match(SpecialValueMatch match) {
        this.internal.match = match;
        return this;
    }
    
    public Dashboardv2beta1SpecialValueMapOptionsBuilder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result) {
    ValueMappingResult resultResource = result.build();
        this.internal.result = resultResource;
        return this;
    }
    public Dashboardv2beta1SpecialValueMapOptions build() {
        return this.internal;
    }
}
