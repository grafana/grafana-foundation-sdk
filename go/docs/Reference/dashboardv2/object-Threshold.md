---
title: <span class="badge object-type-struct"></span> Threshold
---
# <span class="badge object-type-struct"></span> Threshold

## Definition

```go
type Threshold struct {
    // Value null means -Infinity
    Value *float64 `json:"value"`
    Color string `json:"color"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Threshold` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (threshold *Threshold) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Threshold` objects.

```go
func (threshold *Threshold) Equals(other Threshold) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Threshold` fields for violations and returns them.

```go
func (threshold *Threshold) Validate() error
```

