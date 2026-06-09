---
title: <span class="badge builder"></span> RequestBuilder
---
# <span class="badge builder"></span> RequestBuilder

## Constructor

```go
func NewRequestBuilder() *RequestBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RequestBuilder) Build() (variants.Dataquery, error)
```

### <span class="badge object-method"></span> AnnotationQuery

```go
func (builder *RequestBuilder) AnnotationQuery(annotationQuery cog.Builder[cloudwatch.AnnotationQuery]) *RequestBuilder
```

### <span class="badge object-method"></span> LogsQuery

```go
func (builder *RequestBuilder) LogsQuery(logsQuery cog.Builder[cloudwatch.LogsQuery]) *RequestBuilder
```

### <span class="badge object-method"></span> MetricsQuery

```go
func (builder *RequestBuilder) MetricsQuery(metricsQuery cog.Builder[cloudwatch.MetricsQuery]) *RequestBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [Request](./object-Request.md)
