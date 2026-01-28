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

### <span class="badge object-method"></span> AlignValue

Controls value alignment on the timelines

```go
func (builder *OptionsBuilder) AlignValue(alignValue common.TimelineValueAlignment) *OptionsBuilder
```

### <span class="badge object-method"></span> Legend

```go
func (builder *OptionsBuilder) Legend(legend cog.Builder[common.VizLegendOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> MergeValues

Merge equal consecutive values

```go
func (builder *OptionsBuilder) MergeValues(mergeValues bool) *OptionsBuilder
```

### <span class="badge object-method"></span> RowHeight

Controls the row height

```go
func (builder *OptionsBuilder) RowHeight(rowHeight float64) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowValue

Show timeline values on chart

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
