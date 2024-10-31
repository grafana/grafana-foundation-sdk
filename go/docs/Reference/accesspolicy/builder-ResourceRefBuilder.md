---
title: <span class="badge builder"></span> ResourceRefBuilder
---
# <span class="badge builder"></span> ResourceRefBuilder

## Constructor

```go
func NewResourceRefBuilder() *ResourceRefBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ResourceRefBuilder) Build() (ResourceRef, error)
```

### <span class="badge object-method"></span> Kind

```go
func (builder *ResourceRefBuilder) Kind(kind string) *ResourceRefBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *ResourceRefBuilder) Name(name string) *ResourceRefBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ResourceRef](./object-ResourceRef.md)
