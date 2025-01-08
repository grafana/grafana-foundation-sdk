// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class AzureMetricQueryBuilder implements com.grafana.foundation.cog.Builder<AzureMetricQuery> {
    protected final AzureMetricQuery internal;
    
    public AzureMetricQueryBuilder() {
        this.internal = new AzureMetricQuery();
    }
    public AzureMetricQueryBuilder resources(com.grafana.foundation.cog.Builder<List<AzureMonitorResource>> resources) {
    this.internal.resources = resources.build();
        return this;
    }
    
    public AzureMetricQueryBuilder metricNamespace(String metricNamespace) {
    this.internal.metricNamespace = metricNamespace;
        return this;
    }
    
    public AzureMetricQueryBuilder customNamespace(String customNamespace) {
    this.internal.customNamespace = customNamespace;
        return this;
    }
    
    public AzureMetricQueryBuilder metricName(String metricName) {
    this.internal.metricName = metricName;
        return this;
    }
    
    public AzureMetricQueryBuilder region(String region) {
    this.internal.region = region;
        return this;
    }
    
    public AzureMetricQueryBuilder timeGrain(String timeGrain) {
    this.internal.timeGrain = timeGrain;
        return this;
    }
    
    public AzureMetricQueryBuilder aggregation(String aggregation) {
    this.internal.aggregation = aggregation;
        return this;
    }
    
    public AzureMetricQueryBuilder dimensionFilters(com.grafana.foundation.cog.Builder<List<AzureMetricDimension>> dimensionFilters) {
    this.internal.dimensionFilters = dimensionFilters.build();
        return this;
    }
    
    public AzureMetricQueryBuilder top(String top) {
    this.internal.top = top;
        return this;
    }
    
    public AzureMetricQueryBuilder allowedTimeGrainsMs(List<Long> allowedTimeGrainsMs) {
    this.internal.allowedTimeGrainsMs = allowedTimeGrainsMs;
        return this;
    }
    
    public AzureMetricQueryBuilder alias(String alias) {
    this.internal.alias = alias;
        return this;
    }
    
    public AzureMetricQueryBuilder timeGrainUnit(String timeGrainUnit) {
    this.internal.timeGrainUnit = timeGrainUnit;
        return this;
    }
    
    public AzureMetricQueryBuilder dimension(String dimension) {
    this.internal.dimension = dimension;
        return this;
    }
    
    public AzureMetricQueryBuilder dimensionFilter(String dimensionFilter) {
    this.internal.dimensionFilter = dimensionFilter;
        return this;
    }
    
    public AzureMetricQueryBuilder metricDefinition(String metricDefinition) {
    this.internal.metricDefinition = metricDefinition;
        return this;
    }
    
    public AzureMetricQueryBuilder resourceUri(String resourceUri) {
    this.internal.resourceUri = resourceUri;
        return this;
    }
    
    public AzureMetricQueryBuilder resourceGroup(String resourceGroup) {
    this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public AzureMetricQueryBuilder resourceName(String resourceName) {
    this.internal.resourceName = resourceName;
        return this;
    }
    public AzureMetricQuery build() {
        return this.internal;
    }
}
