---
title: <span class="badge object-type-struct"></span> NodeOptions
---
# <span class="badge object-type-struct"></span> NodeOptions

## Definition

```go
type NodeOptions struct {
    // Unit for the main stat to override what ever is set in the data frame.
    MainStatUnit *string `json:"mainStatUnit,omitempty"`
    // Unit for the secondary stat to override what ever is set in the data frame.
    SecondaryStatUnit *string `json:"secondaryStatUnit,omitempty"`
    // Define which fields are shown as part of the node arc (colored circle around the node).
    Arcs []nodegraph.ArcOption `json:"arcs,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `NodeOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (nodeOptions *NodeOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `NodeOptions` objects.

```go
func (nodeOptions *NodeOptions) Equals(other NodeOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `NodeOptions` fields for violations and returns them.

```go
func (nodeOptions *NodeOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [NodeOptionsBuilder](./builder-NodeOptionsBuilder.md)
