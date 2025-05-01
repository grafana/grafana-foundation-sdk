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

### <span class="badge object-method"></span> alias

Alias pattern

```java
public dataqueryBuilder alias(String alias)
```

### <span class="badge object-method"></span> bucketAggs

List of bucket aggregations

```java
public dataqueryBuilder bucketAggs(List<BucketAggregation> bucketAggs)
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```java
public dataqueryBuilder datasource(DataSourceRef datasource)
```

### <span class="badge object-method"></span> hide

If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.

```java
public dataqueryBuilder hide(Boolean hide)
```

### <span class="badge object-method"></span> metrics

List of metric aggregations

```java
public dataqueryBuilder metrics(List<MetricAggregation> metrics)
```

### <span class="badge object-method"></span> query

Lucene query

```java
public dataqueryBuilder query(String query)
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

### <span class="badge object-method"></span> timeField

Name of time field

```java
public dataqueryBuilder timeField(String timeField)
```

## See also

 * <span class="badge object-type-class"></span> [Dataquery](./object-Dataquery.md)
