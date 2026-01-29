---
title: <span class="badge builder"></span> TargetBuilder
---
# <span class="badge builder"></span> TargetBuilder

## Constructor

```go
func NewTargetBuilder() *TargetBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TargetBuilder) Build() (PanelQueryKind, error)
```

### <span class="badge object-method"></span> Hidden

```go
func (builder *TargetBuilder) Hidden(hidden bool) *TargetBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *TargetBuilder) Query(query cog.Builder[dashboardv2beta1.DataQueryKind]) *TargetBuilder
```

### <span class="badge object-method"></span> RefId

```go
func (builder *TargetBuilder) RefId(refId string) *TargetBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PanelQueryKind](./object-PanelQueryKind.md)
