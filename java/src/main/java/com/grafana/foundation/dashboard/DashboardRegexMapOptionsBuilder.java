// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class DashboardRegexMapOptionsBuilder implements com.grafana.foundation.cog.Builder<DashboardRegexMapOptions> {
    protected final DashboardRegexMapOptions internal;
    
    public DashboardRegexMapOptionsBuilder() {
        this.internal = new DashboardRegexMapOptions();
    }
    public DashboardRegexMapOptionsBuilder pattern(String pattern) {
    this.internal.pattern = pattern;
        return this;
    }
    
    public DashboardRegexMapOptionsBuilder result(ValueMappingResult result) {
    this.internal.result = result;
        return this;
    }
    public DashboardRegexMapOptions build() {
        return this.internal;
    }
}
