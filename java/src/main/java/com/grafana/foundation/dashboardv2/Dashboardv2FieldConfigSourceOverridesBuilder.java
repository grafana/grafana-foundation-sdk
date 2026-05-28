// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import java.util.List;

public class Dashboardv2FieldConfigSourceOverridesBuilder implements com.grafana.foundation.cog.Builder<Dashboardv2FieldConfigSourceOverrides> {
    protected final Dashboardv2FieldConfigSourceOverrides internal;
    
    public Dashboardv2FieldConfigSourceOverridesBuilder() {
        this.internal = new Dashboardv2FieldConfigSourceOverrides();
    }
    public Dashboardv2FieldConfigSourceOverridesBuilder systemRef(String systemRef) {
        this.internal.systemRef = systemRef;
        return this;
    }
    
    public Dashboardv2FieldConfigSourceOverridesBuilder matcher(MatcherConfig matcher) {
        this.internal.matcher = matcher;
        return this;
    }
    
    public Dashboardv2FieldConfigSourceOverridesBuilder properties(List<DynamicConfigValue> properties) {
        this.internal.properties = properties;
        return this;
    }
    public Dashboardv2FieldConfigSourceOverrides build() {
        return this.internal;
    }
}
