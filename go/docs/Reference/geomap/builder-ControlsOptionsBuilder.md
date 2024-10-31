---
title: <span class="badge builder"></span> ControlsOptionsBuilder
---
# <span class="badge builder"></span> ControlsOptionsBuilder

## Constructor

```go
func NewControlsOptionsBuilder() *ControlsOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ControlsOptionsBuilder) Build() (ControlsOptions, error)
```

### <span class="badge object-method"></span> MouseWheelZoom

let the mouse wheel zoom

```go
func (builder *ControlsOptionsBuilder) MouseWheelZoom(mouseWheelZoom bool) *ControlsOptionsBuilder
```

### <span class="badge object-method"></span> ShowAttribution

Lower right

```go
func (builder *ControlsOptionsBuilder) ShowAttribution(showAttribution bool) *ControlsOptionsBuilder
```

### <span class="badge object-method"></span> ShowDebug

Show debug

```go
func (builder *ControlsOptionsBuilder) ShowDebug(showDebug bool) *ControlsOptionsBuilder
```

### <span class="badge object-method"></span> ShowMeasure

Show measure

```go
func (builder *ControlsOptionsBuilder) ShowMeasure(showMeasure bool) *ControlsOptionsBuilder
```

### <span class="badge object-method"></span> ShowScale

Scale options

```go
func (builder *ControlsOptionsBuilder) ShowScale(showScale bool) *ControlsOptionsBuilder
```

### <span class="badge object-method"></span> ShowZoom

Zoom (upper left)

```go
func (builder *ControlsOptionsBuilder) ShowZoom(showZoom bool) *ControlsOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ControlsOptions](./object-ControlsOptions.md)
