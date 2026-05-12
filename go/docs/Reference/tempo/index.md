# <span class="badge package-variant-dataquery"></span> tempo

## Objects

 * <span class="badge object-type-struct"></span> [Dataquery](./object-Dataquery.md)
 * <span class="badge object-type-enum"></span> [MetricsQueryType](./object-MetricsQueryType.md)
 * <span class="badge object-type-enum"></span> [SearchStreamingState](./object-SearchStreamingState.md)
 * <span class="badge object-type-enum"></span> [SearchTableType](./object-SearchTableType.md)
 * <span class="badge object-type-struct"></span> [StringOrArrayOfString](./object-StringOrArrayOfString.md)
 * <span class="badge object-type-enum"></span> [TempoQueryType](./object-TempoQueryType.md)
 * <span class="badge object-type-struct"></span> [TraceqlFilter](./object-TraceqlFilter.md)
 * <span class="badge object-type-enum"></span> [TraceqlSearchScope](./object-TraceqlSearchScope.md)
## Builders

 * <span class="badge builder"></span> [DataqueryBuilder](./builder-DataqueryBuilder.md)
 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)
 * <span class="badge builder"></span> [QueryV2Builder](./builder-QueryV2Builder.md)
 * <span class="badge builder"></span> [TraceqlFilterBuilder](./builder-TraceqlFilterBuilder.md)
## Functions

### <span class="badge function"></span> NewDataquery

NewDataquery creates a new Dataquery object.

```go
func NewDataquery() *Dataquery
```

### <span class="badge function"></span> NewTraceqlFilter

NewTraceqlFilter creates a new TraceqlFilter object.

```go
func NewTraceqlFilter() *TraceqlFilter
```

### <span class="badge function"></span> NewStringOrArrayOfString

NewStringOrArrayOfString creates a new StringOrArrayOfString object.

```go
func NewStringOrArrayOfString() *StringOrArrayOfString
```

### <span class="badge function"></span> VariantConfig

VariantConfig returns the configuration related to tempo dataqueries.

This configuration describes how to unmarshal it, convert it to code, …

```go
func VariantConfig() variants.DataqueryConfig
```

### <span class="badge function"></span> QueryV2Converter

QueryV2Converter accepts a `QueryV2` object and generates the Go code to build this object using builders.

```go
func QueryV2Converter(input dashboardv2.DataQueryKind) string
```

### <span class="badge function"></span> DataqueryConverter

DataqueryConverter accepts a `Dataquery` object and generates the Go code to build this object using builders.

```go
func DataqueryConverter(input Dataquery) string
```

### <span class="badge function"></span> TraceqlFilterConverter

TraceqlFilterConverter accepts a `TraceqlFilter` object and generates the Go code to build this object using builders.

```go
func TraceqlFilterConverter(input TraceqlFilter) string
```

### <span class="badge function"></span> QueryConverter

QueryConverter accepts a `Query` object and generates the Go code to build this object using builders.

```go
func QueryConverter(input dashboardv2beta1.DataQueryKind) string
```

