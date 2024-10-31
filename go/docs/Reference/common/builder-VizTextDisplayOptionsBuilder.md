---
title: <span class="badge builder"></span> VizTextDisplayOptionsBuilder
---
# <span class="badge builder"></span> VizTextDisplayOptionsBuilder

## Constructor

```go
func NewVizTextDisplayOptionsBuilder() *VizTextDisplayOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *VizTextDisplayOptionsBuilder) Build() (VizTextDisplayOptions, error)
```

### <span class="badge object-method"></span> TitleSize

Explicit title text size

```go
func (builder *VizTextDisplayOptionsBuilder) TitleSize(titleSize float64) *VizTextDisplayOptionsBuilder
```

### <span class="badge object-method"></span> ValueSize

Explicit value text size

```go
func (builder *VizTextDisplayOptionsBuilder) ValueSize(valueSize float64) *VizTextDisplayOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VizTextDisplayOptions](./object-VizTextDisplayOptions.md)
