// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class SpecialValueMapBuilder implements com.grafana.foundation.cog.Builder<SpecialValueMap> {
    protected final SpecialValueMap internal;
    
    public SpecialValueMapBuilder() {
        this.internal = new SpecialValueMap();
    this.internal.type = "special";
    }
    public SpecialValueMapBuilder options(com.grafana.foundation.cog.Builder<DashboardSpecialValueMapOptions> options) {
    this.internal.options = options.build();
        return this;
    }
    public SpecialValueMap build() {
        return this.internal;
    }
}
