---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```go
func NewQueryBuilder() *QueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error)
```

### <span class="badge object-method"></span> AzureLogAnalytics

Azure Monitor Logs sub-query properties.

```go
func (builder *QueryBuilder) AzureLogAnalytics(azureLogAnalytics cog.Builder[azuremonitor.AzureLogsQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> AzureMonitor

Azure Monitor Metrics sub-query properties.

```go
func (builder *QueryBuilder) AzureMonitor(azureMonitor cog.Builder[azuremonitor.AzureMetricQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> AzureResourceGraph

Azure Resource Graph sub-query properties.

```go
func (builder *QueryBuilder) AzureResourceGraph(azureResourceGraph cog.Builder[azuremonitor.AzureResourceGraphQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> AzureTraces

Application Insights Traces sub-query properties.

```go
func (builder *QueryBuilder) AzureTraces(azureTraces cog.Builder[azuremonitor.AzureTracesQuery]) *QueryBuilder
```

### <span class="badge object-method"></span> CustomNamespace

Custom namespace used in template variable queries

```go
func (builder *QueryBuilder) CustomNamespace(customNamespace string) *QueryBuilder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder
```

### <span class="badge object-method"></span> GrafanaTemplateVariableFn

@deprecated Legacy template variable support.

```go
func (builder *QueryBuilder) GrafanaTemplateVariableFn(grafanaTemplateVariableFn azuremonitor.GrafanaTemplateVariableQuery) *QueryBuilder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder
```

### <span class="badge object-method"></span> Namespace

Namespace used in template variable queries

```go
func (builder *QueryBuilder) Namespace(namespace string) *QueryBuilder
```

### <span class="badge object-method"></span> Query

Used only for exemplar queries from Prometheus

```go
func (builder *QueryBuilder) Query(query string) *QueryBuilder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder
```

### <span class="badge object-method"></span> Region

Region used in template variable queries

```go
func (builder *QueryBuilder) Region(region string) *QueryBuilder
```

### <span class="badge object-method"></span> Resource

Resource used in template variable queries

```go
func (builder *QueryBuilder) Resource(resource string) *QueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

Resource group used in template variable queries

```go
func (builder *QueryBuilder) ResourceGroup(resourceGroup string) *QueryBuilder
```

### <span class="badge object-method"></span> Subscription

Azure subscription containing the resource(s) to be queried.

Also used for template variable queries

```go
func (builder *QueryBuilder) Subscription(subscription string) *QueryBuilder
```

### <span class="badge object-method"></span> Subscriptions

Subscriptions to be queried via Azure Resource Graph.

```go
func (builder *QueryBuilder) Subscriptions(subscriptions []string) *QueryBuilder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryBuilder) Version(version string) *QueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
