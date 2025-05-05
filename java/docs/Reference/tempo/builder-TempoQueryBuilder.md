---
title: <span class="badge builder"></span> TempoQueryBuilder
---
# <span class="badge builder"></span> TempoQueryBuilder

## Constructor

```java
new TempoQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TempoQuery build()
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public TempoQueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> exemplars

For metric queries, how many exemplars to request, 0 means no exemplars

```java
public TempoQueryBuilder exemplars(Long exemplars)
```

### <span class="badge object-method"></span> filters

```java
public TempoQueryBuilder filters(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> filters)
```

### <span class="badge object-method"></span> groupBy

deprecated Filters that are used to query the metrics summary

```java
public TempoQueryBuilder groupBy(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> groupBy)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public TempoQueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> limit

Defines the maximum number of traces that are returned from Tempo

```java
public TempoQueryBuilder limit(Long limit)
```

### <span class="badge object-method"></span> maxDuration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```java
public TempoQueryBuilder maxDuration(String maxDuration)
```

### <span class="badge object-method"></span> metricsQueryType

For metric queries, whether to run instant or range queries

```java
public TempoQueryBuilder metricsQueryType(MetricsQueryType metricsQueryType)
```

### <span class="badge object-method"></span> minDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```java
public TempoQueryBuilder minDuration(String minDuration)
```

### <span class="badge object-method"></span> query

TraceQL query or trace ID

```java
public TempoQueryBuilder query(String query)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public TempoQueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public TempoQueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```java
public TempoQueryBuilder search(String search)
```

### <span class="badge object-method"></span> serviceMapIncludeNamespace

Use service.namespace in addition to service.name to uniquely identify a service.

```java
public TempoQueryBuilder serviceMapIncludeNamespace(Boolean serviceMapIncludeNamespace)
```

### <span class="badge object-method"></span> serviceMapQuery

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.

```java
public TempoQueryBuilder serviceMapQuery(StringOrArrayOfString serviceMapQuery)
```

### <span class="badge object-method"></span> serviceName

@deprecated Query traces by service name

```java
public TempoQueryBuilder serviceName(String serviceName)
```

### <span class="badge object-method"></span> spanName

@deprecated Query traces by span name

```java
public TempoQueryBuilder spanName(String spanName)
```

### <span class="badge object-method"></span> spss

Defines the maximum number of spans per spanset that are returned from Tempo

```java
public TempoQueryBuilder spss(Long spss)
```

### <span class="badge object-method"></span> step

For metric queries, the step size to use

```java
public TempoQueryBuilder step(String step)
```

### <span class="badge object-method"></span> tableType

The type of the table that is used to display the search results

```java
public TempoQueryBuilder tableType(SearchTableType tableType)
```

## See also

 * <span class="badge object-type-class"></span> [TempoQuery](./object-TempoQuery.md)
