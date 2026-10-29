---
title: <span class="badge object-type-struct"></span> SimplifiedRoutingOrNamedRoutingTree
---
# <span class="badge object-type-struct"></span> SimplifiedRoutingOrNamedRoutingTree

## Definition

```go
type SimplifiedRoutingOrNamedRoutingTree struct {
    SimplifiedRouting *alertingv0alpha1.SimplifiedRouting `json:"SimplifiedRouting,omitempty"`
    NamedRoutingTree *alertingv0alpha1.NamedRoutingTree `json:"NamedRoutingTree,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `SimplifiedRoutingOrNamedRoutingTree` as JSON.

```go
func (simplifiedRoutingOrNamedRoutingTree *SimplifiedRoutingOrNamedRoutingTree) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `SimplifiedRoutingOrNamedRoutingTree` from JSON.

```go
func (simplifiedRoutingOrNamedRoutingTree *SimplifiedRoutingOrNamedRoutingTree) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SimplifiedRoutingOrNamedRoutingTree` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (simplifiedRoutingOrNamedRoutingTree *SimplifiedRoutingOrNamedRoutingTree) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SimplifiedRoutingOrNamedRoutingTree` objects.

```go
func (simplifiedRoutingOrNamedRoutingTree *SimplifiedRoutingOrNamedRoutingTree) Equals(other SimplifiedRoutingOrNamedRoutingTree) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SimplifiedRoutingOrNamedRoutingTree` fields for violations and returns them.

```go
func (simplifiedRoutingOrNamedRoutingTree *SimplifiedRoutingOrNamedRoutingTree) Validate() error
```

