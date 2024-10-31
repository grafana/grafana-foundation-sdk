---
title: <span class="badge builder"></span> MetricQueryBuilder
---
# <span class="badge builder"></span> MetricQueryBuilder

## Constructor

```go
func NewMetricQueryBuilder() *MetricQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricQueryBuilder) Build() (MetricQuery, error)
```

### <span class="badge object-method"></span> AliasBy

Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.

```go
func (builder *MetricQueryBuilder) AliasBy(aliasBy string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> AlignmentPeriod

Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.

```go
func (builder *MetricQueryBuilder) AlignmentPeriod(alignmentPeriod string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> CrossSeriesReducer

Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.

```go
func (builder *MetricQueryBuilder) CrossSeriesReducer(crossSeriesReducer string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> EditorMode

```go
func (builder *MetricQueryBuilder) EditorMode(editorMode string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> Filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```go
func (builder *MetricQueryBuilder) Filters(filters []string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> GraphPeriod

To disable the graphPeriod, it should explictly be set to 'disabled'.

```go
func (builder *MetricQueryBuilder) GraphPeriod(graphPeriod string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> GroupBys

Array of labels to group data by.

```go
func (builder *MetricQueryBuilder) GroupBys(groupBys []string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> MetricKind

```go
func (builder *MetricQueryBuilder) MetricKind(metricKind googlecloudmonitoring.MetricKind) *MetricQueryBuilder
```

### <span class="badge object-method"></span> MetricType

```go
func (builder *MetricQueryBuilder) MetricType(metricType string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> PerSeriesAligner

Alignment function to be used. Defaults to ALIGN_MEAN.

```go
func (builder *MetricQueryBuilder) PerSeriesAligner(perSeriesAligner string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> Preprocessor

Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters

```go
func (builder *MetricQueryBuilder) Preprocessor(preprocessor googlecloudmonitoring.PreprocessorType) *MetricQueryBuilder
```

### <span class="badge object-method"></span> ProjectName

GCP project to execute the query against.

```go
func (builder *MetricQueryBuilder) ProjectName(projectName string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> Query

MQL query to be executed.

```go
func (builder *MetricQueryBuilder) Query(query string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> ValueType

```go
func (builder *MetricQueryBuilder) ValueType(valueType string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> View

```go
func (builder *MetricQueryBuilder) View(view string) *MetricQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricQuery](./object-MetricQuery.md)
