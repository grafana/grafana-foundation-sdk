---
title: <span class="badge builder"></span> RepeatOptionsBuilder
---
# <span class="badge builder"></span> RepeatOptionsBuilder

## Constructor

```go
func NewRepeatOptionsBuilder() *RepeatOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RepeatOptionsBuilder) Build() (RepeatOptions, error)
```

### <span class="badge object-method"></span> Direction

```go
func (builder *RepeatOptionsBuilder) Direction(direction dashboardv2beta1.RepeatOptionsDirection) *RepeatOptionsBuilder
```

### <span class="badge object-method"></span> MaxPerRow

```go
func (builder *RepeatOptionsBuilder) MaxPerRow(maxPerRow int64) *RepeatOptionsBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *RepeatOptionsBuilder) Value(value string) *RepeatOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RepeatOptions](./object-RepeatOptions.md)
