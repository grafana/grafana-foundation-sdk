---
title: <span class="badge builder"></span> AdhocFiltersBuilder
---
# <span class="badge builder"></span> AdhocFiltersBuilder

## Constructor

```go
func NewAdhocFiltersBuilder() *AdhocFiltersBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AdhocFiltersBuilder) Build() (AdhocFilters, error)
```

### <span class="badge object-method"></span> Key

```go
func (builder *AdhocFiltersBuilder) Key(key string) *AdhocFiltersBuilder
```

### <span class="badge object-method"></span> Operator

```go
func (builder *AdhocFiltersBuilder) Operator(operator string) *AdhocFiltersBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *AdhocFiltersBuilder) Value(value string) *AdhocFiltersBuilder
```

### <span class="badge object-method"></span> Values

```go
func (builder *AdhocFiltersBuilder) Values(values []string) *AdhocFiltersBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AdhocFilters](./object-AdhocFilters.md)
