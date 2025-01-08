// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class AzureMonitorResourceBuilder implements com.grafana.foundation.cog.Builder<AzureMonitorResource> {
    protected final AzureMonitorResource internal;
    
    public AzureMonitorResourceBuilder() {
        this.internal = new AzureMonitorResource();
    }
    public AzureMonitorResourceBuilder subscription(String subscription) {
    this.internal.subscription = subscription;
        return this;
    }
    
    public AzureMonitorResourceBuilder resourceGroup(String resourceGroup) {
    this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public AzureMonitorResourceBuilder resourceName(String resourceName) {
    this.internal.resourceName = resourceName;
        return this;
    }
    
    public AzureMonitorResourceBuilder metricNamespace(String metricNamespace) {
    this.internal.metricNamespace = metricNamespace;
        return this;
    }
    
    public AzureMonitorResourceBuilder region(String region) {
    this.internal.region = region;
        return this;
    }
    public AzureMonitorResource build() {
        return this.internal;
    }
}
