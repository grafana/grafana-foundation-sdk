// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import com.grafana.foundation.common.DataSourceRef;

public class MonitorQueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final MonitorQuery internal;
    
    public MonitorQueryBuilder() {
        this.internal = new MonitorQuery();
    }
    public MonitorQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public MonitorQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public MonitorQueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public MonitorQueryBuilder subscription(String subscription) {
        this.internal.subscription = subscription;
        return this;
    }
    
    public MonitorQueryBuilder subscriptions(List<String> subscriptions) {
        this.internal.subscriptions = subscriptions;
        return this;
    }
    
    public MonitorQueryBuilder azureMonitor(com.grafana.foundation.cog.Builder<MetricQuery> azureMonitor) {
    MetricQuery azureMonitorResource = azureMonitor.build();
        this.internal.azureMonitor = azureMonitorResource;
        return this;
    }
    
    public MonitorQueryBuilder azureLogAnalytics(com.grafana.foundation.cog.Builder<LogsQuery> azureLogAnalytics) {
    LogsQuery azureLogAnalyticsResource = azureLogAnalytics.build();
        this.internal.azureLogAnalytics = azureLogAnalyticsResource;
        return this;
    }
    
    public MonitorQueryBuilder azureResourceGraph(com.grafana.foundation.cog.Builder<ResourceGraphQuery> azureResourceGraph) {
    ResourceGraphQuery azureResourceGraphResource = azureResourceGraph.build();
        this.internal.azureResourceGraph = azureResourceGraphResource;
        return this;
    }
    
    public MonitorQueryBuilder azureTraces(com.grafana.foundation.cog.Builder<TracesQuery> azureTraces) {
    TracesQuery azureTracesResource = azureTraces.build();
        this.internal.azureTraces = azureTracesResource;
        return this;
    }
    
    public MonitorQueryBuilder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn) {
        this.internal.grafanaTemplateVariableFn = grafanaTemplateVariableFn;
        return this;
    }
    
    public MonitorQueryBuilder resourceGroup(String resourceGroup) {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public MonitorQueryBuilder namespace(String namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    
    public MonitorQueryBuilder resource(String resource) {
        this.internal.resource = resource;
        return this;
    }
    
    public MonitorQueryBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public MonitorQueryBuilder customNamespace(String customNamespace) {
        this.internal.customNamespace = customNamespace;
        return this;
    }
    
    public MonitorQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public MonitorQueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    public MonitorQuery build() {
        return this.internal;
    }
}
