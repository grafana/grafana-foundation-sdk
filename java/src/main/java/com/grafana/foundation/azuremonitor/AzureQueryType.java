// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonValue;


// Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated
@JsonFormat(shape = JsonFormat.Shape.OBJECT)
public enum AzureQueryType {
    AZURE_MONITOR("Azure Monitor"),
    LOG_ANALYTICS("Azure Log Analytics"),
    AZURE_RESOURCE_GRAPH("Azure Resource Graph"),
    AZURE_TRACES("Azure Traces"),
    SUBSCRIPTIONS_QUERY("Azure Subscriptions"),
    RESOURCE_GROUPS_QUERY("Azure Resource Groups"),
    NAMESPACES_QUERY("Azure Namespaces"),
    RESOURCE_NAMES_QUERY("Azure Resource Names"),
    METRIC_NAMES_QUERY("Azure Metric Names"),
    WORKSPACES_QUERY("Azure Workspaces"),
    LOCATIONS_QUERY("Azure Regions"),
    GRAFANA_TEMPLATE_VARIABLE_FN("Grafana Template Variable Function"),
    _EMPTY("");

    private final String value;

    private AzureQueryType(String value) {
        this.value = value;
    }

    @JsonValue
    public String Value() {
        return value;
    }
}
