---
title: <span class="badge object-type-struct"></span> StringOrBoolOrInt64
---
# <span class="badge object-type-struct"></span> StringOrBoolOrInt64

## Definition

```go
type StringOrBoolOrInt64 struct {
    String *string `json:"String,omitempty"`
    Bool *bool `json:"Bool,omitempty"`
    Int64 *int64 `json:"Int64,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrInt64` as JSON.

```go
func (stringOrBoolOrInt64 *StringOrBoolOrInt64) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrInt64` from JSON.

```go
func (stringOrBoolOrInt64 *StringOrBoolOrInt64) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrInt64` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrBoolOrInt64 *StringOrBoolOrInt64) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrBoolOrInt64` objects.

```go
func (stringOrBoolOrInt64 *StringOrBoolOrInt64) Equals(other StringOrBoolOrInt64) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrBoolOrInt64` fields for violations and returns them.

```go
func (stringOrBoolOrInt64 *StringOrBoolOrInt64) Validate() error
```

