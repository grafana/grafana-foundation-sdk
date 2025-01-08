// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class DashboardFieldConfigSourceOverridesBuilder implements com.grafana.foundation.cog.Builder<DashboardFieldConfigSourceOverrides> {
    protected final DashboardFieldConfigSourceOverrides internal;
    
    public DashboardFieldConfigSourceOverridesBuilder() {
        this.internal = new DashboardFieldConfigSourceOverrides();
    }
    public DashboardFieldConfigSourceOverridesBuilder matcher(MatcherConfig matcher) {
    this.internal.matcher = matcher;
        return this;
    }
    
    public DashboardFieldConfigSourceOverridesBuilder properties(List<DynamicConfigValue> properties) {
    this.internal.properties = properties;
        return this;
    }
    public DashboardFieldConfigSourceOverrides build() {
        return this.internal;
    }
}
