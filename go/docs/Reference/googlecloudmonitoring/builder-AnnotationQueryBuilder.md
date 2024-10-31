---
title: <span class="badge builder"></span> AnnotationQueryBuilder
---
# <span class="badge builder"></span> AnnotationQueryBuilder

## Constructor

```go
func NewAnnotationQueryBuilder() *AnnotationQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AnnotationQueryBuilder) Build() (AnnotationQuery, error)
```

### <span class="badge object-method"></span> AlignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```go
func (builder *AnnotationQueryBuilder) AlignmentPeriod(alignmentPeriod string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> CrossSeriesReducer

Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```go
func (builder *AnnotationQueryBuilder) CrossSeriesReducer(crossSeriesReducer string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```go
func (builder *AnnotationQueryBuilder) Filters(filters []string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> GroupBys

Array of labels to group data by.

```go
func (builder *AnnotationQueryBuilder) GroupBys(groupBys []string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> PerSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```go
func (builder *AnnotationQueryBuilder) PerSeriesAligner(perSeriesAligner string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Preprocessor

Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters

```go
func (builder *AnnotationQueryBuilder) Preprocessor(preprocessor googlecloudmonitoring.PreprocessorType) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> ProjectName

GCP project to execute the query against.

```go
func (builder *AnnotationQueryBuilder) ProjectName(projectName string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> SecondaryAlignmentPeriod

Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```go
func (builder *AnnotationQueryBuilder) SecondaryAlignmentPeriod(secondaryAlignmentPeriod string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> SecondaryCrossSeriesReducer

Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```go
func (builder *AnnotationQueryBuilder) SecondaryCrossSeriesReducer(secondaryCrossSeriesReducer string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> SecondaryGroupBys

Only present if a preprocessor is selected. Array of labels to group data by.

```go
func (builder *AnnotationQueryBuilder) SecondaryGroupBys(secondaryGroupBys []string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> SecondaryPerSeriesAligner

Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.

```go
func (builder *AnnotationQueryBuilder) SecondaryPerSeriesAligner(secondaryPerSeriesAligner string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Text

Annotation text.

```go
func (builder *AnnotationQueryBuilder) Text(text string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> Title

Annotation title.

```go
func (builder *AnnotationQueryBuilder) Title(title string) *AnnotationQueryBuilder
```

### <span class="badge object-method"></span> View

Data view, defaults to FULL.

```go
func (builder *AnnotationQueryBuilder) View(view string) *AnnotationQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AnnotationQuery](./object-AnnotationQuery.md)
