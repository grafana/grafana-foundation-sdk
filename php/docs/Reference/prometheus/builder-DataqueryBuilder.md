---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```php
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```php
datasource(\Grafana\Foundation\Dashboard\DataSourceRef $datasource)
```

### <span class="badge object-method"></span> editorMode

Specifies which editor is being used to prepare the query. It can be "code" or "builder"

```php
editorMode(\Grafana\Foundation\Prometheus\QueryEditorMode $editorMode)
```

### <span class="badge object-method"></span> exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```php
exemplar(bool $exemplar)
```

### <span class="badge object-method"></span> expr

The actual expression/query that will be evaluated by Prometheus

```php
expr(string $expr)
```

### <span class="badge object-method"></span> format

Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"

```php
format(\Grafana\Foundation\Prometheus\PromQueryFormat $format)
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```php
hide(bool $hide)
```

### <span class="badge object-method"></span> instant

Returns only the latest value that Prometheus has scraped for the requested time series

```php
instant()
```

### <span class="badge object-method"></span> interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```php
interval(string $interval)
```

### <span class="badge object-method"></span> intervalFactor

@deprecated Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

```php
intervalFactor(float $intervalFactor)
```

### <span class="badge object-method"></span> legendFormat

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```php
legendFormat(string $legendFormat)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```php
queryType(string $queryType)
```

### <span class="badge object-method"></span> range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```php
range()
```

### <span class="badge object-method"></span> rangeAndInstant

```php
rangeAndInstant()
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```php
refId(string $refId)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
