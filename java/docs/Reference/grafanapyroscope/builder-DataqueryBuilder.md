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

### <span class="badge object-method"></span> groupBy

Allows to group the results.

```java
public dataqueryBuilder groupBy(List<String> groupBy)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public dataqueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> labelSelector

Specifies the query label selectors.

```java
public dataqueryBuilder labelSelector(String labelSelector)
```

### <span class="badge object-method"></span> limit

Sets the maximum number of time series.

```java
public dataqueryBuilder limit(Long limit)
```

### <span class="badge object-method"></span> maxNodes

Sets the maximum number of nodes in the flamegraph.

```java
public dataqueryBuilder maxNodes(Long maxNodes)
```

### <span class="badge object-method"></span> profileTypeId

Specifies the type of profile to query.

```java
public dataqueryBuilder profileTypeId(String profileTypeId)
```

### <span class="badge object-method"></span> queryType

Specify the query flavor

TODO make this required and give it a default

```java
public dataqueryBuilder queryType(String queryType)
```

### <span class="badge object-method"></span> refId

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```java
public dataqueryBuilder refId(String refId)
```

### <span class="badge object-method"></span> spanSelector

Specifies the query span selectors.

```java
public dataqueryBuilder spanSelector(List<String> spanSelector)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
