---
title: <span class="badge builder"></span> VariableValueBuilder
---
# <span class="badge builder"></span> VariableValueBuilder

## Constructor

```go
func NewVariableValueBuilder() *VariableValueBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VariableValueBuilder) Build() (VariableValue, error)
```

### <span class="badge object-method"></span> ArrayOfVariableValueSingle

```go
func (builder *VariableValueBuilder) ArrayOfVariableValueSingle(arrayOfVariableValueSingle []cog.Builder[dashboardv2beta1.VariableValueSingle]) *VariableValueBuilder
```

### <span class="badge object-method"></span> Bool

```go
func (builder *VariableValueBuilder) Bool(boolArg bool) *VariableValueBuilder
```

### <span class="badge object-method"></span> CustomVariableValue

```go
func (builder *VariableValueBuilder) CustomVariableValue(customVariableValue dashboardv2beta1.CustomVariableValue) *VariableValueBuilder
```

### <span class="badge object-method"></span> Float64

```go
func (builder *VariableValueBuilder) Float64(float64Arg float64) *VariableValueBuilder
```

### <span class="badge object-method"></span> String

```go
func (builder *VariableValueBuilder) String(stringArg string) *VariableValueBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [VariableValue](./object-VariableValue.md)
