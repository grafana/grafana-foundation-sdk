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

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```typescript
datasource(ref: {
	name?: string;
})
```

### <span class="badge object-method"></span> exemplars

For metric queries, how many exemplars to request, 0 means no exemplars

```typescript
exemplars(exemplars: number)
```

### <span class="badge object-method"></span> filters

```typescript
filters(filters: cog.Builder<tempo.TraceqlFilter>[])
```

### <span class="badge object-method"></span> groupBy

Filters that are used to query the metrics summary

```typescript
groupBy(groupBy: cog.Builder<tempo.TraceqlFilter>[])
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> limit

Defines the maximum number of traces that are returned from Tempo

```typescript
limit(limit: number)
```

### <span class="badge object-method"></span> maxDuration

@deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms

```typescript
maxDuration(maxDuration: string)
```

### <span class="badge object-method"></span> metricsQueryType

For metric queries, whether to run instant or range queries

```typescript
metricsQueryType(metricsQueryType: tempo.MetricsQueryType)
```

### <span class="badge object-method"></span> minDuration

@deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms

```typescript
minDuration(minDuration: string)
```

### <span class="badge object-method"></span> query

TraceQL query or trace ID

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

### <span class="badge object-method"></span> search

@deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true

```typescript
search(search: string)
```

### <span class="badge object-method"></span> serviceMapIncludeNamespace

Use service.namespace in addition to service.name to uniquely identify a service.

```typescript
serviceMapIncludeNamespace(serviceMapIncludeNamespace: boolean)
```

### <span class="badge object-method"></span> serviceMapQuery

Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.

```typescript
serviceMapQuery(serviceMapQuery: string | string[])
```

### <span class="badge object-method"></span> serviceName

@deprecated Query traces by service name

```typescript
serviceName(serviceName: string)
```

### <span class="badge object-method"></span> spanName

@deprecated Query traces by span name

```typescript
spanName(spanName: string)
```

### <span class="badge object-method"></span> spss

Defines the maximum number of spans per spanset that are returned from Tempo

```typescript
spss(spss: number)
```

### <span class="badge object-method"></span> step

For metric queries, the step size to use

```typescript
step(step: string)
```

### <span class="badge object-method"></span> tableType

The type of the table that is used to display the search results

```typescript
tableType(tableType: tempo.SearchTableType)
```

### <span class="badge object-method"></span> version

```typescript
version(version: string)
```

## See also

 * <span class="badge object-type-interface"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
