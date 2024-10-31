---
title: <span class="badge builder"></span> RangeMapBuilder
---
# <span class="badge builder"></span> RangeMapBuilder

## Constructor

```go
func NewRangeMapBuilder() *RangeMapBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RangeMapBuilder) Build() (RangeMap, error)
```

### <span class="badge object-method"></span> Options

Range to match against and the result to apply when the value is within the range

```go
func (builder *RangeMapBuilder) Options(options cog.Builder[dashboard.DashboardRangeMapOptions]) *RangeMapBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RangeMap](./object-RangeMap.md)
