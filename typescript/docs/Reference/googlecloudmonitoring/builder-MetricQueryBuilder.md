---
title: <span class="badge builder"></span> MetricQueryBuilder
---
# <span class="badge builder"></span> MetricQueryBuilder

## Constructor

```typescript
new MetricQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> aliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```typescript
aliasBy(aliasBy: string)
```

### <span class="badge object-method"></span> alignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```typescript
alignmentPeriod(alignmentPeriod: string)
```

### <span class="badge object-method"></span> crossSeriesReducer

Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```typescript
crossSeriesReducer(crossSeriesReducer: string)
```

### <span class="badge object-method"></span> editorMode

```typescript
editorMode(editorMode: string)
```

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```typescript
filters(filters: string[])
```

### <span class="badge object-method"></span> graphPeriod

To disable the graphPeriod, it should explictly be set to 'disabled'.

```typescript
graphPeriod(graphPeriod: string)
```

### <span class="badge object-method"></span> groupBys

Array of labels to group data by.

```typescript
groupBys(groupBys: string[])
```

### <span class="badge object-method"></span> metricKind

```typescript
metricKind(metricKind: googlecloudmonitoring.MetricKind)
```

### <span class="badge object-method"></span> metricType

```typescript
metricType(metricType: string)
```

### <span class="badge object-method"></span> perSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```typescript
perSeriesAligner(perSeriesAligner: string)
```

### <span class="badge object-method"></span> preprocessor

Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters

```typescript
preprocessor(preprocessor: googlecloudmonitoring.PreprocessorType)
```

### <span class="badge object-method"></span> projectName

GCP project to execute the query against.

```typescript
projectName(projectName: string)
```

### <span class="badge object-method"></span> query

MQL query to be executed.

```typescript
query(query: string)
```

### <span class="badge object-method"></span> valueType

```typescript
valueType(valueType: string)
```

### <span class="badge object-method"></span> view

```typescript
view(view: string)
```

## See also

 * <span class="badge object-type-interface"></span> [MetricQuery](./object-MetricQuery.md)
