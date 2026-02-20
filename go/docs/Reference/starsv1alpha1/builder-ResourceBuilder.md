---
title: <span class="badge builder"></span> ResourceBuilder
---
# <span class="badge builder"></span> ResourceBuilder

## Constructor

```go
func NewResourceBuilder() *ResourceBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ResourceBuilder) Build() (Resource, error)
```

### <span class="badge object-method"></span> Group

```go
func (builder *ResourceBuilder) Group(group string) *ResourceBuilder
```

### <span class="badge object-method"></span> Kind

```go
func (builder *ResourceBuilder) Kind(kind string) *ResourceBuilder
```

### <span class="badge object-method"></span> Names

The set of resources

+listType=set

```go
func (builder *ResourceBuilder) Names(names []string) *ResourceBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Resource](./object-Resource.md)
