---
title: <span class="badge builder"></span> DataqueryBuilder
---
# <span class="badge builder"></span> DataqueryBuilder

## Constructor

```java
new DataqueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Dataquery build()
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public dataqueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> editorMode

```java
public dataqueryBuilder editorMode(QueryEditorMode editorMode)
```

### <span class="badge object-method"></span> expr

The LogQL query.

```java
public dataqueryBuilder expr(String expr)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public dataqueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> instant

@deprecated, now use queryType.

```java
public dataqueryBuilder instant(Boolean instant)
```

### <span class="badge object-method"></span> legendFormat

Used to override the name of the series.

```java
public dataqueryBuilder legendFormat(String legendFormat)
```

### <span class="badge object-method"></span> maxLines

Used to limit the number of log rows returned.

```java
public dataqueryBuilder maxLines(Long maxLines)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public dataqueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> range

@deprecated, now use queryType.

```java
public dataqueryBuilder range(Boolean range)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public dataqueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> resolution

@deprecated, now use step.

```java
public dataqueryBuilder resolution(Long resolution)
```

### <span class="badge object-method"></span> step

Used to set step value for range queries.

```java
public dataqueryBuilder step(String step)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
