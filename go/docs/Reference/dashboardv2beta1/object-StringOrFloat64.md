---
title: <span class="badge object-type-struct"></span> StringOrFloat64
---
# <span class="badge object-type-struct"></span> StringOrFloat64

## Definition

```go
type StringOrFloat64 struct {
    String *string `json:"String,omitempty"`
    Float64 *float64 `json:"Float64,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `StringOrFloat64` as JSON.

```go
func (stringOrFloat64 *StringOrFloat64) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrFloat64` from JSON.

```go
func (stringOrFloat64 *StringOrFloat64) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrFloat64` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (stringOrFloat64 *StringOrFloat64) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StringOrFloat64` objects.

```go
func (stringOrFloat64 *StringOrFloat64) Equals(other StringOrFloat64) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StringOrFloat64` fields for violations and returns them.

```go
func (stringOrFloat64 *StringOrFloat64) Validate() error
```

