// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class MonitorResourceBuilder implements com.grafana.foundation.cog.Builder<MonitorResource> {
    protected final MonitorResource internal;
    
    public MonitorResourceBuilder() {
        this.internal = new MonitorResource();
    }
    public MonitorResourceBuilder subscription(String subscription) {
        this.internal.subscription = subscription;
        return this;
    }
    
    public MonitorResourceBuilder resourceGroup(String resourceGroup) {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public MonitorResourceBuilder resourceName(String resourceName) {
        this.internal.resourceName = resourceName;
        return this;
    }
    
    public MonitorResourceBuilder metricNamespace(String metricNamespace) {
        this.internal.metricNamespace = metricNamespace;
        return this;
    }
    
    public MonitorResourceBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    public MonitorResource build() {
        return this.internal;
    }
}
