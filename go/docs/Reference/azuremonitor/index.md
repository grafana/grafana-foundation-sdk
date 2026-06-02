# <span class="badge package-variant-dataquery"></span> azuremonitor

## Objects

 * <span class="badge object-type-struct"></span> [AppInsightsGroupByQuery](./object-AppInsightsGroupByQuery.md)
 * <span class="badge object-type-struct"></span> [AppInsightsMetricNameQuery](./object-AppInsightsMetricNameQuery.md)
 * <span class="badge object-type-struct"></span> [AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery](./object-AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery.md)
 * <span class="badge object-type-struct"></span> [BaseGrafanaTemplateVariableQuery](./object-BaseGrafanaTemplateVariableQuery.md)
 * <span class="badge object-type-ref"></span> [GrafanaTemplateVariableQuery](./object-GrafanaTemplateVariableQuery.md)
 * <span class="badge object-type-enum"></span> [GrafanaTemplateVariableQueryType](./object-GrafanaTemplateVariableQueryType.md)
 * <span class="badge object-type-struct"></span> [LogsQuery](./object-LogsQuery.md)
 * <span class="badge object-type-struct"></span> [MetricDefinitionsQuery](./object-MetricDefinitionsQuery.md)
 * <span class="badge object-type-struct"></span> [MetricDimension](./object-MetricDimension.md)
 * <span class="badge object-type-struct"></span> [MetricNamesQuery](./object-MetricNamesQuery.md)
 * <span class="badge object-type-struct"></span> [MetricNamespaceQuery](./object-MetricNamespaceQuery.md)
 * <span class="badge object-type-struct"></span> [MetricQuery](./object-MetricQuery.md)
 * <span class="badge object-type-struct"></span> [MonitorQuery](./object-MonitorQuery.md)
 * <span class="badge object-type-struct"></span> [MonitorResource](./object-MonitorResource.md)
 * <span class="badge object-type-enum"></span> [QueryType](./object-QueryType.md)
 * <span class="badge object-type-struct"></span> [ResourceGraphQuery](./object-ResourceGraphQuery.md)
 * <span class="badge object-type-struct"></span> [ResourceGroupsQuery](./object-ResourceGroupsQuery.md)
 * <span class="badge object-type-struct"></span> [ResourceNamesQuery](./object-ResourceNamesQuery.md)
 * <span class="badge object-type-enum"></span> [ResultFormat](./object-ResultFormat.md)
 * <span class="badge object-type-struct"></span> [SubscriptionsQuery](./object-SubscriptionsQuery.md)
 * <span class="badge object-type-struct"></span> [TracesFilter](./object-TracesFilter.md)
 * <span class="badge object-type-struct"></span> [TracesQuery](./object-TracesQuery.md)
 * <span class="badge object-type-struct"></span> [UnknownQuery](./object-UnknownQuery.md)
 * <span class="badge object-type-struct"></span> [WorkspacesQuery](./object-WorkspacesQuery.md)
## Builders

 * <span class="badge builder"></span> [AppInsightsGroupByQueryBuilder](./builder-AppInsightsGroupByQueryBuilder.md)
 * <span class="badge builder"></span> [AppInsightsMetricNameQueryBuilder](./builder-AppInsightsMetricNameQueryBuilder.md)
 * <span class="badge builder"></span> [BaseGrafanaTemplateVariableQueryBuilder](./builder-BaseGrafanaTemplateVariableQueryBuilder.md)
 * <span class="badge builder"></span> [LogsQueryBuilder](./builder-LogsQueryBuilder.md)
 * <span class="badge builder"></span> [MetricDefinitionsQueryBuilder](./builder-MetricDefinitionsQueryBuilder.md)
 * <span class="badge builder"></span> [MetricDimensionBuilder](./builder-MetricDimensionBuilder.md)
 * <span class="badge builder"></span> [MetricNamesQueryBuilder](./builder-MetricNamesQueryBuilder.md)
 * <span class="badge builder"></span> [MetricNamespaceQueryBuilder](./builder-MetricNamespaceQueryBuilder.md)
 * <span class="badge builder"></span> [MetricQueryBuilder](./builder-MetricQueryBuilder.md)
 * <span class="badge builder"></span> [MonitorQueryBuilder](./builder-MonitorQueryBuilder.md)
 * <span class="badge builder"></span> [MonitorResourceBuilder](./builder-MonitorResourceBuilder.md)
 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)
 * <span class="badge builder"></span> [QueryV2Builder](./builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [ResourceGraphQueryBuilder](./builder-ResourceGraphQueryBuilder.md)
 * <span class="badge builder"></span> [ResourceGroupsQueryBuilder](./builder-ResourceGroupsQueryBuilder.md)
 * <span class="badge builder"></span> [ResourceNamesQueryBuilder](./builder-ResourceNamesQueryBuilder.md)
 * <span class="badge builder"></span> [SubscriptionsQueryBuilder](./builder-SubscriptionsQueryBuilder.md)
 * <span class="badge builder"></span> [TracesFilterBuilder](./builder-TracesFilterBuilder.md)
 * <span class="badge builder"></span> [TracesQueryBuilder](./builder-TracesQueryBuilder.md)
 * <span class="badge builder"></span> [UnknownQueryBuilder](./builder-UnknownQueryBuilder.md)
 * <span class="badge builder"></span> [WorkspacesQueryBuilder](./builder-WorkspacesQueryBuilder.md)
## Functions

### <span class="badge function"></span> NewMonitorQuery

NewMonitorQuery creates a new MonitorQuery object.

```go
func NewMonitorQuery() *MonitorQuery
```

### <span class="badge function"></span> NewMetricQuery

NewMetricQuery creates a new MetricQuery object.

```go
func NewMetricQuery() *MetricQuery
```

### <span class="badge function"></span> NewMonitorResource

NewMonitorResource creates a new MonitorResource object.

```go
func NewMonitorResource() *MonitorResource
```

### <span class="badge function"></span> NewMetricDimension

NewMetricDimension creates a new MetricDimension object.

```go
func NewMetricDimension() *MetricDimension
```

### <span class="badge function"></span> NewLogsQuery

NewLogsQuery creates a new LogsQuery object.

```go
func NewLogsQuery() *LogsQuery
```

### <span class="badge function"></span> NewResourceGraphQuery

NewResourceGraphQuery creates a new ResourceGraphQuery object.

```go
func NewResourceGraphQuery() *ResourceGraphQuery
```

### <span class="badge function"></span> NewTracesQuery

NewTracesQuery creates a new TracesQuery object.

```go
func NewTracesQuery() *TracesQuery
```

### <span class="badge function"></span> NewTracesFilter

NewTracesFilter creates a new TracesFilter object.

```go
func NewTracesFilter() *TracesFilter
```

### <span class="badge function"></span> NewGrafanaTemplateVariableQuery

NewGrafanaTemplateVariableQuery creates a new GrafanaTemplateVariableQuery object.

```go
func NewGrafanaTemplateVariableQuery() *GrafanaTemplateVariableQuery
```

### <span class="badge function"></span> NewAppInsightsMetricNameQuery

NewAppInsightsMetricNameQuery creates a new AppInsightsMetricNameQuery object.

```go
func NewAppInsightsMetricNameQuery() *AppInsightsMetricNameQuery
```

### <span class="badge function"></span> NewAppInsightsGroupByQuery

NewAppInsightsGroupByQuery creates a new AppInsightsGroupByQuery object.

```go
func NewAppInsightsGroupByQuery() *AppInsightsGroupByQuery
```

### <span class="badge function"></span> NewSubscriptionsQuery

NewSubscriptionsQuery creates a new SubscriptionsQuery object.

```go
func NewSubscriptionsQuery() *SubscriptionsQuery
```

### <span class="badge function"></span> NewResourceGroupsQuery

NewResourceGroupsQuery creates a new ResourceGroupsQuery object.

```go
func NewResourceGroupsQuery() *ResourceGroupsQuery
```

### <span class="badge function"></span> NewResourceNamesQuery

NewResourceNamesQuery creates a new ResourceNamesQuery object.

```go
func NewResourceNamesQuery() *ResourceNamesQuery
```

### <span class="badge function"></span> NewMetricNamespaceQuery

NewMetricNamespaceQuery creates a new MetricNamespaceQuery object.

```go
func NewMetricNamespaceQuery() *MetricNamespaceQuery
```

### <span class="badge function"></span> NewMetricDefinitionsQuery

NewMetricDefinitionsQuery creates a new MetricDefinitionsQuery object.

```go
func NewMetricDefinitionsQuery() *MetricDefinitionsQuery
```

### <span class="badge function"></span> NewMetricNamesQuery

NewMetricNamesQuery creates a new MetricNamesQuery object.

```go
func NewMetricNamesQuery() *MetricNamesQuery
```

### <span class="badge function"></span> NewWorkspacesQuery

NewWorkspacesQuery creates a new WorkspacesQuery object.

```go
func NewWorkspacesQuery() *WorkspacesQuery
```

### <span class="badge function"></span> NewUnknownQuery

NewUnknownQuery creates a new UnknownQuery object.

```go
func NewUnknownQuery() *UnknownQuery
```

### <span class="badge function"></span> NewBaseGrafanaTemplateVariableQuery

NewBaseGrafanaTemplateVariableQuery creates a new BaseGrafanaTemplateVariableQuery object.

```go
func NewBaseGrafanaTemplateVariableQuery() *BaseGrafanaTemplateVariableQuery
```

### <span class="badge function"></span> NewAppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery

NewAppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery creates a new AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery object.

```go
func NewAppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery() *AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to grafana-azure-monitor-datasource dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> QueryV2Converter

QueryV2Converter accepts a `QueryV2` object and generates the Go code to build this object using builders.

```go
func QueryV2Converter(input dashboardv2.DataQueryKind) string
```

### <span class="badge function"></span> MonitorQueryConverter

MonitorQueryConverter accepts a `MonitorQuery` object and generates the Go code to build this object using builders.

```go
func MonitorQueryConverter(input MonitorQuery) string
```

### <span class="badge function"></span> MetricQueryConverter

MetricQueryConverter accepts a `MetricQuery` object and generates the Go code to build this object using builders.

```go
func MetricQueryConverter(input MetricQuery) string
```

### <span class="badge function"></span> MonitorResourceConverter

MonitorResourceConverter accepts a `MonitorResource` object and generates the Go code to build this object using builders.

```go
func MonitorResourceConverter(input MonitorResource) string
```

### <span class="badge function"></span> MetricDimensionConverter

MetricDimensionConverter accepts a `MetricDimension` object and generates the Go code to build this object using builders.

```go
func MetricDimensionConverter(input MetricDimension) string
```

### <span class="badge function"></span> LogsQueryConverter

LogsQueryConverter accepts a `LogsQuery` object and generates the Go code to build this object using builders.

```go
func LogsQueryConverter(input LogsQuery) string
```

### <span class="badge function"></span> ResourceGraphQueryConverter

ResourceGraphQueryConverter accepts a `ResourceGraphQuery` object and generates the Go code to build this object using builders.

```go
func ResourceGraphQueryConverter(input ResourceGraphQuery) string
```

### <span class="badge function"></span> TracesQueryConverter

TracesQueryConverter accepts a `TracesQuery` object and generates the Go code to build this object using builders.

```go
func TracesQueryConverter(input TracesQuery) string
```

### <span class="badge function"></span> TracesFilterConverter

TracesFilterConverter accepts a `TracesFilter` object and generates the Go code to build this object using builders.

```go
func TracesFilterConverter(input TracesFilter) string
```

### <span class="badge function"></span> AppInsightsMetricNameQueryConverter

AppInsightsMetricNameQueryConverter accepts a `AppInsightsMetricNameQuery` object and generates the Go code to build this object using builders.

```go
func AppInsightsMetricNameQueryConverter(input AppInsightsMetricNameQuery) string
```

### <span class="badge function"></span> AppInsightsGroupByQueryConverter

AppInsightsGroupByQueryConverter accepts a `AppInsightsGroupByQuery` object and generates the Go code to build this object using builders.

```go
func AppInsightsGroupByQueryConverter(input AppInsightsGroupByQuery) string
```

### <span class="badge function"></span> SubscriptionsQueryConverter

SubscriptionsQueryConverter accepts a `SubscriptionsQuery` object and generates the Go code to build this object using builders.

```go
func SubscriptionsQueryConverter(input SubscriptionsQuery) string
```

### <span class="badge function"></span> ResourceGroupsQueryConverter

ResourceGroupsQueryConverter accepts a `ResourceGroupsQuery` object and generates the Go code to build this object using builders.

```go
func ResourceGroupsQueryConverter(input ResourceGroupsQuery) string
```

### <span class="badge function"></span> ResourceNamesQueryConverter

ResourceNamesQueryConverter accepts a `ResourceNamesQuery` object and generates the Go code to build this object using builders.

```go
func ResourceNamesQueryConverter(input ResourceNamesQuery) string
```

### <span class="badge function"></span> MetricNamespaceQueryConverter

MetricNamespaceQueryConverter accepts a `MetricNamespaceQuery` object and generates the Go code to build this object using builders.

```go
func MetricNamespaceQueryConverter(input MetricNamespaceQuery) string
```

### <span class="badge function"></span> MetricDefinitionsQueryConverter

MetricDefinitionsQueryConverter accepts a `MetricDefinitionsQuery` object and generates the Go code to build this object using builders.

```go
func MetricDefinitionsQueryConverter(input MetricDefinitionsQuery) string
```

### <span class="badge function"></span> MetricNamesQueryConverter

MetricNamesQueryConverter accepts a `MetricNamesQuery` object and generates the Go code to build this object using builders.

```go
func MetricNamesQueryConverter(input MetricNamesQuery) string
```

### <span class="badge function"></span> WorkspacesQueryConverter

WorkspacesQueryConverter accepts a `WorkspacesQuery` object and generates the Go code to build this object using builders.

```go
func WorkspacesQueryConverter(input WorkspacesQuery) string
```

### <span class="badge function"></span> UnknownQueryConverter

UnknownQueryConverter accepts a `UnknownQuery` object and generates the Go code to build this object using builders.

```go
func UnknownQueryConverter(input UnknownQuery) string
```

### <span class="badge function"></span> BaseGrafanaTemplateVariableQueryConverter

BaseGrafanaTemplateVariableQueryConverter accepts a `BaseGrafanaTemplateVariableQuery` object and generates the Go code to build this object using builders.

```go
func BaseGrafanaTemplateVariableQueryConverter(input BaseGrafanaTemplateVariableQuery) string
```

### <span class="badge function"></span> QueryConverter

QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.

```go
func QueryConverter(input dashboardv2beta1.DataQueryKind) string
```

