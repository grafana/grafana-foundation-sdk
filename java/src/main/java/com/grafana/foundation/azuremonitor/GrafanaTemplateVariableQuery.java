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
    public static GrafanaTemplateVariableQuery createAppInsightsMetricNameQuery(com.grafana.foundation.cog.Builder<AppInsightsMetricNameQuery> appInsightsMetricNameQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.appInsightsMetricNameQuery = appInsightsMetricNameQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createAppInsightsGroupByQuery(com.grafana.foundation.cog.Builder<AppInsightsGroupByQuery> appInsightsGroupByQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.appInsightsGroupByQuery = appInsightsGroupByQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createSubscriptionsQuery(com.grafana.foundation.cog.Builder<SubscriptionsQuery> subscriptionsQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.subscriptionsQuery = subscriptionsQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createResourceGroupsQuery(com.grafana.foundation.cog.Builder<ResourceGroupsQuery> resourceGroupsQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.resourceGroupsQuery = resourceGroupsQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createResourceNamesQuery(com.grafana.foundation.cog.Builder<ResourceNamesQuery> resourceNamesQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.resourceNamesQuery = resourceNamesQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createMetricNamespaceQuery(com.grafana.foundation.cog.Builder<MetricNamespaceQuery> metricNamespaceQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.metricNamespaceQuery = metricNamespaceQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createMetricDefinitionsQuery(com.grafana.foundation.cog.Builder<MetricDefinitionsQuery> metricDefinitionsQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.metricDefinitionsQuery = metricDefinitionsQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createMetricNamesQuery(com.grafana.foundation.cog.Builder<MetricNamesQuery> metricNamesQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.metricNamesQuery = metricNamesQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createWorkspacesQuery(com.grafana.foundation.cog.Builder<WorkspacesQuery> workspacesQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.workspacesQuery = workspacesQuery.build();
        return grafanaTemplateVariableQuery;
    }
    public static GrafanaTemplateVariableQuery createUnknownQuery(com.grafana.foundation.cog.Builder<UnknownQuery> unknownQuery) {
        GrafanaTemplateVariableQuery grafanaTemplateVariableQuery = new GrafanaTemplateVariableQuery();
        grafanaTemplateVariableQuery.unknownQuery = unknownQuery.build();
        return grafanaTemplateVariableQuery;
    }
    
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
