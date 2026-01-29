// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class MetricFindValueBuilder implements com.grafana.foundation.cog.Builder<MetricFindValue> {
    protected final MetricFindValue internal;
    
    public MetricFindValueBuilder() {
        this.internal = new MetricFindValue();
    }
    public MetricFindValueBuilder text(String text) {
        this.internal.text = text;
        return this;
    }
    
    public MetricFindValueBuilder value(StringOrFloat64 value) {
        this.internal.value = value;
        return this;
    }
    
    public MetricFindValueBuilder group(String group) {
        this.internal.group = group;
        return this;
    }
    
    public MetricFindValueBuilder expandable(Boolean expandable) {
        this.internal.expandable = expandable;
        return this;
    }
    public MetricFindValue build() {
        return this.internal;
    }
}
