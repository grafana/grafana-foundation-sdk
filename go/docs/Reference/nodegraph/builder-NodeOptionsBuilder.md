---
title: <span class="badge builder"></span> NodeOptionsBuilder
---
# <span class="badge builder"></span> NodeOptionsBuilder

## Constructor

```go
func NewNodeOptionsBuilder() *NodeOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *NodeOptionsBuilder) Build() (NodeOptions, error)
```

### <span class="badge object-method"></span> Arcs

Define which fields are shown as part of the node arc (colored circle around the node).

```go
func (builder *NodeOptionsBuilder) Arcs(arcs []cog.Builder[nodegraph.ArcOption]) *NodeOptionsBuilder
```

### <span class="badge object-method"></span> MainStatUnit

Unit for the main stat to override what ever is set in the data frame.

```go
func (builder *NodeOptionsBuilder) MainStatUnit(mainStatUnit string) *NodeOptionsBuilder
```

### <span class="badge object-method"></span> SecondaryStatUnit

Unit for the secondary stat to override what ever is set in the data frame.

```go
func (builder *NodeOptionsBuilder) SecondaryStatUnit(secondaryStatUnit string) *NodeOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [NodeOptions](./object-NodeOptions.md)
