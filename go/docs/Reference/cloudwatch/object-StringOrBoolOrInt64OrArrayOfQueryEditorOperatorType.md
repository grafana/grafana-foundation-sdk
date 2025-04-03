---
title: <span class="badge object-type-struct"></span> StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType
---
# <span class="badge object-type-struct"></span> StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType

## Definition

```go
type StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType struct {
    String *string `json:"String,omitempty"`
    Bool *bool `json:"Bool,omitempty"`
    Int64 *int64 `json:"Int64,omitempty"`
    ArrayOfQueryEditorOperatorType []cloudwatch.QueryEditorOperatorType `json:"ArrayOfQueryEditorOperatorType,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` as JSON.

```go
func (stringOrBoolOrInt64OrArrayOfQueryEditorOperatorType *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` from JSON.

```go
func (stringOrBoolOrInt64OrArrayOfQueryEditorOperatorType *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrBoolOrInt64OrArrayOfQueryEditorOperatorType *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` objects.

```go
func (stringOrBoolOrInt64OrArrayOfQueryEditorOperatorType *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) Equals(other StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` fields for violations and returns them.

```go
func (stringOrBoolOrInt64OrArrayOfQueryEditorOperatorType *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) Validate() error
```

