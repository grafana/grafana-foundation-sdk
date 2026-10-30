---
title: <span class="badge object-type-struct"></span> NamedRoutingTree
---
# <span class="badge object-type-struct"></span> NamedRoutingTree

## Definition

```go
type NamedRoutingTree struct {
    Type alertingv1beta1.NotificationSettingsType `json:"type"`
    RoutingTree string `json:"routingTree"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NamedRoutingTree` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (namedRoutingTree *NamedRoutingTree) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NamedRoutingTree` objects.

```go
func (namedRoutingTree *NamedRoutingTree) Equals(other NamedRoutingTree) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NamedRoutingTree` fields for violations and returns them.

```go
func (namedRoutingTree *NamedRoutingTree) Validate() error
```

## See also

 * <span class="badge builder"></span> [NamedRoutingTreeBuilder](./builder-NamedRoutingTreeBuilder.md)
