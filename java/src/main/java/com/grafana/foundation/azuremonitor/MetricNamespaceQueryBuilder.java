// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class MetricNamespaceQueryBuilder implements com.grafana.foundation.cog.Builder<MetricNamespaceQuery> {
    protected final MetricNamespaceQuery internal;
    
    public MetricNamespaceQueryBuilder() {
        this.internal = new MetricNamespaceQuery();
    this.internal.kind = "MetricNamespaceQuery";
    }
    public MetricNamespaceQueryBuilder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public MetricNamespaceQueryBuilder subscription(String subscription) {
    this.internal.subscription = subscription;
        return this;
    }
    
    public MetricNamespaceQueryBuilder resourceGroup(String resourceGroup) {
    this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public MetricNamespaceQueryBuilder metricNamespace(String metricNamespace) {
    this.internal.metricNamespace = metricNamespace;
        return this;
    }
    
    public MetricNamespaceQueryBuilder resourceName(String resourceName) {
    this.internal.resourceName = resourceName;
        return this;
    }
    public MetricNamespaceQuery build() {
        return this.internal;
    }
}
