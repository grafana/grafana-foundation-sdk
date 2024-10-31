---
title: <span class="badge builder"></span> SerialDiffBuilder
---
# <span class="badge builder"></span> SerialDiffBuilder

## Constructor

```go
func NewSerialDiffBuilder() *SerialDiffBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SerialDiffBuilder) Build() (SerialDiff, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *SerialDiffBuilder) Field(field string) *SerialDiffBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *SerialDiffBuilder) Hide(hide bool) *SerialDiffBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *SerialDiffBuilder) Id(id string) *SerialDiffBuilder
```

### <span class="badge object-method"></span> PipelineAgg

```go
func (builder *SerialDiffBuilder) PipelineAgg(pipelineAgg string) *SerialDiffBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *SerialDiffBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchSerialDiffSettings]) *SerialDiffBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SerialDiff](./object-SerialDiff.md)
