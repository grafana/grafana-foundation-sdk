---
title: <span class="badge builder"></span> StringOrBoolOrFloat64OrCustomVariableValueBuilder
---
# <span class="badge builder"></span> StringOrBoolOrFloat64OrCustomVariableValueBuilder

## Constructor

```go
func NewStringOrBoolOrFloat64OrCustomVariableValueBuilder() *StringOrBoolOrFloat64OrCustomVariableValueBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) Build() (StringOrBoolOrFloat64OrCustomVariableValue, error)
```

### <span class="badge object-method"></span> Bool

```go
func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) Bool(boolArg bool) *StringOrBoolOrFloat64OrCustomVariableValueBuilder
```

### <span class="badge object-method"></span> CustomVariableValue

```go
func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) CustomVariableValue(customVariableValue dashboardv2beta1.CustomVariableValue) *StringOrBoolOrFloat64OrCustomVariableValueBuilder
```

### <span class="badge object-method"></span> Float64

```go
func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) Float64(float64Arg float64) *StringOrBoolOrFloat64OrCustomVariableValueBuilder
```

### <span class="badge object-method"></span> String

```go
func (builder *StringOrBoolOrFloat64OrCustomVariableValueBuilder) String(stringArg string) *StringOrBoolOrFloat64OrCustomVariableValueBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [StringOrBoolOrFloat64OrCustomVariableValue](./object-StringOrBoolOrFloat64OrCustomVariableValue.md)
