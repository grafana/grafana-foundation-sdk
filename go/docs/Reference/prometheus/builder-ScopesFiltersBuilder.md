---
title: <span class="badge builder"></span> ScopesFiltersBuilder
---
# <span class="badge builder"></span> ScopesFiltersBuilder

## Constructor

```go
func NewScopesFiltersBuilder() *ScopesFiltersBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ScopesFiltersBuilder) Build() (ScopesFilters, error)
```

### <span class="badge object-method"></span> Key

```go
func (builder *ScopesFiltersBuilder) Key(key string) *ScopesFiltersBuilder
```

### <span class="badge object-method"></span> Operator

```go
func (builder *ScopesFiltersBuilder) Operator(operator string) *ScopesFiltersBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *ScopesFiltersBuilder) Value(value string) *ScopesFiltersBuilder
```

### <span class="badge object-method"></span> Values

```go
func (builder *ScopesFiltersBuilder) Values(values []string) *ScopesFiltersBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ScopesFilters](./object-ScopesFilters.md)
