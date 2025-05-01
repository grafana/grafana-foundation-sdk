// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;
import java.util.LinkedList;

public class DashboardDashboardTemplatingBuilder implements com.grafana.foundation.cog.Builder<DashboardDashboardTemplating> {
    protected final DashboardDashboardTemplating internal;
    
    public DashboardDashboardTemplatingBuilder() {
        this.internal = new DashboardDashboardTemplating();
    }
    public DashboardDashboardTemplatingBuilder list(List<com.grafana.foundation.cog.Builder<VariableModel>> list) {
        List<VariableModel> listResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<VariableModel> r1 : list) {
                VariableModel listDepth1 = r1.build();
                listResources.add(listDepth1); 
        }
        this.internal.list = listResources;
        return this;
    }
    public DashboardDashboardTemplating build() {
        return this.internal;
    }
}
