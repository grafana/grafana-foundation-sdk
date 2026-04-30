// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import java.util.List;

public class AzuremonitorQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public AzuremonitorQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-azure-monitor-datasource";
    }
    public AzuremonitorQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public AzuremonitorQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public AzuremonitorQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public AzuremonitorQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).refId = refId;
        return this;
    }
    
    public AzuremonitorQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).hide = hide;
        return this;
    }
    
    public AzuremonitorQueryBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public AzuremonitorQueryBuilder subscription(String subscription) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).subscription = subscription;
        return this;
    }
    
    public AzuremonitorQueryBuilder subscriptions(List<String> subscriptions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).subscriptions = subscriptions;
        return this;
    }
    
    public AzuremonitorQueryBuilder azureMonitor(com.grafana.foundation.cog.Builder<AzureMetricQuery> azureMonitor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
    AzureMetricQuery azureMonitorResource = azureMonitor.build();
        ((AzureMonitorQuery) this.internal.spec).azureMonitor = azureMonitorResource;
        return this;
    }
    
    public AzuremonitorQueryBuilder azureLogAnalytics(com.grafana.foundation.cog.Builder<AzureLogsQuery> azureLogAnalytics) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
    AzureLogsQuery azureLogAnalyticsResource = azureLogAnalytics.build();
        ((AzureMonitorQuery) this.internal.spec).azureLogAnalytics = azureLogAnalyticsResource;
        return this;
    }
    
    public AzuremonitorQueryBuilder azureResourceGraph(com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> azureResourceGraph) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
    AzureResourceGraphQuery azureResourceGraphResource = azureResourceGraph.build();
        ((AzureMonitorQuery) this.internal.spec).azureResourceGraph = azureResourceGraphResource;
        return this;
    }
    
    public AzuremonitorQueryBuilder azureTraces(com.grafana.foundation.cog.Builder<AzureTracesQuery> azureTraces) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
    AzureTracesQuery azureTracesResource = azureTraces.build();
        ((AzureMonitorQuery) this.internal.spec).azureTraces = azureTracesResource;
        return this;
    }
    
    public AzuremonitorQueryBuilder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).grafanaTemplateVariableFn = grafanaTemplateVariableFn;
        return this;
    }
    
    public AzuremonitorQueryBuilder resourceGroup(String resourceGroup) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).resourceGroup = resourceGroup;
        return this;
    }
    
    public AzuremonitorQueryBuilder namespace(String namespace) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).namespace = namespace;
        return this;
    }
    
    public AzuremonitorQueryBuilder resource(String resource) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).resource = resource;
        return this;
    }
    
    public AzuremonitorQueryBuilder region(String region) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).region = region;
        return this;
    }
    
    public AzuremonitorQueryBuilder customNamespace(String customNamespace) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).customNamespace = customNamespace;
        return this;
    }
    
    public AzuremonitorQueryBuilder query(String query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).query = query;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
