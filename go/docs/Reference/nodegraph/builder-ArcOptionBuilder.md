---
title: <span class="badge builder"></span> ArcOptionBuilder
---
# <span class="badge builder"></span> ArcOptionBuilder

## Constructor

```go
func NewArcOptionBuilder() *ArcOptionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ArcOptionBuilder) Build() (ArcOption, error)
```

### <span class="badge object-method"></span> Color

The color of the arc.

```go
func (builder *ArcOptionBuilder) Color(color string) *ArcOptionBuilder
```

### <span class="badge object-method"></span> Field

Field from which to get the value. Values should be less than 1, representing fraction of a circle.

```go
func (builder *ArcOptionBuilder) Field(field string) *ArcOptionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ArcOption](./object-ArcOption.md)
