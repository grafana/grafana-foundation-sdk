// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class DashboardDashboardTimeBuilder implements com.grafana.foundation.cog.Builder<DashboardDashboardTime> {
    protected final DashboardDashboardTime internal;
    
    public DashboardDashboardTimeBuilder() {
        this.internal = new DashboardDashboardTime();
    }
    public DashboardDashboardTimeBuilder from(String from) {
    this.internal.from = from;
        return this;
    }
    
    public DashboardDashboardTimeBuilder to(String to) {
    this.internal.to = to;
        return this;
    }
    public DashboardDashboardTime build() {
        return this.internal;
    }
}
