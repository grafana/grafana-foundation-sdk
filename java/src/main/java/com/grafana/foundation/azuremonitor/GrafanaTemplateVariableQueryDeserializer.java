// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class GrafanaTemplateVariableQueryDeserializer extends JsonDeserializer<GrafanaTemplateVariableQuery> {

    @Override
    public GrafanaTemplateVariableQuery deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        if (!root.has("kind")) {
            throw new IOException("Cannot find discriminator for GrafanaTemplateVariableQuery");
        }
        String discriminator = root.get("kind").asText();  
        
        switch (discriminator) {
        case "AppInsightsGroupByQuery":
            grafanaTemplateVariableQuery.appInsightsGroupByQuery = mapper.convertValue(root, AppInsightsGroupByQuery.class);
            break;
        case "AppInsightsMetricNameQuery":
            grafanaTemplateVariableQuery.appInsightsMetricNameQuery = mapper.convertValue(root, AppInsightsMetricNameQuery.class);
            break;
        case "MetricDefinitionsQuery":
            grafanaTemplateVariableQuery.metricDefinitionsQuery = mapper.convertValue(root, MetricDefinitionsQuery.class);
            break;
        case "MetricNamesQuery":
            grafanaTemplateVariableQuery.metricNamesQuery = mapper.convertValue(root, MetricNamesQuery.class);
            break;
        case "MetricNamespaceQuery":
            grafanaTemplateVariableQuery.metricNamespaceQuery = mapper.convertValue(root, MetricNamespaceQuery.class);
            break;
        case "ResourceGroupsQuery":
            grafanaTemplateVariableQuery.resourceGroupsQuery = mapper.convertValue(root, ResourceGroupsQuery.class);
            break;
        case "ResourceNamesQuery":
            grafanaTemplateVariableQuery.resourceNamesQuery = mapper.convertValue(root, ResourceNamesQuery.class);
            break;
        case "SubscriptionsQuery":
            grafanaTemplateVariableQuery.subscriptionsQuery = mapper.convertValue(root, SubscriptionsQuery.class);
            break;
        case "UnknownQuery":
            grafanaTemplateVariableQuery.unknownQuery = mapper.convertValue(root, UnknownQuery.class);
            break;
        case "WorkspacesQuery":
            grafanaTemplateVariableQuery.workspacesQuery = mapper.convertValue(root, WorkspacesQuery.class);
            break;
        }
        
        return grafanaTemplateVariableQuery;
    }
}
