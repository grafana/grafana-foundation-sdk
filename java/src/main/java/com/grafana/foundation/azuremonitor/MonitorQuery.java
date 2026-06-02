// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class MonitorQuery implements com.grafana.foundation.cog.variants.Dataquery {
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
    // Also used for template variable queries
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
    public MetricQuery azureMonitor;
    // Azure Monitor Logs sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("azureLogAnalytics")
    public LogsQuery azureLogAnalytics;
    // Azure Resource Graph sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("azureResourceGraph")
    public ResourceGraphQuery azureResourceGraph;
    // Application Insights Traces sub-query properties.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("azureTraces")
    public TracesQuery azureTraces;
    // @deprecated Legacy template variable support.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("grafanaTemplateVariableFn")
    public GrafanaTemplateVariableQuery grafanaTemplateVariableFn;
    // Resource group used in template variable queries
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resourceGroup")
    public String resourceGroup;
    // Namespace used in template variable queries
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("namespace")
    public String namespace;
    // Resource used in template variable queries
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("resource")
    public String resource;
    // Region used in template variable queries
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("region")
    public String region;
    // Custom namespace used in template variable queries
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("customNamespace")
    public String customNamespace;
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("datasource")
    public Object datasource;
    // Used only for exemplar queries from Prometheus
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("query")
    public String query;
    public MonitorQuery() {
        this.refId = "";
    }
    public MonitorQuery(String refId,Boolean hide,String queryType,String subscription,List<String> subscriptions,MetricQuery azureMonitor,LogsQuery azureLogAnalytics,ResourceGraphQuery azureResourceGraph,TracesQuery azureTraces,GrafanaTemplateVariableQuery grafanaTemplateVariableFn,String resourceGroup,String namespace,String resource,String region,String customNamespace,Object datasource,String query) {
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
        this.customNamespace = customNamespace;
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
