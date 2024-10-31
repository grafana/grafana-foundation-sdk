---
title: <span class="badge object-type-struct"></span> BoolOrFloat64
---
# <span class="badge object-type-struct"></span> BoolOrFloat64

## Definition

```go
type BoolOrFloat64 struct {
    Bool *bool `json:"Bool,omitempty"`
    Float64 *float64 `json:"Float64,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `BoolOrFloat64` as JSON.

```go
func (boolOrFloat64 *BoolOrFloat64) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `BoolOrFloat64` from JSON.

```go
func (boolOrFloat64 *BoolOrFloat64) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BoolOrFloat64` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (boolOrFloat64 *BoolOrFloat64) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BoolOrFloat64` objects.

```go
func (boolOrFloat64 *BoolOrFloat64) Equals(other BoolOrFloat64) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BoolOrFloat64` fields for violations and returns them.

```go
func (boolOrFloat64 *BoolOrFloat64) Validate() error
```

