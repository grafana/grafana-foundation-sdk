// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class MetricDefinitionsQueryBuilder implements com.grafana.foundation.cog.Builder<MetricDefinitionsQuery> {
    protected final MetricDefinitionsQuery internal;
    
    public MetricDefinitionsQueryBuilder() {
        this.internal = new MetricDefinitionsQuery();
    this.internal.kind = "MetricDefinitionsQuery";
    }
    public MetricDefinitionsQueryBuilder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public MetricDefinitionsQueryBuilder subscription(String subscription) {
    this.internal.subscription = subscription;
        return this;
    }
    
    public MetricDefinitionsQueryBuilder resourceGroup(String resourceGroup) {
    this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public MetricDefinitionsQueryBuilder metricNamespace(String metricNamespace) {
    this.internal.metricNamespace = metricNamespace;
        return this;
    }
    
    public MetricDefinitionsQueryBuilder resourceName(String resourceName) {
    this.internal.resourceName = resourceName;
        return this;
    }
    public MetricDefinitionsQuery build() {
        return this.internal;
    }
}
