---
title: <span class="badge builder"></span> VizTooltipOptionsBuilder
---
# <span class="badge builder"></span> VizTooltipOptionsBuilder

## Constructor

```go
func NewVizTooltipOptionsBuilder() *VizTooltipOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VizTooltipOptionsBuilder) Build() (VizTooltipOptions, error)
```

### <span class="badge object-method"></span> Mode

```go
func (builder *VizTooltipOptionsBuilder) Mode(mode common.TooltipDisplayMode) *VizTooltipOptionsBuilder
```

### <span class="badge object-method"></span> Sort

```go
func (builder *VizTooltipOptionsBuilder) Sort(sort common.SortOrder) *VizTooltipOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VizTooltipOptions](./object-VizTooltipOptions.md)
