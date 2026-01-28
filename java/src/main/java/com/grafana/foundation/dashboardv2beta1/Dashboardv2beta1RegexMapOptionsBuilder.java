// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class Dashboardv2beta1RegexMapOptionsBuilder implements com.grafana.foundation.cog.Builder<Dashboardv2beta1RegexMapOptions> {
    protected final Dashboardv2beta1RegexMapOptions internal;
    
    public Dashboardv2beta1RegexMapOptionsBuilder() {
        this.internal = new Dashboardv2beta1RegexMapOptions();
    }
    public Dashboardv2beta1RegexMapOptionsBuilder pattern(String pattern) {
        this.internal.pattern = pattern;
        return this;
    }
    
    public Dashboardv2beta1RegexMapOptionsBuilder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result) {
    ValueMappingResult resultResource = result.build();
        this.internal.result = resultResource;
        return this;
    }
    public Dashboardv2beta1RegexMapOptions build() {
        return this.internal;
    }
}
