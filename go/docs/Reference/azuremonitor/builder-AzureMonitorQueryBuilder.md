---
title: <span class="badge builder"></span> AzureMonitorQueryBuilder
---
# <span class="badge builder"></span> AzureMonitorQueryBuilder

## Constructor

```go
func NewAzureMonitorQueryBuilder() *AzureMonitorQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AzureMonitorQueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> AzureLogAnalytics

Azure Monitor Logs sub-query properties.

```go
func (builder *AzureMonitorQueryBuilder) AzureLogAnalytics(azureLogAnalytics cog.Builder[azuremonitor.AzureLogsQuery]) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> AzureMonitor

Azure Monitor Metrics sub-query properties.

```go
func (builder *AzureMonitorQueryBuilder) AzureMonitor(azureMonitor cog.Builder[azuremonitor.AzureMetricQuery]) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> AzureResourceGraph

Azure Resource Graph sub-query properties.

```go
func (builder *AzureMonitorQueryBuilder) AzureResourceGraph(azureResourceGraph cog.Builder[azuremonitor.AzureResourceGraphQuery]) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> AzureTraces

Application Insights Traces sub-query properties.

```go
func (builder *AzureMonitorQueryBuilder) AzureTraces(azureTraces cog.Builder[azuremonitor.AzureTracesQuery]) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *AzureMonitorQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> GrafanaTemplateVariableFn

@deprecated Legacy template variable support.

```go
func (builder *AzureMonitorQueryBuilder) GrafanaTemplateVariableFn(grafanaTemplateVariableFn azuremonitor.GrafanaTemplateVariableQuery) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> Hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```go
func (builder *AzureMonitorQueryBuilder) Hide(hide bool) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> Namespace

```go
func (builder *AzureMonitorQueryBuilder) Namespace(namespace string) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *AzureMonitorQueryBuilder) QueryType(queryType string) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *AzureMonitorQueryBuilder) RefId(refId string) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> Region

Azure Monitor query type.

queryType: #AzureQueryType

```go
func (builder *AzureMonitorQueryBuilder) Region(region string) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> Resource

```go
func (builder *AzureMonitorQueryBuilder) Resource(resource string) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

Template variables params. These exist for backwards compatiblity with legacy template variables.

```go
func (builder *AzureMonitorQueryBuilder) ResourceGroup(resourceGroup string) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> Subscription

Azure subscription containing the resource(s) to be queried.

```go
func (builder *AzureMonitorQueryBuilder) Subscription(subscription string) *AzureMonitorQueryBuilder
```

### <span class="badge object-method"></span> Subscriptions

Subscriptions to be queried via Azure Resource Graph.

```go
func (builder *AzureMonitorQueryBuilder) Subscriptions(subscriptions []string) *AzureMonitorQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AzureMonitorQuery](./object-AzureMonitorQuery.md)
