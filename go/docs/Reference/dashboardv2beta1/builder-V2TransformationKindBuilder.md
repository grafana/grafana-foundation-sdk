---
title: <span class="badge builder"></span> V2TransformationKindBuilder
---
# <span class="badge builder"></span> V2TransformationKindBuilder

Dashboard v2 shape: the transformation ID moved from `kind` to `group`.

## Constructor

```go
func NewV2TransformationKindBuilder() *V2TransformationKindBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *V2TransformationKindBuilder) Build() (V2TransformationKind, error)
```

### <span class="badge object-method"></span> Group

The group is the transformation ID

```go
func (builder *V2TransformationKindBuilder) Group(group string) *V2TransformationKindBuilder
```

### <span class="badge object-method"></span> Spec

```go
func (builder *V2TransformationKindBuilder) Spec(spec cog.Builder[dashboardv2beta1.V2TransformationSpec]) *V2TransformationKindBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [V2TransformationKind](./object-V2TransformationKind.md)
