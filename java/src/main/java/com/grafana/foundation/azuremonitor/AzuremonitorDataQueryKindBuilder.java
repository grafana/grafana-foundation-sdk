// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import java.util.List;
import com.grafana.foundation.common.DataSourceRef;

public class AzuremonitorDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public AzuremonitorDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-azure-monitor-datasource";
    }
    public AzuremonitorDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).refId = refId;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).hide = hide;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder subscription(String subscription) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).subscription = subscription;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder subscriptions(List<String> subscriptions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).subscriptions = subscriptions;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder azureMonitor(com.grafana.foundation.cog.Builder<AzureMetricQuery> azureMonitor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
    AzureMetricQuery azureMonitorResource = azureMonitor.build();
        ((AzureMonitorQuery) this.internal.spec).azureMonitor = azureMonitorResource;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder azureLogAnalytics(com.grafana.foundation.cog.Builder<AzureLogsQuery> azureLogAnalytics) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
    AzureLogsQuery azureLogAnalyticsResource = azureLogAnalytics.build();
        ((AzureMonitorQuery) this.internal.spec).azureLogAnalytics = azureLogAnalyticsResource;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder azureResourceGraph(com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> azureResourceGraph) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
    AzureResourceGraphQuery azureResourceGraphResource = azureResourceGraph.build();
        ((AzureMonitorQuery) this.internal.spec).azureResourceGraph = azureResourceGraphResource;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder azureTraces(com.grafana.foundation.cog.Builder<AzureTracesQuery> azureTraces) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
    AzureTracesQuery azureTracesResource = azureTraces.build();
        ((AzureMonitorQuery) this.internal.spec).azureTraces = azureTracesResource;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).grafanaTemplateVariableFn = grafanaTemplateVariableFn;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder resourceGroup(String resourceGroup) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).resourceGroup = resourceGroup;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder namespace(String namespace) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).namespace = namespace;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder resource(String resource) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).resource = resource;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder region(String region) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).region = region;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder customNamespace(String customNamespace) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).customNamespace = customNamespace;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder oldDatasource(DataSourceRef datasource) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.azuremonitor.AzureMonitorQueryBuilder().build();
		}
        ((AzureMonitorQuery) this.internal.spec).datasource = datasource;
        return this;
    }
    
    public AzuremonitorDataQueryKindBuilder query(String query) {
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
