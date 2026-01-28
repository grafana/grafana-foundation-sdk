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

### <span class="badge object-method"></span> CellHeight

Controls the height of the rows

```go
func (builder *OptionsBuilder) CellHeight(cellHeight common.TableCellHeight) *OptionsBuilder
```

### <span class="badge object-method"></span> Footer

Controls footer options

```go
func (builder *OptionsBuilder) Footer(footer cog.Builder[common.TableFooterOptions]) *OptionsBuilder
```

### <span class="badge object-method"></span> FrameIndex

Represents the index of the selected frame

```go
func (builder *OptionsBuilder) FrameIndex(frameIndex float64) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowHeader

Controls whether the panel should show the header

```go
func (builder *OptionsBuilder) ShowHeader(showHeader bool) *OptionsBuilder
```

### <span class="badge object-method"></span> ShowTypeIcons

Controls whether the header should show icons for the column types

```go
func (builder *OptionsBuilder) ShowTypeIcons(showTypeIcons bool) *OptionsBuilder
```

### <span class="badge object-method"></span> SortBy

Used to control row sorting

```go
func (builder *OptionsBuilder) SortBy(sortBy []cog.Builder[common.TableSortByFieldState]) *OptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Options](./object-Options.md)
