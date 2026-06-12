---
title: <span class="badge builder"></span> ScopesBuilder
---
# <span class="badge builder"></span> ScopesBuilder

## Constructor

```go
func NewScopesBuilder() *ScopesBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ScopesBuilder) Build() (Scopes, error)
```

### <span class="badge object-method"></span> DefaultPath

```go
func (builder *ScopesBuilder) DefaultPath(defaultPath []string) *ScopesBuilder
```

### <span class="badge object-method"></span> Filters

```go
func (builder *ScopesBuilder) Filters(filters []cog.Builder[prometheus.ScopesFilters]) *ScopesBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *ScopesBuilder) Title(title string) *ScopesBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Scopes](./object-Scopes.md)
