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
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("region")
    public String region;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public DataSourceRef datasource;
    // Used only for exemplar queries from Prometheus
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("query")
    public String query;
    public AzureMonitorQuery() {
    }
    
    public AzureMonitorQuery(String refId,Boolean hide,String queryType,String subscription,List<String> subscriptions,AzureMetricQuery azureMonitor,AzureLogsQuery azureLogAnalytics,AzureResourceGraphQuery azureResourceGraph,AzureTracesQuery azureTraces,GrafanaTemplateVariableQuery grafanaTemplateVariableFn,String resourceGroup,String namespace,String resource,String region,DataSourceRef datasource,String query) {
        this.refId = refId;
        this.hide = hide;
        this.queryType = queryType;
        this.subscription = subscription;
        this.subscriptions = subscriptions;
        this.azureMonitor = azureMonitor;
        this.azureLogAnalytics = azureLogAnalytics;
        this.azureResourceGraph = azureResourceGraph;
        this.azureTraces = azureTraces;
        this.grafanaTemplateVariableFn = grafanaTemplateVariableFn;
        this.resourceGroup = resourceGroup;
        this.namespace = namespace;
        this.resource = resource;
        this.region = region;
        this.datasource = datasource;
        this.query = query;
    }
    public String dataqueryName() {
        return "grafana-azure-monitor-datasource";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
