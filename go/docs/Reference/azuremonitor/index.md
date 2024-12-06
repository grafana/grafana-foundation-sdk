# azuremonitor

## Objects

 * <span class="badge object-type-struct"></span> [AppInsightsGroupByQuery](./object-AppInsightsGroupByQuery.md)
 * <span class="badge object-type-struct"></span> [AppInsightsMetricNameQuery](./object-AppInsightsMetricNameQuery.md)
 * <span class="badge object-type-struct"></span> [AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery](./object-AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery.md)
 * <span class="badge object-type-struct"></span> [AzureLogsQuery](./object-AzureLogsQuery.md)
 * <span class="badge object-type-struct"></span> [AzureMetricDimension](./object-AzureMetricDimension.md)
 * <span class="badge object-type-struct"></span> [AzureMetricQuery](./object-AzureMetricQuery.md)
 * <span class="badge object-type-struct"></span> [AzureMonitorQuery](./object-AzureMonitorQuery.md)
 * <span class="badge object-type-struct"></span> [AzureMonitorResource](./object-AzureMonitorResource.md)
 * <span class="badge object-type-enum"></span> [AzureQueryType](./object-AzureQueryType.md)
 * <span class="badge object-type-struct"></span> [AzureResourceGraphQuery](./object-AzureResourceGraphQuery.md)
 * <span class="badge object-type-struct"></span> [AzureTracesFilter](./object-AzureTracesFilter.md)
 * <span class="badge object-type-struct"></span> [AzureTracesQuery](./object-AzureTracesQuery.md)
 * <span class="badge object-type-struct"></span> [BaseGrafanaTemplateVariableQuery](./object-BaseGrafanaTemplateVariableQuery.md)
 * <span class="badge object-type-ref"></span> [GrafanaTemplateVariableQuery](./object-GrafanaTemplateVariableQuery.md)
 * <span class="badge object-type-enum"></span> [GrafanaTemplateVariableQueryType](./object-GrafanaTemplateVariableQueryType.md)
 * <span class="badge object-type-struct"></span> [MetricDefinitionsQuery](./object-MetricDefinitionsQuery.md)
 * <span class="badge object-type-struct"></span> [MetricNamesQuery](./object-MetricNamesQuery.md)
 * <span class="badge object-type-struct"></span> [MetricNamespaceQuery](./object-MetricNamespaceQuery.md)
 * <span class="badge object-type-struct"></span> [ResourceGroupsQuery](./object-ResourceGroupsQuery.md)
 * <span class="badge object-type-struct"></span> [ResourceNamesQuery](./object-ResourceNamesQuery.md)
 * <span class="badge object-type-enum"></span> [ResultFormat](./object-ResultFormat.md)
 * <span class="badge object-type-struct"></span> [SubscriptionsQuery](./object-SubscriptionsQuery.md)
 * <span class="badge object-type-struct"></span> [UnknownQuery](./object-UnknownQuery.md)
 * <span class="badge object-type-struct"></span> [WorkspacesQuery](./object-WorkspacesQuery.md)
## Builders

 * <span class="badge builder"></span> [AppInsightsGroupByQueryBuilder](./builder-AppInsightsGroupByQueryBuilder.md)
 * <span class="badge builder"></span> [AppInsightsMetricNameQueryBuilder](./builder-AppInsightsMetricNameQueryBuilder.md)
 * <span class="badge builder"></span> [AzureLogsQueryBuilder](./builder-AzureLogsQueryBuilder.md)
 * <span class="badge builder"></span> [AzureMetricDimensionBuilder](./builder-AzureMetricDimensionBuilder.md)
 * <span class="badge builder"></span> [AzureMetricQueryBuilder](./builder-AzureMetricQueryBuilder.md)
 * <span class="badge builder"></span> [AzureMonitorQueryBuilder](./builder-AzureMonitorQueryBuilder.md)
 * <span class="badge builder"></span> [AzureMonitorResourceBuilder](./builder-AzureMonitorResourceBuilder.md)
 * <span class="badge builder"></span> [AzureResourceGraphQueryBuilder](./builder-AzureResourceGraphQueryBuilder.md)
 * <span class="badge builder"></span> [AzureTracesFilterBuilder](./builder-AzureTracesFilterBuilder.md)
 * <span class="badge builder"></span> [AzureTracesQueryBuilder](./builder-AzureTracesQueryBuilder.md)
 * <span class="badge builder"></span> [BaseGrafanaTemplateVariableQueryBuilder](./builder-BaseGrafanaTemplateVariableQueryBuilder.md)
 * <span class="badge builder"></span> [MetricDefinitionsQueryBuilder](./builder-MetricDefinitionsQueryBuilder.md)
 * <span class="badge builder"></span> [MetricNamesQueryBuilder](./builder-MetricNamesQueryBuilder.md)
 * <span class="badge builder"></span> [MetricNamespaceQueryBuilder](./builder-MetricNamespaceQueryBuilder.md)
 * <span class="badge builder"></span> [ResourceGroupsQueryBuilder](./builder-ResourceGroupsQueryBuilder.md)
 * <span class="badge builder"></span> [ResourceNamesQueryBuilder](./builder-ResourceNamesQueryBuilder.md)
 * <span class="badge builder"></span> [SubscriptionsQueryBuilder](./builder-SubscriptionsQueryBuilder.md)
 * <span class="badge builder"></span> [UnknownQueryBuilder](./builder-UnknownQueryBuilder.md)
 * <span class="badge builder"></span> [WorkspacesQueryBuilder](./builder-WorkspacesQueryBuilder.md)
## Functions

### <span class="badge function"></span> NewAzureMonitorQuery

NewAzureMonitorQuery creates a new AzureMonitorQuery object.

```go
func NewAzureMonitorQuery() *AzureMonitorQuery
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to grafana-azure-monitor-datasource dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> NewAzureMetricQuery

NewAzureMetricQuery creates a new AzureMetricQuery object.

```go
func NewAzureMetricQuery() *AzureMetricQuery
```

### <span class="badge function"></span> NewAzureLogsQuery

NewAzureLogsQuery creates a new AzureLogsQuery object.

```go
func NewAzureLogsQuery() *AzureLogsQuery
```

### <span class="badge function"></span> NewAzureTracesQuery

NewAzureTracesQuery creates a new AzureTracesQuery object.

```go
func NewAzureTracesQuery() *AzureTracesQuery
```

### <span class="badge function"></span> NewAzureTracesFilter

NewAzureTracesFilter creates a new AzureTracesFilter object.

```go
func NewAzureTracesFilter() *AzureTracesFilter
```

### <span class="badge function"></span> NewAzureResourceGraphQuery

NewAzureResourceGraphQuery creates a new AzureResourceGraphQuery object.

```go
func NewAzureResourceGraphQuery() *AzureResourceGraphQuery
```

### <span class="badge function"></span> NewAzureMonitorResource

NewAzureMonitorResource creates a new AzureMonitorResource object.

```go
func NewAzureMonitorResource() *AzureMonitorResource
```

### <span class="badge function"></span> NewAzureMetricDimension

NewAzureMetricDimension creates a new AzureMetricDimension object.

```go
func NewAzureMetricDimension() *AzureMetricDimension
```

### <span class="badge function"></span> NewBaseGrafanaTemplateVariableQuery

NewBaseGrafanaTemplateVariableQuery creates a new BaseGrafanaTemplateVariableQuery object.

```go
func NewBaseGrafanaTemplateVariableQuery() *BaseGrafanaTemplateVariableQuery
```

### <span class="badge function"></span> NewUnknownQuery

NewUnknownQuery creates a new UnknownQuery object.

```go
func NewUnknownQuery() *UnknownQuery
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

### <span class="badge function"></span> NewGrafanaTemplateVariableQuery

NewGrafanaTemplateVariableQuery creates a new GrafanaTemplateVariableQuery object.

```go
func NewGrafanaTemplateVariableQuery() *GrafanaTemplateVariableQuery
```

### <span class="badge function"></span> NewAppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery

NewAppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery creates a new AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery object.

```go
func NewAppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery() *AppInsightsMetricNameQueryOrAppInsightsGroupByQueryOrSubscriptionsQueryOrResourceGroupsQueryOrResourceNamesQueryOrMetricNamespaceQueryOrMetricDefinitionsQueryOrMetricNamesQueryOrWorkspacesQueryOrUnknownQuery
```

### <span class="badge function"></span> AzureMonitorQueryConverter

AzureMonitorQueryConverter accepts a `AzureMonitorQuery` object and generates the Go code to build this object using builders.

```go
func AzureMonitorQueryConverter(input AzureMonitorQuery) string
```

### <span class="badge function"></span> AzureMetricQueryConverter

AzureMetricQueryConverter accepts a `AzureMetricQuery` object and generates the Go code to build this object using builders.

```go
func AzureMetricQueryConverter(input AzureMetricQuery) string
```

### <span class="badge function"></span> AzureLogsQueryConverter

AzureLogsQueryConverter accepts a `AzureLogsQuery` object and generates the Go code to build this object using builders.

```go
func AzureLogsQueryConverter(input AzureLogsQuery) string
```

### <span class="badge function"></span> AzureTracesQueryConverter

AzureTracesQueryConverter accepts a `AzureTracesQuery` object and generates the Go code to build this object using builders.

```go
func AzureTracesQueryConverter(input AzureTracesQuery) string
```

### <span class="badge function"></span> AzureTracesFilterConverter

AzureTracesFilterConverter accepts a `AzureTracesFilter` object and generates the Go code to build this object using builders.

```go
func AzureTracesFilterConverter(input AzureTracesFilter) string
```

### <span class="badge function"></span> AzureResourceGraphQueryConverter

AzureResourceGraphQueryConverter accepts a `AzureResourceGraphQuery` object and generates the Go code to build this object using builders.

```go
func AzureResourceGraphQueryConverter(input AzureResourceGraphQuery) string
```

### <span class="badge function"></span> AzureMonitorResourceConverter

AzureMonitorResourceConverter accepts a `AzureMonitorResource` object and generates the Go code to build this object using builders.

```go
func AzureMonitorResourceConverter(input AzureMonitorResource) string
```

### <span class="badge function"></span> AzureMetricDimensionConverter

AzureMetricDimensionConverter accepts a `AzureMetricDimension` object and generates the Go code to build this object using builders.

```go
func AzureMetricDimensionConverter(input AzureMetricDimension) string
```

### <span class="badge function"></span> BaseGrafanaTemplateVariableQueryConverter

BaseGrafanaTemplateVariableQueryConverter accepts a `BaseGrafanaTemplateVariableQuery` object and generates the Go code to build this object using builders.

```go
func BaseGrafanaTemplateVariableQueryConverter(input BaseGrafanaTemplateVariableQuery) string
```

### <span class="badge function"></span> UnknownQueryConverter

UnknownQueryConverter accepts a `UnknownQuery` object and generates the Go code to build this object using builders.

```go
func UnknownQueryConverter(input UnknownQuery) string
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

