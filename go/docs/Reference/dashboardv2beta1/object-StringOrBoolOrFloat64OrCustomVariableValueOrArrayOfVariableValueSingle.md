---
title: <span class="badge object-type-struct"></span> StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle
---
# <span class="badge object-type-struct"></span> StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle

## Definition

```go
type StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle struct {
    String *string `json:"String,omitempty"`
    Bool *bool `json:"Bool,omitempty"`
    Float64 *float64 `json:"Float64,omitempty"`
    CustomVariableValue *dashboardv2beta1.CustomVariableValue `json:"CustomVariableValue,omitempty"`
    ArrayOfVariableValueSingle []dashboardv2beta1.VariableValueSingle `json:"ArrayOfVariableValueSingle,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` as JSON.

```go
func (stringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` from JSON.

```go
func (stringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` objects.

```go
func (stringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) Equals(other StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle` fields for violations and returns them.

```go
func (stringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle *StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingle) Validate() error
```

## See also

 * <span class="badge builder"></span> [StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder](./builder-StringOrBoolOrFloat64OrCustomVariableValueOrArrayOfVariableValueSingleBuilder.md)
