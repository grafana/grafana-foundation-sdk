# <span class="badge package-variant-dataquery"></span> azuremonitor

## Objects

 * <span class="badge object-type-class"></span> [AppInsightsGroupByQuery](./object-AppInsightsGroupByQuery.md)
 * <span class="badge object-type-class"></span> [AppInsightsMetricNameQuery](./object-AppInsightsMetricNameQuery.md)
 * <span class="badge object-type-class"></span> [BaseGrafanaTemplateVariableQuery](./object-BaseGrafanaTemplateVariableQuery.md)
 * <span class="badge object-type-disjunction"></span> [GrafanaTemplateVariableQuery](./object-GrafanaTemplateVariableQuery.md)
 * <span class="badge object-type-enum"></span> [GrafanaTemplateVariableQueryType](./object-GrafanaTemplateVariableQueryType.md)
 * <span class="badge object-type-class"></span> [LogsQuery](./object-LogsQuery.md)
 * <span class="badge object-type-class"></span> [MetricDefinitionsQuery](./object-MetricDefinitionsQuery.md)
 * <span class="badge object-type-class"></span> [MetricDimension](./object-MetricDimension.md)
 * <span class="badge object-type-class"></span> [MetricNamesQuery](./object-MetricNamesQuery.md)
 * <span class="badge object-type-class"></span> [MetricNamespaceQuery](./object-MetricNamespaceQuery.md)
 * <span class="badge object-type-class"></span> [MetricQuery](./object-MetricQuery.md)
 * <span class="badge object-type-class"></span> [MonitorQuery](./object-MonitorQuery.md)
 * <span class="badge object-type-class"></span> [MonitorResource](./object-MonitorResource.md)
 * <span class="badge object-type-enum"></span> [QueryType](./object-QueryType.md)
 * <span class="badge object-type-class"></span> [ResourceGraphQuery](./object-ResourceGraphQuery.md)
 * <span class="badge object-type-class"></span> [ResourceGroupsQuery](./object-ResourceGroupsQuery.md)
 * <span class="badge object-type-class"></span> [ResourceNamesQuery](./object-ResourceNamesQuery.md)
 * <span class="badge object-type-enum"></span> [ResultFormat](./object-ResultFormat.md)
 * <span class="badge object-type-class"></span> [SubscriptionsQuery](./object-SubscriptionsQuery.md)
 * <span class="badge object-type-class"></span> [TracesFilter](./object-TracesFilter.md)
 * <span class="badge object-type-class"></span> [TracesQuery](./object-TracesQuery.md)
 * <span class="badge object-type-class"></span> [UnknownQuery](./object-UnknownQuery.md)
 * <span class="badge object-type-class"></span> [WorkspacesQuery](./object-WorkspacesQuery.md)
## Builders

 * <span class="badge builder"></span> [AppInsightsGroupByQuery](./builder-AppInsightsGroupByQuery.md)
 * <span class="badge builder"></span> [AppInsightsMetricNameQuery](./builder-AppInsightsMetricNameQuery.md)
 * <span class="badge builder"></span> [BaseGrafanaTemplateVariableQuery](./builder-BaseGrafanaTemplateVariableQuery.md)
 * <span class="badge builder"></span> [LogsQuery](./builder-LogsQuery.md)
 * <span class="badge builder"></span> [MetricDefinitionsQuery](./builder-MetricDefinitionsQuery.md)
 * <span class="badge builder"></span> [MetricDimension](./builder-MetricDimension.md)
 * <span class="badge builder"></span> [MetricNamesQuery](./builder-MetricNamesQuery.md)
 * <span class="badge builder"></span> [MetricNamespaceQuery](./builder-MetricNamespaceQuery.md)
 * <span class="badge builder"></span> [MetricQuery](./builder-MetricQuery.md)
 * <span class="badge builder"></span> [MonitorQuery](./builder-MonitorQuery.md)
 * <span class="badge builder"></span> [MonitorResource](./builder-MonitorResource.md)
 * <span class="badge builder"></span> [Query](./builder-Query.md)
 * <span class="badge builder"></span> [QueryV2](./builder-QueryV2.md)
 * <span class="badge builder"></span> [ResourceGraphQuery](./builder-ResourceGraphQuery.md)
 * <span class="badge builder"></span> [ResourceGroupsQuery](./builder-ResourceGroupsQuery.md)
 * <span class="badge builder"></span> [ResourceNamesQuery](./builder-ResourceNamesQuery.md)
 * <span class="badge builder"></span> [SubscriptionsQuery](./builder-SubscriptionsQuery.md)
 * <span class="badge builder"></span> [TracesFilter](./builder-TracesFilter.md)
 * <span class="badge builder"></span> [TracesQuery](./builder-TracesQuery.md)
 * <span class="badge builder"></span> [UnknownQuery](./builder-UnknownQuery.md)
 * <span class="badge builder"></span> [WorkspacesQuery](./builder-WorkspacesQuery.md)
## Functions

### <span class="badge function"></span> variant_config

variant_config returns the configuration related to grafana-azure-monitor-datasource data queries.

This configuration describes how to unmarshal it, convert it to code, …

```python
def variant_config() -> variants.DataqueryConfig
```

