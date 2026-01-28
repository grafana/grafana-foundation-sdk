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

### <span class="badge object-method"></span> ColorMode

```go
func (builder *OptionsBuilder) ColorMode(colorMode common.BigValueColorMode) *OptionsBuilder
```

### <span class="badge object-method"></span> GraphMode

```go
func (builder *OptionsBuilder) GraphMode(graphMode common.BigValueGraphMode) *OptionsBuilder
```

### <span class="badge object-method"></span> JustifyMode

```go
func (builder *OptionsBuilder) JustifyMode(justifyMode common.BigValueJustifyMode) *OptionsBuilder
```

### <span class="badge object-method"></span> Orientation

```go
func (builder *OptionsBuilder) Orientation(orientation common.VizOrientation) *OptionsBuilder
```

### <span class="badge object-method"></span> ReduceOptions

```go
func (builder *OptionsBuilder) ReduceOptions(reduceOptions cog.Builder[common.ReduceDataOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> Text

```go
func (builder *OptionsBuilder) Text(text cog.Builder[common.VizTextDisplayOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> TextMode

```go
func (builder *OptionsBuilder) TextMode(textMode common.BigValueTextMode) *OptionsBuilder
```

### <span class="badge object-method"></span> WideLayout

```go
func (builder *OptionsBuilder) WideLayout(wideLayout bool) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
