// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class DashboardDashboardTemplatingBuilder implements com.grafana.foundation.cog.Builder<DashboardDashboardTemplating> {
    protected final DashboardDashboardTemplating internal;
    
    public DashboardDashboardTemplatingBuilder() {
        this.internal = new DashboardDashboardTemplating();
    }
    public DashboardDashboardTemplatingBuilder list(com.grafana.foundation.cog.Builder<List<VariableModel>> list) {
        this.internal.list = list.build();
        return this;
    }
    public DashboardDashboardTemplating build() {
        return this.internal;
    }
}
