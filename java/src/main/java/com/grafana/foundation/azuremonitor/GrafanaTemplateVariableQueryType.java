// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum GrafanaTemplateVariableQueryType {
    APP_INSIGHTS_METRIC_NAME_QUERY("AppInsightsMetricNameQuery"),
    APP_INSIGHTS_GROUP_BY_QUERY("AppInsightsGroupByQuery"),
    SUBSCRIPTIONS_QUERY("SubscriptionsQuery"),
    RESOURCE_GROUPS_QUERY("ResourceGroupsQuery"),
    RESOURCE_NAMES_QUERY("ResourceNamesQuery"),
    METRIC_NAMESPACE_QUERY("MetricNamespaceQuery"),
    METRIC_NAMES_QUERY("MetricNamesQuery"),
    WORKSPACES_QUERY("WorkspacesQuery"),
    UNKNOWN_QUERY("UnknownQuery"),
    _EMPTY("");

    private final String value;

    private GrafanaTemplateVariableQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
