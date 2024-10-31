---
title: <span class="badge object-type-struct"></span> BoolOrUint32
---
# <span class="badge object-type-struct"></span> BoolOrUint32

## Definition

```go
type BoolOrUint32 struct {
    Bool *bool `json:"Bool,omitempty"`
    Uint32 *uint32 `json:"Uint32,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `BoolOrUint32` as JSON.

```go
func (boolOrUint32 *BoolOrUint32) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `BoolOrUint32` from JSON.

```go
func (boolOrUint32 *BoolOrUint32) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BoolOrUint32` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (boolOrUint32 *BoolOrUint32) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BoolOrUint32` objects.

```go
func (boolOrUint32 *BoolOrUint32) Equals(other BoolOrUint32) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BoolOrUint32` fields for violations and returns them.

```go
func (boolOrUint32 *BoolOrUint32) Validate() error
```

