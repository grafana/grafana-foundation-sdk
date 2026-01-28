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

### <span class="badge object-method"></span> Legend

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

```go
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> XField

Name of the x field to use (defaults to first number)

```go
func (builder *OptionsBuilder) XField(xField string) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
