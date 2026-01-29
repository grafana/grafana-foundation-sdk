---
title: <span class="badge object-type-struct"></span> Kind
---
# <span class="badge object-type-struct"></span> Kind

--- Common types ---

## Definition

```go
type Kind struct {
    Kind string `json:"kind"`
    Spec any `json:"spec"`
    Metadata any `json:"metadata,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Kind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (kind *Kind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Kind` objects.

```go
func (kind *Kind) Equals(other Kind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Kind` fields for violations and returns them.

```go
func (kind *Kind) Validate() error
```

## See also

 * <span class="badge builder"></span> [KindBuilder](./builder-KindBuilder.md)
