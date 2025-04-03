// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class ResourceNamesQueryBuilder implements com.grafana.foundation.cog.Builder<ResourceNamesQuery> {
    protected final ResourceNamesQuery internal;
    
    public ResourceNamesQueryBuilder() {
        this.internal = new ResourceNamesQuery();
        this.internal.kind = "ResourceNamesQuery";
    }
    public ResourceNamesQueryBuilder rawQuery(String rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public ResourceNamesQueryBuilder subscription(String subscription) {
        this.internal.subscription = subscription;
        return this;
    }
    
    public ResourceNamesQueryBuilder resourceGroup(String resourceGroup) {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public ResourceNamesQueryBuilder metricNamespace(String metricNamespace) {
        this.internal.metricNamespace = metricNamespace;
        return this;
    }
    public ResourceNamesQuery build() {
        return this.internal;
    }
}
