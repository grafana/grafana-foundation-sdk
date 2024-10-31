---
title: <span class="badge builder"></span> AzureMonitorQueryBuilder
---
# <span class="badge builder"></span> AzureMonitorQueryBuilder

## Constructor

```typescript
new AzureMonitorQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> azureLogAnalytics

Azure Monitor Logs sub-query properties.

```typescript
azureLogAnalytics(azureLogAnalytics: cog.Builder<azuremonitor.AzureLogsQuery>)
```

### <span class="badge object-method"></span> azureMonitor

Azure Monitor Metrics sub-query properties.

```typescript
azureMonitor(azureMonitor: cog.Builder<azuremonitor.AzureMetricQuery>)
```

### <span class="badge object-method"></span> azureResourceGraph

Azure Resource Graph sub-query properties.

```typescript
azureResourceGraph(azureResourceGraph: cog.Builder<azuremonitor.AzureResourceGraphQuery>)
```

### <span class="badge object-method"></span> azureTraces

Application Insights Traces sub-query properties.

```typescript
azureTraces(azureTraces: cog.Builder<azuremonitor.AzureTracesQuery>)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```typescript
datasource(datasource: dashboard.DataSourceRef)
```

### <span class="badge object-method"></span> grafanaTemplateVariableFn

@deprecated Legacy template variable support.

```typescript
grafanaTemplateVariableFn(grafanaTemplateVariableFn: azuremonitor.GrafanaTemplateVariableQuery)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> namespace

```typescript
namespace(namespace: string)
```

### <span class="badge object-method"></span> query

Used only for exemplar queries from Prometheus

```typescript
query(query: string)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```typescript
refId(refId: string)
```

### <span class="badge object-method"></span> region

```typescript
region(region: string)
```

### <span class="badge object-method"></span> resource

```typescript
resource(resource: string)
```

### <span class="badge object-method"></span> resourceGroup

Template variables params. These exist for backwards compatiblity with legacy template variables.

```typescript
resourceGroup(resourceGroup: string)
```

### <span class="badge object-method"></span> subscription

Azure subscription containing the resource(s) to be queried.

```typescript
subscription(subscription: string)
```

### <span class="badge object-method"></span> subscriptions

Subscriptions to be queried via Azure Resource Graph.

```typescript
subscriptions(subscriptions: string[])
```

## See also

 * <span class="badge object-type-interface"></span> [AzureMonitorQuery](./object-AzureMonitorQuery.md)
