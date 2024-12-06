---
title: <span class="badge builder"></span> TempoQueryBuilder
---
# <span class="badge builder"></span> TempoQueryBuilder

## Constructor

```typescript
new TempoQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```typescript
datasource(datasource: dashboard.DataSourceRef)
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

## See also

 * <span class="badge object-type-interface"></span> [TempoQuery](./object-TempoQuery.md)
