// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class MetricDimensionBuilder implements com.grafana.foundation.cog.Builder<MetricDimension> {
    protected final MetricDimension internal;
    
    public MetricDimensionBuilder() {
        this.internal = new MetricDimension();
    }
    public MetricDimensionBuilder dimension(String dimension) {
        this.internal.dimension = dimension;
        return this;
    }
    
    public MetricDimensionBuilder operator(String operator) {
        this.internal.operator = operator;
        return this;
    }
    
    public MetricDimensionBuilder filters(List<String> filters) {
        this.internal.filters = filters;
        return this;
    }
    
    public MetricDimensionBuilder filter(String filter) {
        this.internal.filter = filter;
        return this;
    }
    public MetricDimension build() {
        return this.internal;
    }
}
