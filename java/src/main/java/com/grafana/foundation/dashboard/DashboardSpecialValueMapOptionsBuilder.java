// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class DashboardSpecialValueMapOptionsBuilder implements com.grafana.foundation.cog.Builder<DashboardSpecialValueMapOptions> {
    protected final DashboardSpecialValueMapOptions internal;
    
    public DashboardSpecialValueMapOptionsBuilder() {
        this.internal = new DashboardSpecialValueMapOptions();
    }
    public DashboardSpecialValueMapOptionsBuilder match(SpecialValueMatch match) {
    this.internal.match = match;
        return this;
    }
    
    public DashboardSpecialValueMapOptionsBuilder result(ValueMappingResult result) {
    this.internal.result = result;
        return this;
    }
    public DashboardSpecialValueMapOptions build() {
        return this.internal;
    }
}
