---
title: <span class="badge object-type-struct"></span> StringOrArrayOfString
---
# <span class="badge object-type-struct"></span> StringOrArrayOfString

## Definition

```go
type StringOrArrayOfString struct {
    String *string `json:"String,omitempty"`
    ArrayOfString []string `json:"ArrayOfString,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrArrayOfString` as JSON.

```go
func (stringOrArrayOfString *StringOrArrayOfString) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.

```go
func (stringOrArrayOfString *StringOrArrayOfString) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrArrayOfString *StringOrArrayOfString) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrArrayOfString` objects.

```go
func (stringOrArrayOfString *StringOrArrayOfString) Equals(other StringOrArrayOfString) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrArrayOfString` fields for violations and returns them.

```go
func (stringOrArrayOfString *StringOrArrayOfString) Validate() error
```

