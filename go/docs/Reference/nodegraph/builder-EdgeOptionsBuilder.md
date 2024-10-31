---
title: <span class="badge builder"></span> EdgeOptionsBuilder
---
# <span class="badge builder"></span> EdgeOptionsBuilder

## Constructor

```go
func NewEdgeOptionsBuilder() *EdgeOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *EdgeOptionsBuilder) Build() (EdgeOptions, error)
```

### <span class="badge object-method"></span> MainStatUnit

Unit for the main stat to override what ever is set in the data frame.

```go
func (builder *EdgeOptionsBuilder) MainStatUnit(mainStatUnit string) *EdgeOptionsBuilder
```

### <span class="badge object-method"></span> SecondaryStatUnit

Unit for the secondary stat to override what ever is set in the data frame.

```go
func (builder *EdgeOptionsBuilder) SecondaryStatUnit(secondaryStatUnit string) *EdgeOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [EdgeOptions](./object-EdgeOptions.md)
