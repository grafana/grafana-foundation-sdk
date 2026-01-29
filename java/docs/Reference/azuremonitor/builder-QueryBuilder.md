---
title: <span class="badge builder"></span> QueryBuilder
---
# <span class="badge builder"></span> QueryBuilder

## Constructor

```java
new QueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Query build()
```

### <span class="badge object-method"></span> azureLogAnalytics

Azure Monitor Logs sub-query properties.

```java
public QueryBuilder azureLogAnalytics(com.grafana.foundation.cog.Builder<AzureLogsQuery> azureLogAnalytics)
```

### <span class="badge object-method"></span> azureMonitor

Azure Monitor Metrics sub-query properties.

```java
public QueryBuilder azureMonitor(com.grafana.foundation.cog.Builder<AzureMetricQuery> azureMonitor)
```

### <span class="badge object-method"></span> azureResourceGraph

Azure Resource Graph sub-query properties.

```java
public QueryBuilder azureResourceGraph(com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> azureResourceGraph)
```

### <span class="badge object-method"></span> azureTraces

Application Insights Traces sub-query properties.

```java
public QueryBuilder azureTraces(com.grafana.foundation.cog.Builder<AzureTracesQuery> azureTraces)
```

### <span class="badge object-method"></span> customNamespace

Custom namespace used in template variable queries

```java
public QueryBuilder customNamespace(String customNamespace)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> grafanaTemplateVariableFn

@deprecated Legacy template variable support.

```java
public QueryBuilder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> namespace

Namespace used in template variable queries

```java
public QueryBuilder namespace(String namespace)
```

### <span class="badge object-method"></span> oldDatasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public QueryBuilder oldDatasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> query

Used only for exemplar queries from Prometheus

```java
public QueryBuilder query(String query)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> region

Region used in template variable queries

```java
public QueryBuilder region(String region)
```

### <span class="badge object-method"></span> resource

Resource used in template variable queries

```java
public QueryBuilder resource(String resource)
```

### <span class="badge object-method"></span> resourceGroup

Resource group used in template variable queries

```java
public QueryBuilder resourceGroup(String resourceGroup)
```

### <span class="badge object-method"></span> subscription

Azure subscription containing the resource(s) to be queried.

Also used for template variable queries

```java
public QueryBuilder subscription(String subscription)
```

### <span class="badge object-method"></span> subscriptions

Subscriptions to be queried via Azure Resource Graph.

```java
public QueryBuilder subscriptions(List<String> subscriptions)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
