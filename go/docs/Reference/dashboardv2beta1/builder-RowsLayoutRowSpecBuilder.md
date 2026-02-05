---
title: <span class="badge builder"></span> RowsLayoutRowSpecBuilder
---
# <span class="badge builder"></span> RowsLayoutRowSpecBuilder

## Constructor

```go
func NewRowsLayoutRowSpecBuilder() *RowsLayoutRowSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RowsLayoutRowSpecBuilder) Build() (RowsLayoutRowSpec, error)
```

### <span class="badge object-method"></span> Collapse

```go
func (builder *RowsLayoutRowSpecBuilder) Collapse(collapse bool) *RowsLayoutRowSpecBuilder
```

### <span class="badge object-method"></span> ConditionalRendering

```go
func (builder *RowsLayoutRowSpecBuilder) ConditionalRendering(conditionalRendering cog.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) *RowsLayoutRowSpecBuilder
```

### <span class="badge object-method"></span> FillScreen

```go
func (builder *RowsLayoutRowSpecBuilder) FillScreen(fillScreen bool) *RowsLayoutRowSpecBuilder
```

### <span class="badge object-method"></span> HideHeader

```go
func (builder *RowsLayoutRowSpecBuilder) HideHeader(hideHeader bool) *RowsLayoutRowSpecBuilder
```

### <span class="badge object-method"></span> Layout

```go
func (builder *RowsLayoutRowSpecBuilder) Layout(layout dashboardv2beta1.GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) *RowsLayoutRowSpecBuilder
```

### <span class="badge object-method"></span> Repeat

```go
func (builder *RowsLayoutRowSpecBuilder) Repeat(repeat cog.Builder[dashboardv2beta1.RowRepeatOptions]) *RowsLayoutRowSpecBuilder
```

### <span class="badge object-method"></span> Title

```go
func (builder *RowsLayoutRowSpecBuilder) Title(title string) *RowsLayoutRowSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RowsLayoutRowSpec](./object-RowsLayoutRowSpec.md)
