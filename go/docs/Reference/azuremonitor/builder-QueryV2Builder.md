---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```go
func NewQueryV2Builder() *QueryV2Builder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error)
```

### <span class="badge object-method"></span> AzureLogAnalytics

Azure Monitor Logs sub-query properties.

```go
func (builder *QueryV2Builder) AzureLogAnalytics(azureLogAnalytics cog.Builder[azuremonitor.AzureLogsQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> AzureMonitor

Azure Monitor Metrics sub-query properties.

```go
func (builder *QueryV2Builder) AzureMonitor(azureMonitor cog.Builder[azuremonitor.AzureMetricQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> AzureResourceGraph

Azure Resource Graph sub-query properties.

```go
func (builder *QueryV2Builder) AzureResourceGraph(azureResourceGraph cog.Builder[azuremonitor.AzureResourceGraphQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> AzureTraces

Application Insights Traces sub-query properties.

```go
func (builder *QueryV2Builder) AzureTraces(azureTraces cog.Builder[azuremonitor.AzureTracesQuery]) *QueryV2Builder
```

### <span class="badge object-method"></span> CustomNamespace

Custom namespace used in template variable queries

```go
func (builder *QueryV2Builder) CustomNamespace(customNamespace string) *QueryV2Builder
```

### <span class="badge object-method"></span> Datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```go
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder
```

### <span class="badge object-method"></span> GrafanaTemplateVariableFn

@deprecated Legacy template variable support.

```go
func (builder *QueryV2Builder) GrafanaTemplateVariableFn(grafanaTemplateVariableFn azuremonitor.GrafanaTemplateVariableQuery) *QueryV2Builder
```

### <span class="badge object-method"></span> Hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```go
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder
```

### <span class="badge object-method"></span> Labels

```go
func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder
```

### <span class="badge object-method"></span> Namespace

Namespace used in template variable queries

```go
func (builder *QueryV2Builder) Namespace(namespace string) *QueryV2Builder
```

### <span class="badge object-method"></span> Query

Used only for exemplar queries from Prometheus

```go
func (builder *QueryV2Builder) Query(query string) *QueryV2Builder
```

### <span class="badge object-method"></span> QueryType

Specify the query flavor

TODO make this required and give it a default

```go
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder
```

### <span class="badge object-method"></span> RefId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```go
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder
```

### <span class="badge object-method"></span> Region

Region used in template variable queries

```go
func (builder *QueryV2Builder) Region(region string) *QueryV2Builder
```

### <span class="badge object-method"></span> Resource

Resource used in template variable queries

```go
func (builder *QueryV2Builder) Resource(resource string) *QueryV2Builder
```

### <span class="badge object-method"></span> ResourceGroup

Resource group used in template variable queries

```go
func (builder *QueryV2Builder) ResourceGroup(resourceGroup string) *QueryV2Builder
```

### <span class="badge object-method"></span> Subscription

Azure subscription containing the resource(s) to be queried.

Also used for template variable queries

```go
func (builder *QueryV2Builder) Subscription(subscription string) *QueryV2Builder
```

### <span class="badge object-method"></span> Subscriptions

Subscriptions to be queried via Azure Resource Graph.

```go
func (builder *QueryV2Builder) Subscriptions(subscriptions []string) *QueryV2Builder
```

### <span class="badge object-method"></span> Version

```go
func (builder *QueryV2Builder) Version(version string) *QueryV2Builder
```

## See also

 * <span class="badge object-type-struct"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
