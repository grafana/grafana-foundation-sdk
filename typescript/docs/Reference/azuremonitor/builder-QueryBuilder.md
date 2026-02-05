---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```typescript
new QueryBuilder()
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

### <span class="badge object-method"></span> customNamespace

Custom namespace used in template variable queries

```typescript
customNamespace(customNamespace: string)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```typescript
datasource(ref: {
	name?: string;
})
```

### <span class="badge object-method"></span> grafanaTemplateVariableFn

@deprecated Legacy template variable support.

```typescript
grafanaTemplateVariableFn(grafanaTemplateVariableFn: cog.Builder<azuremonitor.GrafanaTemplateVariableQuery>)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> namespace

Namespace used in template variable queries

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

Region used in template variable queries

```typescript
region(region: string)
```

### <span class="badge object-method"></span> resource

Resource used in template variable queries

```typescript
resource(resource: string)
```

### <span class="badge object-method"></span> resourceGroup

Resource group used in template variable queries

```typescript
resourceGroup(resourceGroup: string)
```

### <span class="badge object-method"></span> subscription

Azure subscription containing the resource(s) to be queried.

Also used for template variable queries

```typescript
subscription(subscription: string)
```

### <span class="badge object-method"></span> subscriptions

Subscriptions to be queried via Azure Resource Graph.

```typescript
subscriptions(subscriptions: string[])
```

### <span class="badge object-method"></span> version

```typescript
version(version: string)
```

## See also

 * <span class="badge object-type-interface"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
