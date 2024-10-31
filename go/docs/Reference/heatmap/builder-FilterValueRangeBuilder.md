---
title: <span class="badge builder"></span> FilterValueRangeBuilder
---
# <span class="badge builder"></span> FilterValueRangeBuilder

## Constructor

```go
func NewFilterValueRangeBuilder() *FilterValueRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FilterValueRangeBuilder) Build() (FilterValueRange, error)
```

### <span class="badge object-method"></span> Ge

Sets the filter range to values greater than or equal to the given value

```go
func (builder *FilterValueRangeBuilder) Ge(ge float32) *FilterValueRangeBuilder
```

### <span class="badge object-method"></span> Le

Sets the filter range to values less than or equal to the given value

```go
func (builder *FilterValueRangeBuilder) Le(le float32) *FilterValueRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FilterValueRange](./object-FilterValueRange.md)
