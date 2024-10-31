---
title: <span class="badge builder"></span> RawDataBuilder
---
# <span class="badge builder"></span> RawDataBuilder

## Constructor

```go
func NewRawDataBuilder() *RawDataBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RawDataBuilder) Build() (RawData, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *RawDataBuilder) Hide(hide bool) *RawDataBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *RawDataBuilder) Id(id string) *RawDataBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *RawDataBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchRawDataSettings]) *RawDataBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RawData](./object-RawData.md)
