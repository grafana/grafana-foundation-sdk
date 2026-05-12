---
title: <span class="badge object-type-struct"></span> CustomVariableValue
---
# <span class="badge object-type-struct"></span> CustomVariableValue

Custom variable value

## Definition

```go
type CustomVariableValue struct {
    // The format name or function used in the expression
    Formatter *string `json:"formatter"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CustomVariableValue` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (customVariableValue *CustomVariableValue) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CustomVariableValue` objects.

```go
func (customVariableValue *CustomVariableValue) Equals(other CustomVariableValue) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CustomVariableValue` fields for violations and returns them.

```go
func (customVariableValue *CustomVariableValue) Validate() error
```

