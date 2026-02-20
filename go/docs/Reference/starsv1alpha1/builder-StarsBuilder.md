---
title: <span class="badge builder"></span> StarsBuilder
---
# <span class="badge builder"></span> StarsBuilder

## Constructor

```go
func NewStarsBuilder() *StarsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *StarsBuilder) Build() (Stars, error)
```

### <span class="badge object-method"></span> Resource

```go
func (builder *StarsBuilder) Resource(resource cog.Builder[starsv1alpha1.Resource]) *StarsBuilder
```

### <span class="badge object-method"></span> Resources

```go
func (builder *StarsBuilder) Resources(resource []cog.Builder[starsv1alpha1.Resource]) *StarsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Stars](./object-Stars.md)
