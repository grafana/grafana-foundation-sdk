---
title: <span class="badge builder"></span> SearchResultResourceBuilder
---
# <span class="badge builder"></span> SearchResultResourceBuilder

#SearchResultResource is the full identity of a hit.

## Constructor

```go
func NewSearchResultResourceBuilder() *SearchResultResourceBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SearchResultResourceBuilder) Build() (SearchResultResource, error)
```

### <span class="badge object-method"></span> Group

```go
func (builder *SearchResultResourceBuilder) Group(group string) *SearchResultResourceBuilder
```

### <span class="badge object-method"></span> Kind

```go
func (builder *SearchResultResourceBuilder) Kind(kind string) *SearchResultResourceBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *SearchResultResourceBuilder) Name(name string) *SearchResultResourceBuilder
```

### <span class="badge object-method"></span> Resource

```go
func (builder *SearchResultResourceBuilder) Resource(resource string) *SearchResultResourceBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SearchResultResource](./object-SearchResultResource.md)
