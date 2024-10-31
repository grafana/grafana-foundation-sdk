---
title: <span class="badge builder"></span> TableFieldOptionsBuilder
---
# <span class="badge builder"></span> TableFieldOptionsBuilder

## Constructor

```go
func NewTableFieldOptionsBuilder() *TableFieldOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TableFieldOptionsBuilder) Build() (TableFieldOptions, error)
```

### <span class="badge object-method"></span> Align

```go
func (builder *TableFieldOptionsBuilder) Align(align common.FieldTextAlignment) *TableFieldOptionsBuilder
```

### <span class="badge object-method"></span> CellOptions

```go
func (builder *TableFieldOptionsBuilder) CellOptions(cellOptions common.TableCellOptions) *TableFieldOptionsBuilder
```

### <span class="badge object-method"></span> DisplayMode

This field is deprecated in favor of using cellOptions

```go
func (builder *TableFieldOptionsBuilder) DisplayMode(displayMode common.TableCellDisplayMode) *TableFieldOptionsBuilder
```

### <span class="badge object-method"></span> Filterable

```go
func (builder *TableFieldOptionsBuilder) Filterable(filterable bool) *TableFieldOptionsBuilder
```

### <span class="badge object-method"></span> Hidden

?? default is missing or false ??

```go
func (builder *TableFieldOptionsBuilder) Hidden(hidden bool) *TableFieldOptionsBuilder
```

### <span class="badge object-method"></span> HideHeader

Hides any header for a column, usefull for columns that show some static content or buttons.

```go
func (builder *TableFieldOptionsBuilder) HideHeader(hideHeader bool) *TableFieldOptionsBuilder
```

### <span class="badge object-method"></span> Inspect

```go
func (builder *TableFieldOptionsBuilder) Inspect(inspect bool) *TableFieldOptionsBuilder
```

### <span class="badge object-method"></span> MinWidth

```go
func (builder *TableFieldOptionsBuilder) MinWidth(minWidth float64) *TableFieldOptionsBuilder
```

### <span class="badge object-method"></span> Width

```go
func (builder *TableFieldOptionsBuilder) Width(width float64) *TableFieldOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TableFieldOptions](./object-TableFieldOptions.md)
