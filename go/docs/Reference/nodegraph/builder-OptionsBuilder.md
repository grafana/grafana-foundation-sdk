---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```go
func NewOptionsBuilder() *OptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *OptionsBuilder) Build() (Options, error)
```

### <span class="badge object-method"></span> Edges

```go
func (builder *OptionsBuilder) Edges(edges cog.Builder[nodegraph.EdgeOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Nodes

```go
func (builder *OptionsBuilder) Nodes(nodes cog.Builder[nodegraph.NodeOptions]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
