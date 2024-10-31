---
title: <span class="badge builder"></span> TimeSeriesListBuilder
---
# <span class="badge builder"></span> TimeSeriesListBuilder

## Constructor

```go
func NewTimeSeriesListBuilder() *TimeSeriesListBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TimeSeriesListBuilder) Build() (TimeSeriesList, error)
```

### <span class="badge object-method"></span> AlignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```go
func (builder *TimeSeriesListBuilder) AlignmentPeriod(alignmentPeriod string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> CrossSeriesReducer

Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```go
func (builder *TimeSeriesListBuilder) CrossSeriesReducer(crossSeriesReducer string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> Filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```go
func (builder *TimeSeriesListBuilder) Filters(filters []string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> GroupBys

Array of labels to group data by.

```go
func (builder *TimeSeriesListBuilder) GroupBys(groupBys []string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> PerSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```go
func (builder *TimeSeriesListBuilder) PerSeriesAligner(perSeriesAligner string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> Preprocessor

Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters

```go
func (builder *TimeSeriesListBuilder) Preprocessor(preprocessor googlecloudmonitoring.PreprocessorType) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> ProjectName

GCP project to execute the query against.

```go
func (builder *TimeSeriesListBuilder) ProjectName(projectName string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> SecondaryAlignmentPeriod

Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```go
func (builder *TimeSeriesListBuilder) SecondaryAlignmentPeriod(secondaryAlignmentPeriod string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> SecondaryCrossSeriesReducer

Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```go
func (builder *TimeSeriesListBuilder) SecondaryCrossSeriesReducer(secondaryCrossSeriesReducer string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> SecondaryGroupBys

Only present if a preprocessor is selected. Array of labels to group data by.

```go
func (builder *TimeSeriesListBuilder) SecondaryGroupBys(secondaryGroupBys []string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> SecondaryPerSeriesAligner

Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.

```go
func (builder *TimeSeriesListBuilder) SecondaryPerSeriesAligner(secondaryPerSeriesAligner string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> Text

Annotation text.

```go
func (builder *TimeSeriesListBuilder) Text(text string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> Title

Annotation title.

```go
func (builder *TimeSeriesListBuilder) Title(title string) *TimeSeriesListBuilder
```

### <span class="badge object-method"></span> View

Data view, defaults to FULL.

```go
func (builder *TimeSeriesListBuilder) View(view string) *TimeSeriesListBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimeSeriesList](./object-TimeSeriesList.md)
