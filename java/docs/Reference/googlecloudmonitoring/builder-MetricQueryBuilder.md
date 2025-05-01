---
title: <span class="badge builder"></span> MetricQueryBuilder
---
# <span class="badge builder"></span> MetricQueryBuilder

## Constructor

```java
new MetricQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public MetricQuery build()
```

### <span class="badge object-method"></span> aliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```java
public MetricQueryBuilder aliasBy(String aliasBy)
```

### <span class="badge object-method"></span> alignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```java
public MetricQueryBuilder alignmentPeriod(String alignmentPeriod)
```

### <span class="badge object-method"></span> crossSeriesReducer

Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```java
public MetricQueryBuilder crossSeriesReducer(String crossSeriesReducer)
```

### <span class="badge object-method"></span> editorMode

```java
public MetricQueryBuilder editorMode(String editorMode)
```

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```java
public MetricQueryBuilder filters(List<String> filters)
```

### <span class="badge object-method"></span> graphPeriod

To disable the graphPeriod, it should explictly be set to 'disabled'.

```java
public MetricQueryBuilder graphPeriod(String graphPeriod)
```

### <span class="badge object-method"></span> groupBys

Array of labels to group data by.

```java
public MetricQueryBuilder groupBys(List<String> groupBys)
```

### <span class="badge object-method"></span> metricKind

```java
public MetricQueryBuilder metricKind(MetricKind metricKind)
```

### <span class="badge object-method"></span> metricType

```java
public MetricQueryBuilder metricType(String metricType)
```

### <span class="badge object-method"></span> perSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```java
public MetricQueryBuilder perSeriesAligner(String perSeriesAligner)
```

### <span class="badge object-method"></span> preprocessor

Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters

```java
public MetricQueryBuilder preprocessor(PreprocessorType preprocessor)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```java
public MetricQueryBuilder projectName(String projectName)
```

### <span class="badge object-method"></span> query

MQL query to be executed.

```java
public MetricQueryBuilder query(String query)
```

### <span class="badge object-method"></span> valueType

```java
public MetricQueryBuilder valueType(String valueType)
```

### <span class="badge object-method"></span> view

```java
public MetricQueryBuilder view(String view)
```

## See also

 * <span class="badge object-type-class"></span> [MetricQuery](./object-MetricQuery.md)
