---
title: <span class="badge object-type-enum"></span> QueryType
---
# <span class="badge object-type-enum"></span> QueryType

Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated

## Definition

```python
class QueryType(enum.StrEnum):
    """
    Defines the supported queryTypes. GrafanaTemplateVariableFn is deprecated
    """

    MONITOR = "Azure Monitor"
    LOG_ANALYTICS = "Azure Log Analytics"
    RESOURCE_GRAPH = "Azure Resource Graph"
    TRACES = "Azure Traces"
    SUBSCRIPTIONS_QUERY = "Azure Subscriptions"
    RESOURCE_GROUPS_QUERY = "Azure Resource Groups"
    NAMESPACES_QUERY = "Azure Namespaces"
    RESOURCE_NAMES_QUERY = "Azure Resource Names"
    METRIC_NAMES_QUERY = "Azure Metric Names"
    WORKSPACES_QUERY = "Azure Workspaces"
    LOCATIONS_QUERY = "Azure Regions"
    GRAFANA_TEMPLATE_VARIABLE_FN = "Grafana Template Variable Function"
    TRACE_EXEMPLAR = "traceql"
    CUSTOM_NAMESPACES_QUERY = "Azure Custom Namespaces"
    CUSTOM_METRIC_NAMES_QUERY = "Azure Custom Metric Names"
```
