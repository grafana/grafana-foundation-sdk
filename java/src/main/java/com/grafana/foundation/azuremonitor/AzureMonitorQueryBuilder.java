// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import com.grafana.foundation.common.DataSourceRef;

public class AzureMonitorQueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final AzureMonitorQuery internal;
    
    public AzureMonitorQueryBuilder() {
        this.internal = new AzureMonitorQuery();
    }
    public AzureMonitorQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public AzureMonitorQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public AzureMonitorQueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public AzureMonitorQueryBuilder subscription(String subscription) {
        this.internal.subscription = subscription;
        return this;
    }
    
    public AzureMonitorQueryBuilder subscriptions(List<String> subscriptions) {
        this.internal.subscriptions = subscriptions;
        return this;
    }
    
    public AzureMonitorQueryBuilder azureMonitor(com.grafana.foundation.cog.Builder<AzureMetricQuery> azureMonitor) {
    AzureMetricQuery azureMonitorResource = azureMonitor.build();
        this.internal.azureMonitor = azureMonitorResource;
        return this;
    }
    
    public AzureMonitorQueryBuilder azureLogAnalytics(com.grafana.foundation.cog.Builder<AzureLogsQuery> azureLogAnalytics) {
    AzureLogsQuery azureLogAnalyticsResource = azureLogAnalytics.build();
        this.internal.azureLogAnalytics = azureLogAnalyticsResource;
        return this;
    }
    
    public AzureMonitorQueryBuilder azureResourceGraph(com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> azureResourceGraph) {
    AzureResourceGraphQuery azureResourceGraphResource = azureResourceGraph.build();
        this.internal.azureResourceGraph = azureResourceGraphResource;
        return this;
    }
    
    public AzureMonitorQueryBuilder azureTraces(com.grafana.foundation.cog.Builder<AzureTracesQuery> azureTraces) {
    AzureTracesQuery azureTracesResource = azureTraces.build();
        this.internal.azureTraces = azureTracesResource;
        return this;
    }
    
    public AzureMonitorQueryBuilder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn) {
        this.internal.grafanaTemplateVariableFn = grafanaTemplateVariableFn;
        return this;
    }
    
    public AzureMonitorQueryBuilder resourceGroup(String resourceGroup) {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public AzureMonitorQueryBuilder namespace(String namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    
    public AzureMonitorQueryBuilder resource(String resource) {
        this.internal.resource = resource;
        return this;
    }
    
    public AzureMonitorQueryBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public AzureMonitorQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public AzureMonitorQueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    public AzureMonitorQuery build() {
        return this.internal;
    }
}
