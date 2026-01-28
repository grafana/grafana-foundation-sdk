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

### <span class="badge object-method"></span> Orientation

```go
func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder
```

### <span class="badge object-method"></span> ReduceOptions

```go
func (builder *OptionsBuilder) ReduceOptions(reduceOptions cog.Builder[common.ReduceDataOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowThresholdLabels

```go
func (builder *OptionsBuilder) ShowThresholdLabels(showThresholdLabels bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowThresholdMarkers

```go
func (builder *OptionsBuilder) ShowThresholdMarkers(showThresholdMarkers bool) *OptionsBuilder
```

### <span class="badge object-method"></span> Text

```go
func (builder *OptionsBuilder) Text(text cog.Builder[common.VizTextDisplayOptions]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
