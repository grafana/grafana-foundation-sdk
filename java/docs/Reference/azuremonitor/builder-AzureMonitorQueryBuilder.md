---
title: <span class="badge builder"></span> AzureMonitorQueryBuilder
---
# <span class="badge builder"></span> AzureMonitorQueryBuilder

## Constructor

```java
new AzureMonitorQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AzureMonitorQuery build()
```

### <span class="badge object-method"></span> azureLogAnalytics

Azure Monitor Logs sub-query properties.

```java
public AzureMonitorQueryBuilder azureLogAnalytics(com.grafana.foundation.cog.Builder<AzureLogsQuery> azureLogAnalytics)
```

### <span class="badge object-method"></span> azureMonitor

Azure Monitor Metrics sub-query properties.

```java
public AzureMonitorQueryBuilder azureMonitor(com.grafana.foundation.cog.Builder<AzureMetricQuery> azureMonitor)
```

### <span class="badge object-method"></span> azureResourceGraph

Azure Resource Graph sub-query properties.

```java
public AzureMonitorQueryBuilder azureResourceGraph(com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> azureResourceGraph)
```

### <span class="badge object-method"></span> azureTraces

Application Insights Traces sub-query properties.

```java
public AzureMonitorQueryBuilder azureTraces(com.grafana.foundation.cog.Builder<AzureTracesQuery> azureTraces)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public AzureMonitorQueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> grafanaTemplateVariableFn

@deprecated Legacy template variable support.

```java
public AzureMonitorQueryBuilder grafanaTemplateVariableFn(GrafanaTemplateVariableQuery grafanaTemplateVariableFn)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public AzureMonitorQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> namespace

```java
public AzureMonitorQueryBuilder namespace(String namespace)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public AzureMonitorQueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public AzureMonitorQueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> region

Azure Monitor query type.

queryType: #AzureQueryType

```java
public AzureMonitorQueryBuilder region(String region)
```

### <span class="badge object-method"></span> resource

```java
public AzureMonitorQueryBuilder resource(String resource)
```

### <span class="badge object-method"></span> resourceGroup

Template variables params. These exist for backwards compatiblity with legacy template variables.

```java
public AzureMonitorQueryBuilder resourceGroup(String resourceGroup)
```

### <span class="badge object-method"></span> subscription

Azure subscription containing the resource(s) to be queried.

```java
public AzureMonitorQueryBuilder subscription(String subscription)
```

### <span class="badge object-method"></span> subscriptions

Subscriptions to be queried via Azure Resource Graph.

```java
public AzureMonitorQueryBuilder subscriptions(List<String> subscriptions)
```

## See also

 * <span class="badge object-type-class"></span> [AzureMonitorQuery](./object-AzureMonitorQuery.md)
