---
title: <span class="badge builder"></span> FieldConfigBuilder
---
# <span class="badge builder"></span> FieldConfigBuilder

## Constructor

```go
func NewFieldConfigBuilder() *FieldConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FieldConfigBuilder) Build() (FieldConfig, error)
```

### <span class="badge object-method"></span> Align

```go
func (builder *FieldConfigBuilder) Align(align common.FieldTextAlignment) *FieldConfigBuilder
```

### <span class="badge object-method"></span> CellOptions

```go
func (builder *FieldConfigBuilder) CellOptions(cellOptions common.TableCellOptions) *FieldConfigBuilder
```

### <span class="badge object-method"></span> DisplayMode

This field is deprecated in favor of using cellOptions

```go
func (builder *FieldConfigBuilder) DisplayMode(displayMode common.TableCellDisplayMode) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Filterable

```go
func (builder *FieldConfigBuilder) Filterable(filterable bool) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Hidden

?? default is missing or false ??

```go
func (builder *FieldConfigBuilder) Hidden(hidden bool) *FieldConfigBuilder
```

### <span class="badge object-method"></span> HideHeader

Hides any header for a column, useful for columns that show some static content or buttons.

```go
func (builder *FieldConfigBuilder) HideHeader(hideHeader bool) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Inspect

```go
func (builder *FieldConfigBuilder) Inspect(inspect bool) *FieldConfigBuilder
```

### <span class="badge object-method"></span> MinWidth

```go
func (builder *FieldConfigBuilder) MinWidth(minWidth float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Width

```go
func (builder *FieldConfigBuilder) Width(width float64) *FieldConfigBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [FieldConfig](./object-FieldConfig.md)
