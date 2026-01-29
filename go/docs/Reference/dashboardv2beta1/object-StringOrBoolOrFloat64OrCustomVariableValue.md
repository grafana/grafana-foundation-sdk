---
title: <span class="badge object-type-struct"></span> StringOrBoolOrFloat64OrCustomVariableValue
---
# <span class="badge object-type-struct"></span> StringOrBoolOrFloat64OrCustomVariableValue

## Definition

```go
type StringOrBoolOrFloat64OrCustomVariableValue struct {
    String *string `json:"String,omitempty"`
    Bool *bool `json:"Bool,omitempty"`
    Float64 *float64 `json:"Float64,omitempty"`
    CustomVariableValue *dashboardv2beta1.CustomVariableValue `json:"CustomVariableValue,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrFloat64OrCustomVariableValue` as JSON.

```go
func (stringOrBoolOrFloat64OrCustomVariableValue *StringOrBoolOrFloat64OrCustomVariableValue) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrCustomVariableValue` from JSON.

```go
func (stringOrBoolOrFloat64OrCustomVariableValue *StringOrBoolOrFloat64OrCustomVariableValue) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrCustomVariableValue` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrBoolOrFloat64OrCustomVariableValue *StringOrBoolOrFloat64OrCustomVariableValue) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrBoolOrFloat64OrCustomVariableValue` objects.

```go
func (stringOrBoolOrFloat64OrCustomVariableValue *StringOrBoolOrFloat64OrCustomVariableValue) Equals(other StringOrBoolOrFloat64OrCustomVariableValue) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrBoolOrFloat64OrCustomVariableValue` fields for violations and returns them.

```go
func (stringOrBoolOrFloat64OrCustomVariableValue *StringOrBoolOrFloat64OrCustomVariableValue) Validate() error
```

## See also

 * <span class="badge builder"></span> [StringOrBoolOrFloat64OrCustomVariableValueBuilder](./builder-StringOrBoolOrFloat64OrCustomVariableValueBuilder.md)
