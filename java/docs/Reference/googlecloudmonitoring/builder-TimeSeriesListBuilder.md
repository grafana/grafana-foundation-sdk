---
title: <span class="badge builder"></span> TimeSeriesListBuilder
---
# <span class="badge builder"></span> TimeSeriesListBuilder

## Constructor

```java
new TimeSeriesListBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public TimeSeriesList build()
```

### <span class="badge object-method"></span> alignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```java
public TimeSeriesListBuilder alignmentPeriod(String alignmentPeriod)
```

### <span class="badge object-method"></span> crossSeriesReducer

Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```java
public TimeSeriesListBuilder crossSeriesReducer(String crossSeriesReducer)
```

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```java
public TimeSeriesListBuilder filters(List<String> filters)
```

### <span class="badge object-method"></span> groupBys

Array of labels to group data by.

```java
public TimeSeriesListBuilder groupBys(List<String> groupBys)
```

### <span class="badge object-method"></span> perSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```java
public TimeSeriesListBuilder perSeriesAligner(String perSeriesAligner)
```

### <span class="badge object-method"></span> preprocessor

Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters

```java
public TimeSeriesListBuilder preprocessor(PreprocessorType preprocessor)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```java
public TimeSeriesListBuilder projectName(String projectName)
```

### <span class="badge object-method"></span> secondaryAlignmentPeriod

Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```java
public TimeSeriesListBuilder secondaryAlignmentPeriod(String secondaryAlignmentPeriod)
```

### <span class="badge object-method"></span> secondaryCrossSeriesReducer

Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```java
public TimeSeriesListBuilder secondaryCrossSeriesReducer(String secondaryCrossSeriesReducer)
```

### <span class="badge object-method"></span> secondaryGroupBys

Only present if a preprocessor is selected. Array of labels to group data by.

```java
public TimeSeriesListBuilder secondaryGroupBys(List<String> secondaryGroupBys)
```

### <span class="badge object-method"></span> secondaryPerSeriesAligner

Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.

```java
public TimeSeriesListBuilder secondaryPerSeriesAligner(String secondaryPerSeriesAligner)
```

### <span class="badge object-method"></span> text

Annotation text.

```java
public TimeSeriesListBuilder text(String text)
```

### <span class="badge object-method"></span> title

Annotation title.

```java
public TimeSeriesListBuilder title(String title)
```

### <span class="badge object-method"></span> view

Data view, defaults to FULL.

```java
public TimeSeriesListBuilder view(String view)
```

## See also

 * <span class="badge object-type-class"></span> [TimeSeriesList](./object-TimeSeriesList.md)
