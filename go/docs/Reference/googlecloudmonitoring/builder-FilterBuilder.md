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

### <span class="badge object-method"></span> Condition

Filter condition.

```go
func (builder *FilterBuilder) Condition(condition string) *FilterBuilder
```

### <span class="badge object-method"></span> Key

Filter key.

```go
func (builder *FilterBuilder) Key(key string) *FilterBuilder
```

### <span class="badge object-method"></span> Operator

Filter operator.

```go
func (builder *FilterBuilder) Operator(operator string) *FilterBuilder
```

### <span class="badge object-method"></span> Value

Filter value.

```go
func (builder *FilterBuilder) Value(value string) *FilterBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Filter](./object-Filter.md)
