---
title: <span class="badge object-type-struct"></span> Key
---
# <span class="badge object-type-struct"></span> Key

## Definition

```go
type Key struct {
    Tick float64 `json:"tick"`
    Type string `json:"type"`
    Uid *string `json:"uid,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Key` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (key *Key) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Key` objects.

```go
func (key *Key) Equals(other Key) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Key` fields for violations and returns them.

```go
func (key *Key) Validate() error
```

## See also

 * <span class="badge builder"></span> [KeyBuilder](./builder-KeyBuilder.md)
