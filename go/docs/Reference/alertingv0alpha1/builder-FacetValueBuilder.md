---
title: <span class="badge builder"></span> FacetValueBuilder
---
# <span class="badge builder"></span> FacetValueBuilder

#FacetValue is a single value/count pair in a facet breakdown.

## Constructor

```go
func NewFacetValueBuilder() *FacetValueBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FacetValueBuilder) Build() (FacetValue, error)
```

### <span class="badge object-method"></span> Count

```go
func (builder *FacetValueBuilder) Count(count int64) *FacetValueBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *FacetValueBuilder) Value(value string) *FacetValueBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FacetValue](./object-FacetValue.md)
