// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = GrafanaTemplateVariableQueryDeserializer.class)
public class GrafanaTemplateVariableQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected AppInsightsMetricNameQuery appInsightsMetricNameQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected AppInsightsGroupByQuery appInsightsGroupByQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected SubscriptionsQuery subscriptionsQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected ResourceGroupsQuery resourceGroupsQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected ResourceNamesQuery resourceNamesQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected MetricNamespaceQuery metricNamespaceQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected MetricDefinitionsQuery metricDefinitionsQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected MetricNamesQuery metricNamesQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected WorkspacesQuery workspacesQuery;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected UnknownQuery unknownQuery;
    protected GrafanaTemplateVariableQuery() {}
    public static GrafanaTemplateVariableQuery createAppInsightsMetricNameQuery(AppInsightsMetricNameQuery appInsightsMetricNameQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.appInsightsMetricNameQuery = appInsightsMetricNameQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createAppInsightsGroupByQuery(AppInsightsGroupByQuery appInsightsGroupByQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.appInsightsGroupByQuery = appInsightsGroupByQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createSubscriptionsQuery(SubscriptionsQuery subscriptionsQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.subscriptionsQuery = subscriptionsQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createResourceGroupsQuery(ResourceGroupsQuery resourceGroupsQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.resourceGroupsQuery = resourceGroupsQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createResourceNamesQuery(ResourceNamesQuery resourceNamesQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.resourceNamesQuery = resourceNamesQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createMetricNamespaceQuery(MetricNamespaceQuery metricNamespaceQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.metricNamespaceQuery = metricNamespaceQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createMetricDefinitionsQuery(MetricDefinitionsQuery metricDefinitionsQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.metricDefinitionsQuery = metricDefinitionsQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createMetricNamesQuery(MetricNamesQuery metricNamesQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.metricNamesQuery = metricNamesQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createWorkspacesQuery(WorkspacesQuery workspacesQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.workspacesQuery = workspacesQuery;
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createUnknownQuery(UnknownQuery unknownQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.unknownQuery = unknownQuery;
        return grafanaTemplateVariableQuery;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
