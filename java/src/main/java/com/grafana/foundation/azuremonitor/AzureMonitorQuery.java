// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

public class AzureMonitorQuery implements com.grafana.foundation.cog.variants.Dataquery {
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    @JsonProperty("refId")
    public String refId;
    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    // Specify the query flavor
    // TODO make this required and give it a default
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("queryType")
    public String queryType;
    // Azure subscription containing the resource(s) to be queried.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("subscription")
    public String subscription;
    // Subscriptions to be queried via Azure Resource Graph.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("subscriptions")
    public List<String> subscriptions;
    // Azure Monitor Metrics sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("azureMonitor")
    public AzureMetricQuery azureMonitor;
    // Azure Monitor Logs sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("azureLogAnalytics")
    public AzureLogsQuery azureLogAnalytics;
    // Azure Resource Graph sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("azureResourceGraph")
    public AzureResourceGraphQuery azureResourceGraph;
    // Application Insights Traces sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("azureTraces")
    public AzureTracesQuery azureTraces;
    // @deprecated Legacy template variable support.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("grafanaTemplateVariableFn")
    public GrafanaTemplateVariableQuery grafanaTemplateVariableFn;
    // Template variables params. These exist for backwards compatiblity with legacy template variables.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resourceGroup")
    public String resourceGroup;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("namespace")
    public String namespace;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resource")
    public String resource;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // Azure Monitor query type.
    // queryType: #AzureQueryType
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("region")
    public String region;
    public String dataqueryName() {
        return "grafana-azure-monitor-datasource";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AzureMonitorQuery> {
        protected final AzureMonitorQuery internal;
        
        public Builder() {
            this.internal = new AzureMonitorQuery();
        }
    public Builder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public Builder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public Builder subscription(String subscription) {
    this.internal.subscription = subscription;
        return this;
    }
    
    public Builder subscriptions(List<String> subscriptions) {
    this.internal.subscriptions = subscriptions;
        return this;
    }
    
    public Builder azureMonitor(com.grafana.foundation.cog.Builder<AzureMetricQuery> azureMonitor) {
    this.internal.azureMonitor = azureMonitor.build();
        return this;
    }
    
    public Builder azureLogAnalytics(com.grafana.foundation.cog.Builder<AzureLogsQuery> azureLogAnalytics) {
    this.internal.azureLogAnalytics = azureLogAnalytics.build();
        return this;
    }
    
    public Builder azureResourceGraph(com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> azureResourceGraph) {
    this.internal.azureResourceGraph = azureResourceGraph.build();
        return this;
    }
    
    public Builder azureTraces(com.grafana.foundation.cog.Builder<AzureTracesQuery> azureTraces) {
    this.internal.azureTraces = azureTraces.build();
        return this;
    }
    
    public Builder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn) {
    this.internal.grafanaTemplateVariableFn = grafanaTemplateVariableFn;
        return this;
    }
    
    public Builder resourceGroup(String resourceGroup) {
    this.internal.resourceGroup = resourceGroup;
        return this;
    }
    
    public Builder namespace(String namespace) {
    this.internal.namespace = namespace;
        return this;
    }
    
    public Builder resource(String resource) {
    this.internal.resource = resource;
        return this;
    }
    
    public Builder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public Builder region(String region) {
    this.internal.region = region;
        return this;
    }
    public AzureMonitorQuery build() {
            return this.internal;
        }
    }
}
