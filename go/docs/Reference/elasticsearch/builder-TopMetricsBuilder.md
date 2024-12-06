---
title: <span class="badge builder"></span> TopMetricsBuilder
---
# <span class="badge builder"></span> TopMetricsBuilder

## Constructor

```go
func NewTopMetricsBuilder() *TopMetricsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TopMetricsBuilder) Build() (TopMetrics, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *TopMetricsBuilder) Hide(hide bool) *TopMetricsBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *TopMetricsBuilder) Id(id string) *TopMetricsBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *TopMetricsBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchTopMetricsSettings]) *TopMetricsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TopMetrics](./object-TopMetrics.md)
