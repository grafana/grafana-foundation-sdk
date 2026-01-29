---
title: <span class="badge builder"></span> VariableValueSingleBuilder
---
# <span class="badge builder"></span> VariableValueSingleBuilder

## Constructor

```go
func NewVariableValueSingleBuilder() *VariableValueSingleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VariableValueSingleBuilder) Build() (VariableValueSingle, error)
```

### <span class="badge object-method"></span> Bool

```go
func (builder *VariableValueSingleBuilder) Bool(boolArg bool) *VariableValueSingleBuilder
```

### <span class="badge object-method"></span> CustomVariableValue

```go
func (builder *VariableValueSingleBuilder) CustomVariableValue(customVariableValue dashboardv2beta1.CustomVariableValue) *VariableValueSingleBuilder
```

### <span class="badge object-method"></span> Float64

```go
func (builder *VariableValueSingleBuilder) Float64(float64Arg float64) *VariableValueSingleBuilder
```

### <span class="badge object-method"></span> String

```go
func (builder *VariableValueSingleBuilder) String(stringArg string) *VariableValueSingleBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [VariableValueSingle](./object-VariableValueSingle.md)
