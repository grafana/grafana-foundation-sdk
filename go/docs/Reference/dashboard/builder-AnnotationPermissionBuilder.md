---
title: <span class="badge builder"></span> AnnotationPermissionBuilder
---
# <span class="badge builder"></span> AnnotationPermissionBuilder

## Constructor

```go
func NewAnnotationPermissionBuilder() *AnnotationPermissionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AnnotationPermissionBuilder) Build() (AnnotationPermission, error)
```

### <span class="badge object-method"></span> Dashboard

```go
func (builder *AnnotationPermissionBuilder) Dashboard(dashboard cog.Builder[dashboard.AnnotationActions]) *AnnotationPermissionBuilder
```

### <span class="badge object-method"></span> Organization

```go
func (builder *AnnotationPermissionBuilder) Organization(organization cog.Builder[dashboard.AnnotationActions]) *AnnotationPermissionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AnnotationPermission](./object-AnnotationPermission.md)
