---
title: <span class="badge builder"></span> ExtendedStatBuilder
---
# <span class="badge builder"></span> ExtendedStatBuilder

## Constructor

```go
func NewExtendedStatBuilder() *ExtendedStatBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ExtendedStatBuilder) Build() (ExtendedStat, error)
```

### <span class="badge object-method"></span> Label

```go
func (builder *ExtendedStatBuilder) Label(label string) *ExtendedStatBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *ExtendedStatBuilder) Value(value elasticsearch.ExtendedStatMetaType) *ExtendedStatBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ExtendedStat](./object-ExtendedStat.md)
