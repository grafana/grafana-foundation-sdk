---
title: <span class="badge builder"></span> FilterBuilder
---
# <span class="badge builder"></span> FilterBuilder

## Constructor

```go
func NewFilterBuilder() *FilterBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FilterBuilder) Build() (Filter, error)
```

### <span class="badge object-method"></span> Label

```go
func (builder *FilterBuilder) Label(label string) *FilterBuilder
```

### <span class="badge object-method"></span> Query

```go
func (builder *FilterBuilder) Query(query string) *FilterBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Filter](./object-Filter.md)
