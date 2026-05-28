// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;


public class Dashboardv2RegexMapOptionsBuilder implements com.grafana.foundation.cog.Builder<Dashboardv2RegexMapOptions> {
    protected final Dashboardv2RegexMapOptions internal;
    
    public Dashboardv2RegexMapOptionsBuilder() {
        this.internal = new Dashboardv2RegexMapOptions();
    }
    public Dashboardv2RegexMapOptionsBuilder pattern(String pattern) {
        this.internal.pattern = pattern;
        return this;
    }
    
    public Dashboardv2RegexMapOptionsBuilder result(com.grafana.foundation.cog.Builder<ValueMappingResult> result) {
    ValueMappingResult resultResource = result.build();
        this.internal.result = resultResource;
        return this;
    }
    public Dashboardv2RegexMapOptions build() {
        return this.internal;
    }
}
