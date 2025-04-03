---
title: <span class="badge builder"></span> FieldColorBuilder
---
# <span class="badge builder"></span> FieldColorBuilder

## Constructor

```go
func NewFieldColorBuilder() *FieldColorBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FieldColorBuilder) Build() (FieldColor, error)
```

### <span class="badge object-method"></span> FixedColor

The fixed color value for fixed or shades color modes.

```go
func (builder *FieldColorBuilder) FixedColor(fixedColor string) *FieldColorBuilder
```

### <span class="badge object-method"></span> Mode

The main color scheme mode.

```go
func (builder *FieldColorBuilder) Mode(mode dashboard.FieldColorModeId) *FieldColorBuilder
```

### <span class="badge object-method"></span> SeriesBy

Some visualizations need to know how to assign a series color from by value color schemes.

```go
func (builder *FieldColorBuilder) SeriesBy(seriesBy dashboard.FieldColorSeriesByMode) *FieldColorBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FieldColor](./object-FieldColor.md)
