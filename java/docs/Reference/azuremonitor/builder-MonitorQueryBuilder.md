---
title: <span class="badge builder"></span> MonitorQueryBuilder
---
# <span class="badge builder"></span> MonitorQueryBuilder

## Constructor

```java
new MonitorQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public MonitorQuery build()
```

### <span class="badge object-method"></span> azureLogAnalytics

Azure Monitor Logs sub-query properties.

```java
public MonitorQueryBuilder azureLogAnalytics(com.grafana.foundation.cog.Builder<LogsQuery> azureLogAnalytics)
```

### <span class="badge object-method"></span> azureMonitor

Azure Monitor Metrics sub-query properties.

```java
public MonitorQueryBuilder azureMonitor(com.grafana.foundation.cog.Builder<MetricQuery> azureMonitor)
```

### <span class="badge object-method"></span> azureResourceGraph

Azure Resource Graph sub-query properties.

```java
public MonitorQueryBuilder azureResourceGraph(com.grafana.foundation.cog.Builder<ResourceGraphQuery> azureResourceGraph)
```

### <span class="badge object-method"></span> azureTraces

Application Insights Traces sub-query properties.

```java
public MonitorQueryBuilder azureTraces(com.grafana.foundation.cog.Builder<TracesQuery> azureTraces)
```

### <span class="badge object-method"></span> customNamespace

Custom namespace used in template variable queries

```java
public MonitorQueryBuilder customNamespace(String customNamespace)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public MonitorQueryBuilder datasource(Object datasource)
```

### <span class="badge object-method"></span> grafanaTemplateVariableFn

@deprecated Legacy template variable support.

```java
public MonitorQueryBuilder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public MonitorQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> namespace

Namespace used in template variable queries

```java
public MonitorQueryBuilder namespace(String namespace)
```

### <span class="badge object-method"></span> query

Used only for exemplar queries from Prometheus

```java
public MonitorQueryBuilder query(String query)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public MonitorQueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public MonitorQueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> region

Region used in template variable queries

```java
public MonitorQueryBuilder region(String region)
```

### <span class="badge object-method"></span> resource

Resource used in template variable queries

```java
public MonitorQueryBuilder resource(String resource)
```

### <span class="badge object-method"></span> resourceGroup

Resource group used in template variable queries

```java
public MonitorQueryBuilder resourceGroup(String resourceGroup)
```

### <span class="badge object-method"></span> subscription

Azure subscription containing the resource(s) to be queried.

Also used for template variable queries

```java
public MonitorQueryBuilder subscription(String subscription)
```

### <span class="badge object-method"></span> subscriptions

Subscriptions to be queried via Azure Resource Graph.

```java
public MonitorQueryBuilder subscriptions(List<String> subscriptions)
```

## See also

 * <span class="badge object-type-class"></span> [MonitorQuery](./object-MonitorQuery.md)
