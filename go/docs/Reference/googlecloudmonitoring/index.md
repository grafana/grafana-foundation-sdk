# googlecloudmonitoring

## Objects

 * <span class="badge object-type-enum"></span> [AlignmentTypes](./object-AlignmentTypes.md)
 * <span class="badge object-type-struct"></span> [AnnotationQuery](./object-AnnotationQuery.md)
 * <span class="badge object-type-struct"></span> [CloudMonitoringQuery](./object-CloudMonitoringQuery.md)
 * <span class="badge object-type-struct"></span> [Filter](./object-Filter.md)
 * <span class="badge object-type-struct"></span> [LegacyCloudMonitoringAnnotationQuery](./object-LegacyCloudMonitoringAnnotationQuery.md)
 * <span class="badge object-type-enum"></span> [MetricFindQueryTypes](./object-MetricFindQueryTypes.md)
 * <span class="badge object-type-enum"></span> [MetricKind](./object-MetricKind.md)
 * <span class="badge object-type-struct"></span> [MetricQuery](./object-MetricQuery.md)
 * <span class="badge object-type-enum"></span> [PreprocessorType](./object-PreprocessorType.md)
 * <span class="badge object-type-enum"></span> [QueryType](./object-QueryType.md)
 * <span class="badge object-type-struct"></span> [SLOQuery](./object-SLOQuery.md)
 * <span class="badge object-type-struct"></span> [TimeSeriesList](./object-TimeSeriesList.md)
 * <span class="badge object-type-struct"></span> [TimeSeriesQuery](./object-TimeSeriesQuery.md)
 * <span class="badge object-type-enum"></span> [ValueTypes](./object-ValueTypes.md)
## Builders

 * <span class="badge builder"></span> [AnnotationQueryBuilder](./builder-AnnotationQueryBuilder.md)
 * <span class="badge builder"></span> [CloudMonitoringQueryBuilder](./builder-CloudMonitoringQueryBuilder.md)
 * <span class="badge builder"></span> [FilterBuilder](./builder-FilterBuilder.md)
 * <span class="badge builder"></span> [LegacyCloudMonitoringAnnotationQueryBuilder](./builder-LegacyCloudMonitoringAnnotationQueryBuilder.md)
 * <span class="badge builder"></span> [MetricQueryBuilder](./builder-MetricQueryBuilder.md)
 * <span class="badge builder"></span> [SLOQueryBuilder](./builder-SLOQueryBuilder.md)
 * <span class="badge builder"></span> [TimeSeriesListBuilder](./builder-TimeSeriesListBuilder.md)
 * <span class="badge builder"></span> [TimeSeriesQueryBuilder](./builder-TimeSeriesQueryBuilder.md)
## Functions

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to cloud-monitoring dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> CloudMonitoringQueryConverter

CloudMonitoringQueryConverter accepts a `CloudMonitoringQuery` object and generates the Go code to build this object using builders.

```go
func CloudMonitoringQueryConverter(input CloudMonitoringQuery) string
```

### <span class="badge function"></span> TimeSeriesListConverter

TimeSeriesListConverter accepts a `TimeSeriesList` object and generates the Go code to build this object using builders.

```go
func TimeSeriesListConverter(input TimeSeriesList) string
```

### <span class="badge function"></span> AnnotationQueryConverter

AnnotationQueryConverter accepts a `AnnotationQuery` object and generates the Go code to build this object using builders.

```go
func AnnotationQueryConverter(input AnnotationQuery) string
```

### <span class="badge function"></span> TimeSeriesQueryConverter

TimeSeriesQueryConverter accepts a `TimeSeriesQuery` object and generates the Go code to build this object using builders.

```go
func TimeSeriesQueryConverter(input TimeSeriesQuery) string
```

### <span class="badge function"></span> SLOQueryConverter

SLOQueryConverter accepts a `SLOQuery` object and generates the Go code to build this object using builders.

```go
func SLOQueryConverter(input SLOQuery) string
```

### <span class="badge function"></span> MetricQueryConverter

MetricQueryConverter accepts a `MetricQuery` object and generates the Go code to build this object using builders.

```go
func MetricQueryConverter(input MetricQuery) string
```

### <span class="badge function"></span> LegacyCloudMonitoringAnnotationQueryConverter

LegacyCloudMonitoringAnnotationQueryConverter accepts a `LegacyCloudMonitoringAnnotationQuery` object and generates the Go code to build this object using builders.

```go
func LegacyCloudMonitoringAnnotationQueryConverter(input LegacyCloudMonitoringAnnotationQuery) string
```

### <span class="badge function"></span> FilterConverter

FilterConverter accepts a `Filter` object and generates the Go code to build this object using builders.

```go
func FilterConverter(input Filter) string
```

