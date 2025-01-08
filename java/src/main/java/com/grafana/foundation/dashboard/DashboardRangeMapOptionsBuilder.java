// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class DashboardRangeMapOptionsBuilder implements com.grafana.foundation.cog.Builder<DashboardRangeMapOptions> {
    protected final DashboardRangeMapOptions internal;
    
    public DashboardRangeMapOptionsBuilder() {
        this.internal = new DashboardRangeMapOptions();
    }
    public DashboardRangeMapOptionsBuilder from(Double from) {
    this.internal.from = from;
        return this;
    }
    
    public DashboardRangeMapOptionsBuilder to(Double to) {
    this.internal.to = to;
        return this;
    }
    
    public DashboardRangeMapOptionsBuilder result(ValueMappingResult result) {
    this.internal.result = result;
        return this;
    }
    public DashboardRangeMapOptions build() {
        return this.internal;
    }
}
