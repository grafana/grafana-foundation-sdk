---
title: <span class="badge builder"></span> LegacyCloudMonitoringAnnotationQueryBuilder
---
# <span class="badge builder"></span> LegacyCloudMonitoringAnnotationQueryBuilder

## Constructor

```go
func NewLegacyCloudMonitoringAnnotationQueryBuilder() *LegacyCloudMonitoringAnnotationQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) Build() (LegacyCloudMonitoringAnnotationQuery, error)
```

### <span class="badge object-method"></span> Filters

Array of filters to query data by. Labels that can be filtered on are defined by the metric.

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) Filters(filters []string) *LegacyCloudMonitoringAnnotationQueryBuilder
```

### <span class="badge object-method"></span> MetricKind

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) MetricKind(metricKind googlecloudmonitoring.MetricKind) *LegacyCloudMonitoringAnnotationQueryBuilder
```

### <span class="badge object-method"></span> MetricType

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) MetricType(metricType string) *LegacyCloudMonitoringAnnotationQueryBuilder
```

### <span class="badge object-method"></span> ProjectName

GCP project to execute the query against.

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) ProjectName(projectName string) *LegacyCloudMonitoringAnnotationQueryBuilder
```

### <span class="badge object-method"></span> RefId

Query refId.

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) RefId(refId string) *LegacyCloudMonitoringAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Text

Annotation text.

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) Text(text string) *LegacyCloudMonitoringAnnotationQueryBuilder
```

### <span class="badge object-method"></span> Title

Annotation title.

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) Title(title string) *LegacyCloudMonitoringAnnotationQueryBuilder
```

### <span class="badge object-method"></span> ValueType

```go
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) ValueType(valueType string) *LegacyCloudMonitoringAnnotationQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LegacyCloudMonitoringAnnotationQuery](./object-LegacyCloudMonitoringAnnotationQuery.md)
