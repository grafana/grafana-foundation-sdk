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

### <span class="badge object-method"></span> ColWidth

Controls the column width

```go
func (builder *OptionsBuilder) ColWidth(colWidth float64) *OptionsBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> PerPage

Enables pagination when > 0

```go
func (builder *OptionsBuilder) PerPage(perPage float64) *OptionsBuilder
```

### <span class="badge object-method"></span> RowHeight

Set the height of the rows

```go
func (builder *OptionsBuilder) RowHeight(rowHeight float32) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowValue

Show values on the columns

```go
func (builder *OptionsBuilder) ShowValue(showValue common.VisibilityMode) *OptionsBuilder
```

### <span class="badge object-method"></span> Timezone

```go
func (builder *OptionsBuilder) Timezone(timezone []common.TimeZone) *OptionsBuilder
```

### <span class="badge object-method"></span> Tooltip

```go
func (builder *OptionsBuilder) Tooltip(tooltip cog.Builder[common.VizTooltipOptions]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
