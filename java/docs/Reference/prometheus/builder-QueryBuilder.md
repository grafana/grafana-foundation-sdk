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

### <span class="badge object-method"></span> datasource

New type for datasource reference

Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.

```java
public QueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource)
```

### <span class="badge object-method"></span> editorMode

Specifies which editor is being used to prepare the query. It can be "code" or "builder"

```java
public QueryBuilder editorMode(QueryEditorMode editorMode)
```

### <span class="badge object-method"></span> exemplar

Execute an additional query to identify interesting raw samples relevant for the given expr

```java
public QueryBuilder exemplar(Boolean exemplar)
```

### <span class="badge object-method"></span> expr

The actual expression/query that will be evaluated by Prometheus

```java
public QueryBuilder expr(String expr)
```

### <span class="badge object-method"></span> format

Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"

```java
public QueryBuilder format(PromQueryFormat format)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public QueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> instant

Returns only the latest value that Prometheus has scraped for the requested time series

```java
public QueryBuilder instant()
```

### <span class="badge object-method"></span> interval

An additional lower limit for the step parameter of the Prometheus query and for the

`$__interval` and `$__rate_interval` variables.

```java
public QueryBuilder interval(String interval)
```

### <span class="badge object-method"></span> intervalFactor

@deprecated Used to specify how many times to divide max data points by. We use max data points under query options

See https://github.com/grafana/grafana/issues/48081

```java
public QueryBuilder intervalFactor(Double intervalFactor)
```

### <span class="badge object-method"></span> legendFormat

Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname

```java
public QueryBuilder legendFormat(String legendFormat)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public QueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> range

Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series

```java
public QueryBuilder range()
```

### <span class="badge object-method"></span> rangeAndInstant

```java
public QueryBuilder rangeAndInstant()
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public QueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> version

```java
public QueryBuilder version(String version)
```

## See also

 * <span class="badge object-type-class"></span> [dashboardv2beta1.DataQueryKind](../dashboardv2beta1/object-DataQueryKind.md)
