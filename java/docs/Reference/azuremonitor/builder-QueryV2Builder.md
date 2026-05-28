---
title: <span class="badge builder"></span> QueryV2Builder
---
# <span class="badge builder"></span> QueryV2Builder

## Constructor

```java
new QueryV2Builder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public QueryV2 build()
```

### <span class="badge object-method"></span> azureLogAnalytics

Azure Monitor Logs sub-query properties.

```java
public QueryV2Builder azureLogAnalytics(com.grafana.foundation.cog.Builder<AzureLogsQuery> azureLogAnalytics)
```

### <span class="badge object-method"></span> azureMonitor

Azure Monitor Metrics sub-query properties.

```java
public QueryV2Builder azureMonitor(com.grafana.foundation.cog.Builder<AzureMetricQuery> azureMonitor)
```

### <span class="badge object-method"></span> azureResourceGraph

Azure Resource Graph sub-query properties.

```java
public QueryV2Builder azureResourceGraph(com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> azureResourceGraph)
```

### <span class="badge object-method"></span> azureTraces

Application Insights Traces sub-query properties.

```java
public QueryV2Builder azureTraces(com.grafana.foundation.cog.Builder<AzureTracesQuery> azureTraces)
```

### <span class="badge object-method"></span> customNamespace

Custom namespace used in template variable queries

```java
public QueryV2Builder customNamespace(String customNamespace)
```

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> grafanaTemplateVariableFn

@deprecated Legacy template variable support.

```java
public QueryV2Builder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryV2Builder hide(Boolean hide)
```

### <span class="badge object-method"></span> labels

```java
public QueryV2Builder labels(Map<String, String> labels)
```

### <span class="badge object-method"></span> namespace

Namespace used in template variable queries

```java
public QueryV2Builder namespace(String namespace)
```

### <span class="badge object-method"></span> query

Used only for exemplar queries from Prometheus

```java
public QueryV2Builder query(String query)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryV2Builder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryV2Builder refId(String refId)
```

### <span class="badge object-method"></span> region

Region used in template variable queries

```java
public QueryV2Builder region(String region)
```

### <span class="badge object-method"></span> resource

Resource used in template variable queries

```java
public QueryV2Builder resource(String resource)
```

### <span class="badge object-method"></span> resourceGroup

Resource group used in template variable queries

```java
public QueryV2Builder resourceGroup(String resourceGroup)
```

### <span class="badge object-method"></span> subscription

Azure subscription containing the resource(s) to be queried.

Also used for template variable queries

```java
public QueryV2Builder subscription(String subscription)
```

### <span class="badge object-method"></span> subscriptions

Subscriptions to be queried via Azure Resource Graph.

```java
public QueryV2Builder subscriptions(List<String> subscriptions)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
