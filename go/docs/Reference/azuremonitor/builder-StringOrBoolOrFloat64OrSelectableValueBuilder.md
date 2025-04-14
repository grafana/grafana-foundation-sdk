---
title: <span class="badge builder"></span> StringOrBoolOrFloat64OrSelectableValueBuilder
---
# <span class="badge builder"></span> StringOrBoolOrFloat64OrSelectableValueBuilder

## Constructor

```go
func NewStringOrBoolOrFloat64OrSelectableValueBuilder() *StringOrBoolOrFloat64OrSelectableValueBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) Build() (StringOrBoolOrFloat64OrSelectableValue, error)
```

### <span class="badge object-method"></span> Bool

```go
func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) Bool(boolArg bool) *StringOrBoolOrFloat64OrSelectableValueBuilder
```

### <span class="badge object-method"></span> Float64

```go
func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) Float64(float64Arg float64) *StringOrBoolOrFloat64OrSelectableValueBuilder
```

### <span class="badge object-method"></span> SelectableValue

```go
func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) SelectableValue(selectableValue cog.Builder[azuremonitor.SelectableValue]) *StringOrBoolOrFloat64OrSelectableValueBuilder
```

### <span class="badge object-method"></span> String

```go
func (builder *StringOrBoolOrFloat64OrSelectableValueBuilder) String(stringArg string) *StringOrBoolOrFloat64OrSelectableValueBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [StringOrBoolOrFloat64OrSelectableValue](./object-StringOrBoolOrFloat64OrSelectableValue.md)
