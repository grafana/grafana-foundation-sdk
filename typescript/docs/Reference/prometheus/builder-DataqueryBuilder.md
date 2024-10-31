---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```typescript
new DataqueryBuilder()
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

### <span class="badge object-method"></span> editorMode

Specifies which editor is being used to prepare the query. It can be "code" or "builder"

```typescript
editorMode(editorMode: prometheus.QueryEditorMode)
```

### <span class="badge object-method"></span> exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```typescript
exemplar(exemplar: boolean)
```

### <span class="badge object-method"></span> expr

The actual expression/query that will be evaluated by Prometheus

```typescript
expr(expr: string)
```

### <span class="badge object-method"></span> format

Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"

```typescript
format(format: prometheus.PromQueryFormat)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```typescript
hide(hide: boolean)
```

### <span class="badge object-method"></span> instant

Returns only the latest value that Prometheus has scraped for the requested time series

```typescript
instant()
```

### <span class="badge object-method"></span> interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```typescript
interval(interval: string)
```

### <span class="badge object-method"></span> intervalFactor

@deprecated Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

```typescript
intervalFactor(intervalFactor: number)
```

### <span class="badge object-method"></span> legendFormat

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```typescript
legendFormat(legendFormat: string)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```typescript
queryType(queryType: string)
```

### <span class="badge object-method"></span> range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```typescript
range()
```

### <span class="badge object-method"></span> rangeAndInstant

```typescript
rangeAndInstant()
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```typescript
refId(refId: string)
```

## See also

 * <span class="badge object-type-interface"></span> [dataquery](./object-dataquery.md)
