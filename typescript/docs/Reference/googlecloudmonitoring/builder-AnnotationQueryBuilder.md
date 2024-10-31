---
title: <span class="badge builder"></span> AnnotationQueryBuilder
---
# <span class="badge builder"></span> AnnotationQueryBuilder

## Constructor

```typescript
new AnnotationQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
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

### <span class="badge object-method"></span> filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```typescript
filters(filters: string[])
```

### <span class="badge object-method"></span> groupBys

Array of labels to group data by.

```typescript
groupBys(groupBys: string[])
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

### <span class="badge object-method"></span> secondaryAlignmentPeriod

Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```typescript
secondaryAlignmentPeriod(secondaryAlignmentPeriod: string)
```

### <span class="badge object-method"></span> secondaryCrossSeriesReducer

Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```typescript
secondaryCrossSeriesReducer(secondaryCrossSeriesReducer: string)
```

### <span class="badge object-method"></span> secondaryGroupBys

Only present if a preprocessor is selected. Array of labels to group data by.

```typescript
secondaryGroupBys(secondaryGroupBys: string[])
```

### <span class="badge object-method"></span> secondaryPerSeriesAligner

Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.

```typescript
secondaryPerSeriesAligner(secondaryPerSeriesAligner: string)
```

### <span class="badge object-method"></span> text

Annotation text.

```typescript
text(text: string)
```

### <span class="badge object-method"></span> title

Annotation title.

```typescript
title(title: string)
```

### <span class="badge object-method"></span> view

Data view, defaults to FULL.

```typescript
view(view: string)
```

## See also

 * <span class="badge object-type-interface"></span> [AnnotationQuery](./object-AnnotationQuery.md)
