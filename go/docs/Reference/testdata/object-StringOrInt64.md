---
title: <span class="badge object-type-struct"></span> StringOrInt64
---
# <span class="badge object-type-struct"></span> StringOrInt64

## Definition

```go
type StringOrInt64 struct {
    String *string `json:"String,omitempty"`
    Int64 *int64 `json:"Int64,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrInt64` as JSON.

```go
func (stringOrInt64 *StringOrInt64) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrInt64` from JSON.

```go
func (stringOrInt64 *StringOrInt64) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrInt64` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrInt64 *StringOrInt64) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrInt64` objects.

```go
func (stringOrInt64 *StringOrInt64) Equals(other StringOrInt64) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrInt64` fields for violations and returns them.

```go
func (stringOrInt64 *StringOrInt64) Validate() error
```

