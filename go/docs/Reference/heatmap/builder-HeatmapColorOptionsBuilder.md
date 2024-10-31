---
title: <span class="badge builder"></span> HeatmapColorOptionsBuilder
---
# <span class="badge builder"></span> HeatmapColorOptionsBuilder

## Constructor

```go
func NewHeatmapColorOptionsBuilder() *HeatmapColorOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *HeatmapColorOptionsBuilder) Build() (HeatmapColorOptions, error)
```

### <span class="badge object-method"></span> Exponent

Controls the exponent when scale is set to exponential

```go
func (builder *HeatmapColorOptionsBuilder) Exponent(exponent float32) *HeatmapColorOptionsBuilder
```

### <span class="badge object-method"></span> Fill

Controls the color fill when in opacity mode

```go
func (builder *HeatmapColorOptionsBuilder) Fill(fill string) *HeatmapColorOptionsBuilder
```

### <span class="badge object-method"></span> Max

Sets the maximum value for the color scale

```go
func (builder *HeatmapColorOptionsBuilder) Max(max float32) *HeatmapColorOptionsBuilder
```

### <span class="badge object-method"></span> Min

Sets the minimum value for the color scale

```go
func (builder *HeatmapColorOptionsBuilder) Min(min float32) *HeatmapColorOptionsBuilder
```

### <span class="badge object-method"></span> Mode

Sets the color mode

```go
func (builder *HeatmapColorOptionsBuilder) Mode(mode heatmap.HeatmapColorMode) *HeatmapColorOptionsBuilder
```

### <span class="badge object-method"></span> Reverse

Reverses the color scheme

```go
func (builder *HeatmapColorOptionsBuilder) Reverse(reverse bool) *HeatmapColorOptionsBuilder
```

### <span class="badge object-method"></span> Scale

Controls the color scale

```go
func (builder *HeatmapColorOptionsBuilder) Scale(scale heatmap.HeatmapColorScale) *HeatmapColorOptionsBuilder
```

### <span class="badge object-method"></span> Scheme

Controls the color scheme used

```go
func (builder *HeatmapColorOptionsBuilder) Scheme(scheme string) *HeatmapColorOptionsBuilder
```

### <span class="badge object-method"></span> Steps

Controls the number of color steps

```go
func (builder *HeatmapColorOptionsBuilder) Steps(steps uint64) *HeatmapColorOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [HeatmapColorOptions](./object-HeatmapColorOptions.md)
