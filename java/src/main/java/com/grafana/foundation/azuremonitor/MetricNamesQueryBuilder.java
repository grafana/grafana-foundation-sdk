// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class MetricNamesQueryBuilder implements com.grafana.foundation.cog.Builder<MetricNamesQuery> {
    protected final MetricNamesQuery internal;
    
    public MetricNamesQueryBuilder() {
        this.internal = new MetricNamesQuery();
    this.internal.kind = "MetricNamesQuery";
    }
    public MetricNamesQueryBuilder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public MetricNamesQueryBuilder subscription(String subscription) {
    this.internal.subscription = subscription;
        return this;
    }
    
    public MetricNamesQueryBuilder resourceGroup(String resourceGroup) {
    this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public MetricNamesQueryBuilder resourceName(String resourceName) {
    this.internal.resourceName = resourceName;
        return this;
    }
    
    public MetricNamesQueryBuilder metricNamespace(String metricNamespace) {
    this.internal.metricNamespace = metricNamespace;
        return this;
    }
    public MetricNamesQuery build() {
        return this.internal;
    }
}
