// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class Dashboardv2beta1FieldConfigSourceOverridesBuilder implements com.grafana.foundation.cog.Builder<Dashboardv2beta1FieldConfigSourceOverrides> {
    protected final Dashboardv2beta1FieldConfigSourceOverrides internal;
    
    public Dashboardv2beta1FieldConfigSourceOverridesBuilder() {
        this.internal = new Dashboardv2beta1FieldConfigSourceOverrides();
    }
    public Dashboardv2beta1FieldConfigSourceOverridesBuilder matcher(MatcherConfig matcher) {
        this.internal.matcher = matcher;
        return this;
    }
    
    public Dashboardv2beta1FieldConfigSourceOverridesBuilder properties(List<DynamicConfigValue> properties) {
        this.internal.properties = properties;
        return this;
    }
    public Dashboardv2beta1FieldConfigSourceOverrides build() {
        return this.internal;
    }
}
