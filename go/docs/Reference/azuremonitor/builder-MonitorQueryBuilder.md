---
title: <span class="badge builder"></span> MonitorQueryBuilder
---
# <span class="badge builder"></span> MonitorQueryBuilder

## Constructor

```go
func NewMonitorQueryBuilder() *MonitorQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MonitorQueryBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> AzureLogAnalytics

Azure Monitor Logs sub-query properties.

```go
func (builder *MonitorQueryBuilder) AzureLogAnalytics(azureLogAnalytics cog.Builder[azuremonitor.LogsQuery]) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> AzureMonitor

Azure Monitor Metrics sub-query properties.

```go
func (builder *MonitorQueryBuilder) AzureMonitor(azureMonitor cog.Builder[azuremonitor.MetricQuery]) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> AzureResourceGraph

Azure Resource Graph sub-query properties.

```go
func (builder *MonitorQueryBuilder) AzureResourceGraph(azureResourceGraph cog.Builder[azuremonitor.ResourceGraphQuery]) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> AzureTraces

Application Insights Traces sub-query properties.

```go
func (builder *MonitorQueryBuilder) AzureTraces(azureTraces cog.Builder[azuremonitor.TracesQuery]) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> CustomNamespace

Custom namespace used in template variable queries

```go
func (builder *MonitorQueryBuilder) CustomNamespace(customNamespace string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> Datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```go
func (builder *MonitorQueryBuilder) Datasource(datasource any) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> GrafanaTemplateVariableFn

@deprecated Legacy template variable support.

```go
func (builder *MonitorQueryBuilder) GrafanaTemplateVariableFn(grafanaTemplateVariableFn azuremonitor.GrafanaTemplateVariableQuery) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *MonitorQueryBuilder) Hide(hide bool) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> Namespace

Namespace used in template variable queries

```go
func (builder *MonitorQueryBuilder) Namespace(namespace string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> Query

Used only for exemplar queries from Prometheus

```go
func (builder *MonitorQueryBuilder) Query(query string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *MonitorQueryBuilder) QueryType(queryType string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *MonitorQueryBuilder) RefId(refId string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> Region

Region used in template variable queries

```go
func (builder *MonitorQueryBuilder) Region(region string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> Resource

Resource used in template variable queries

```go
func (builder *MonitorQueryBuilder) Resource(resource string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

Resource group used in template variable queries

```go
func (builder *MonitorQueryBuilder) ResourceGroup(resourceGroup string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> Subscription

Azure subscription containing the resource(s) to be queried.

Also used for template variable queries

```go
func (builder *MonitorQueryBuilder) Subscription(subscription string) *MonitorQueryBuilder
```

### <span class="badge object-method"></span> Subscriptions

Subscriptions to be queried via Azure Resource Graph.

```go
func (builder *MonitorQueryBuilder) Subscriptions(subscriptions []string) *MonitorQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MonitorQuery](./object-MonitorQuery.md)
