---
title: <span class="badge builder"></span> CumulativeSumBuilder
---
# <span class="badge builder"></span> CumulativeSumBuilder

## Constructor

```go
func NewCumulativeSumBuilder() *CumulativeSumBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *CumulativeSumBuilder) Build() (CumulativeSum, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *CumulativeSumBuilder) Field(field string) *CumulativeSumBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *CumulativeSumBuilder) Hide(hide bool) *CumulativeSumBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *CumulativeSumBuilder) Id(id string) *CumulativeSumBuilder
```

### <span class="badge object-method"></span> PipelineAgg

```go
func (builder *CumulativeSumBuilder) PipelineAgg(pipelineAgg string) *CumulativeSumBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *CumulativeSumBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchCumulativeSumSettings]) *CumulativeSumBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [CumulativeSum](./object-CumulativeSum.md)
