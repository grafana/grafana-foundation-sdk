---
title: <span class="badge object-type-struct"></span> StringOrBoolOrFloat64OrSelectableValue
---
# <span class="badge object-type-struct"></span> StringOrBoolOrFloat64OrSelectableValue

## Definition

```go
type StringOrBoolOrFloat64OrSelectableValue struct {
    String *string `json:"String,omitempty"`
    Bool *bool `json:"Bool,omitempty"`
    Float64 *float64 `json:"Float64,omitempty"`
    SelectableValue *azuremonitor.SelectableValue `json:"SelectableValue,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrFloat64OrSelectableValue` as JSON.

```go
func (stringOrBoolOrFloat64OrSelectableValue *StringOrBoolOrFloat64OrSelectableValue) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrSelectableValue` from JSON.

```go
func (stringOrBoolOrFloat64OrSelectableValue *StringOrBoolOrFloat64OrSelectableValue) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrFloat64OrSelectableValue` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrBoolOrFloat64OrSelectableValue *StringOrBoolOrFloat64OrSelectableValue) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrBoolOrFloat64OrSelectableValue` objects.

```go
func (stringOrBoolOrFloat64OrSelectableValue *StringOrBoolOrFloat64OrSelectableValue) Equals(other StringOrBoolOrFloat64OrSelectableValue) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrBoolOrFloat64OrSelectableValue` fields for violations and returns them.

```go
func (stringOrBoolOrFloat64OrSelectableValue *StringOrBoolOrFloat64OrSelectableValue) Validate() error
```

## See also

 * <span class="badge builder"></span> [StringOrBoolOrFloat64OrSelectableValueBuilder](./builder-StringOrBoolOrFloat64OrSelectableValueBuilder.md)
