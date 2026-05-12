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

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> exemplars

For metric queries, how many exemplars to request, 0 means no exemplars

```java
public QueryV2Builder exemplars(Long exemplars)
```

### <span class="badge object-method"></span> filters

```java
public QueryV2Builder filters(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> filters)
```

### <span class="badge object-method"></span> groupBy

Filters that are used to query the metrics summary

```java
public QueryV2Builder groupBy(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> groupBy)
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

### <span class="badge object-method"></span> limit

Defines the maximum number of traces that are returned from Tempo

```java
public QueryV2Builder limit(Long limit)
```

### <span class="badge object-method"></span> maxDuration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```java
public QueryV2Builder maxDuration(String maxDuration)
```

### <span class="badge object-method"></span> metricsQueryType

For metric queries, whether to run instant or range queries

```java
public QueryV2Builder metricsQueryType(MetricsQueryType metricsQueryType)
```

### <span class="badge object-method"></span> minDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```java
public QueryV2Builder minDuration(String minDuration)
```

### <span class="badge object-method"></span> query

TraceQL query or trace ID

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

### <span class="badge object-method"></span> search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```java
public QueryV2Builder search(String search)
```

### <span class="badge object-method"></span> serviceMapIncludeNamespace

Use service.namespace in addition to service.name to uniquely identify a service.

```java
public QueryV2Builder serviceMapIncludeNamespace(Boolean serviceMapIncludeNamespace)
```

### <span class="badge object-method"></span> serviceMapQuery

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.

```java
public QueryV2Builder serviceMapQuery(StringOrArrayOfString serviceMapQuery)
```

### <span class="badge object-method"></span> serviceName

@deprecated Query traces by service name

```java
public QueryV2Builder serviceName(String serviceName)
```

### <span class="badge object-method"></span> spanName

@deprecated Query traces by span name

```java
public QueryV2Builder spanName(String spanName)
```

### <span class="badge object-method"></span> spss

Defines the maximum number of spans per spanset that are returned from Tempo

```java
public QueryV2Builder spss(Long spss)
```

### <span class="badge object-method"></span> step

For metric queries, the step size to use

```java
public QueryV2Builder step(String step)
```

### <span class="badge object-method"></span> tableType

The type of the table that is used to display the search results

```java
public QueryV2Builder tableType(SearchTableType tableType)
```

### <span class="badge object-method"></span> version

```java
public QueryV2Builder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2.DataQueryKind](../dashboardv2/object-DataQueryKind.md)
