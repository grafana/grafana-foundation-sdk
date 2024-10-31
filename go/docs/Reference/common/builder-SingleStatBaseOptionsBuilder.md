---
title: <span class="badge builder"></span> SingleStatBaseOptionsBuilder
---
# <span class="badge builder"></span> SingleStatBaseOptionsBuilder

## Constructor

```go
func NewSingleStatBaseOptionsBuilder() *SingleStatBaseOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SingleStatBaseOptionsBuilder) Build() (SingleStatBaseOptions, error)
```

### <span class="badge object-method"></span> Orientation

```go
func (builder *SingleStatBaseOptionsBuilder) Orientation(orientation common.VizOrientation) *SingleStatBaseOptionsBuilder
```

### <span class="badge object-method"></span> ReduceOptions

```go
func (builder *SingleStatBaseOptionsBuilder) ReduceOptions(reduceOptions cog.Builder[common.ReduceDataOptions]) *SingleStatBaseOptionsBuilder
```

### <span class="badge object-method"></span> Text

```go
func (builder *SingleStatBaseOptionsBuilder) Text(text cog.Builder[common.VizTextDisplayOptions]) *SingleStatBaseOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SingleStatBaseOptions](./object-SingleStatBaseOptions.md)
