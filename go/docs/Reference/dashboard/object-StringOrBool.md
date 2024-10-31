---
title: <span class="badge object-type-struct"></span> StringOrBool
---
# <span class="badge object-type-struct"></span> StringOrBool

## Definition

```go
type StringOrBool struct {
    String *string `json:"String,omitempty"`
    Bool *bool `json:"Bool,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBool` as JSON.

```go
func (stringOrBool *StringOrBool) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBool` from JSON.

```go
func (stringOrBool *StringOrBool) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBool` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrBool *StringOrBool) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrBool` objects.

```go
func (stringOrBool *StringOrBool) Equals(other StringOrBool) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrBool` fields for violations and returns them.

```go
func (stringOrBool *StringOrBool) Validate() error
```

