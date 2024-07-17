// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = GrafanaTemplateVariableQueryDeserializer.class)
public class GrafanaTemplateVariableQuery { 
    @JsonUnwrapped
    public AppInsightsMetricNameQuery appInsightsMetricNameQuery; 
    @JsonUnwrapped
    public AppInsightsGroupByQuery appInsightsGroupByQuery; 
    @JsonUnwrapped
    public SubscriptionsQuery subscriptionsQuery; 
    @JsonUnwrapped
    public ResourceGroupsQuery resourceGroupsQuery; 
    @JsonUnwrapped
    public ResourceNamesQuery resourceNamesQuery; 
    @JsonUnwrapped
    public MetricNamespaceQuery metricNamespaceQuery; 
    @JsonUnwrapped
    public MetricDefinitionsQuery metricDefinitionsQuery; 
    @JsonUnwrapped
    public MetricNamesQuery metricNamesQuery; 
    @JsonUnwrapped
    public WorkspacesQuery workspacesQuery; 
    @JsonUnwrapped
    public UnknownQuery unknownQuery;
    
    public String toJSON() throws JsonProcessingException {
        if (appInsightsMetricNameQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(appInsightsMetricNameQuery);
        }
        if (appInsightsGroupByQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(appInsightsGroupByQuery);
        }
        if (subscriptionsQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(subscriptionsQuery);
        }
        if (resourceGroupsQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(resourceGroupsQuery);
        }
        if (resourceNamesQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(resourceNamesQuery);
        }
        if (metricNamespaceQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(metricNamespaceQuery);
        }
        if (metricDefinitionsQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(metricDefinitionsQuery);
        }
        if (metricNamesQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(metricNamesQuery);
        }
        if (workspacesQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(workspacesQuery);
        }
        if (unknownQuery != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(unknownQuery);
        }
        
        return null;
    }

}
