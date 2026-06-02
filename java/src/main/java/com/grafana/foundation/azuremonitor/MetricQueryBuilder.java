// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class MetricQueryBuilder implements com.grafana.foundation.cog.Builder<MetricQuery> {
    protected final MetricQuery internal;
    
    public MetricQueryBuilder() {
        this.internal = new MetricQuery();
    }
    public MetricQueryBuilder resources(List<com.grafana.foundation.cog.Builder<MonitorResource>> resources) {
        List<MonitorResource> resourcesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<MonitorResource> r1 : resources) {
                MonitorResource resourcesDepth1 = r1.build();
                resourcesResources.add(resourcesDepth1); 
        }
        this.internal.resources = resourcesResources;
        return this;
    }
    
    public MetricQueryBuilder metricNamespace(String metricNamespace) {
        this.internal.metricNamespace = metricNamespace;
        return this;
    }
    
    public MetricQueryBuilder customNamespace(String customNamespace) {
        this.internal.customNamespace = customNamespace;
        return this;
    }
    
    public MetricQueryBuilder metricName(String metricName) {
        this.internal.metricName = metricName;
        return this;
    }
    
    public MetricQueryBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public MetricQueryBuilder timeGrain(String timeGrain) {
        this.internal.timeGrain = timeGrain;
        return this;
    }
    
    public MetricQueryBuilder aggregation(String aggregation) {
        this.internal.aggregation = aggregation;
        return this;
    }
    
    public MetricQueryBuilder dimensionFilters(List<com.grafana.foundation.cog.Builder<MetricDimension>> dimensionFilters) {
        List<MetricDimension> dimensionFiltersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<MetricDimension> r1 : dimensionFilters) {
                MetricDimension dimensionFiltersDepth1 = r1.build();
                dimensionFiltersResources.add(dimensionFiltersDepth1); 
        }
        this.internal.dimensionFilters = dimensionFiltersResources;
        return this;
    }
    
    public MetricQueryBuilder top(String top) {
        this.internal.top = top;
        return this;
    }
    
    public MetricQueryBuilder allowedTimeGrainsMs(List<Long> allowedTimeGrainsMs) {
        this.internal.allowedTimeGrainsMs = allowedTimeGrainsMs;
        return this;
    }
    
    public MetricQueryBuilder alias(String alias) {
        this.internal.alias = alias;
        return this;
    }
    
    public MetricQueryBuilder timeGrainUnit(String timeGrainUnit) {
        this.internal.timeGrainUnit = timeGrainUnit;
        return this;
    }
    
    public MetricQueryBuilder dimension(String dimension) {
        this.internal.dimension = dimension;
        return this;
    }
    
    public MetricQueryBuilder dimensionFilter(String dimensionFilter) {
        this.internal.dimensionFilter = dimensionFilter;
        return this;
    }
    
    public MetricQueryBuilder metricDefinition(String metricDefinition) {
        this.internal.metricDefinition = metricDefinition;
        return this;
    }
    
    public MetricQueryBuilder resourceUri(String resourceUri) {
        this.internal.resourceUri = resourceUri;
        return this;
    }
    
    public MetricQueryBuilder resourceGroup(String resourceGroup) {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public MetricQueryBuilder resourceName(String resourceName) {
        this.internal.resourceName = resourceName;
        return this;
    }
    public MetricQuery build() {
        return this.internal;
    }
}
