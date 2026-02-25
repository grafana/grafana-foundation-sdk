---
title: <span class="badge builder"></span> LibraryPanelBuilder
---
# <span class="badge builder"></span> LibraryPanelBuilder

## Constructor

```go
func NewLibraryPanelBuilder() *LibraryPanelBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LibraryPanelBuilder) Build() (LibraryPanelKind, error)
```

### <span class="badge object-method"></span> Id

Panel ID for the library panel in the dashboard

```go
func (builder *LibraryPanelBuilder) Id(id float64) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> LibraryPanel

```go
func (builder *LibraryPanelBuilder) LibraryPanel(libraryPanel cog.Builder[dashboardv2beta1.LibraryPanelRef]) *LibraryPanelBuilder
```

### <span class="badge object-method"></span> Title

Title for the library panel in the dashboard

```go
func (builder *LibraryPanelBuilder) Title(title string) *LibraryPanelBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LibraryPanelKind](./object-LibraryPanelKind.md)
