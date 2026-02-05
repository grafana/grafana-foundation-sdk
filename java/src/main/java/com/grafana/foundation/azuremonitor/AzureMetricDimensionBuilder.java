// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class AzureMetricDimensionBuilder implements com.grafana.foundation.cog.Builder<AzureMetricDimension> {
    protected final AzureMetricDimension internal;
    
    public AzureMetricDimensionBuilder() {
        this.internal = new AzureMetricDimension();
    }
    public AzureMetricDimensionBuilder dimension(String dimension) {
        this.internal.dimension = dimension;
        return this;
    }
    
    public AzureMetricDimensionBuilder operator(String operator) {
        this.internal.operator = operator;
        return this;
    }
    
    public AzureMetricDimensionBuilder filters(List<String> filters) {
        this.internal.filters = filters;
        return this;
    }
    
    public AzureMetricDimensionBuilder filter(String filter) {
        this.internal.filter = filter;
        return this;
    }
    public AzureMetricDimension build() {
        return this.internal;
    }
}
