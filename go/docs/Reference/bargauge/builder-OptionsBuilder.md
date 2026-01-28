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

### <span class="badge object-method"></span> DisplayMode

```go
func (builder *OptionsBuilder) DisplayMode(displayMode common.BarGaugeDisplayMode) *OptionsBuilder
```

### <span class="badge object-method"></span> MaxVizHeight

```go
func (builder *OptionsBuilder) MaxVizHeight(maxVizHeight uint32) *OptionsBuilder
```

### <span class="badge object-method"></span> MinVizHeight

```go
func (builder *OptionsBuilder) MinVizHeight(minVizHeight uint32) *OptionsBuilder
```

### <span class="badge object-method"></span> MinVizWidth

```go
func (builder *OptionsBuilder) MinVizWidth(minVizWidth uint32) *OptionsBuilder
```

### <span class="badge object-method"></span> NamePlacement

```go
func (builder *OptionsBuilder) NamePlacement(namePlacement common.BarGaugeNamePlacement) *OptionsBuilder
```

### <span class="badge object-method"></span> Orientation

```go
func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder
```

### <span class="badge object-method"></span> ReduceOptions

```go
func (builder *OptionsBuilder) ReduceOptions(reduceOptions cog.Builder[common.ReduceDataOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowUnfilled

```go
func (builder *OptionsBuilder) ShowUnfilled(showUnfilled bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Sizing

```go
func (builder *OptionsBuilder) Sizing(sizing common.BarGaugeSizing) *OptionsBuilder
```

### <span class="badge object-method"></span> Text

```go
func (builder *OptionsBuilder) Text(text cog.Builder[common.VizTextDisplayOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> ValueMode

```go
func (builder *OptionsBuilder) ValueMode(valueMode common.BarGaugeValueMode) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
