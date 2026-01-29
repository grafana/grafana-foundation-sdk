---
title: <span class="badge builder"></span> LibraryPanelKindSpecBuilder
---
# <span class="badge builder"></span> LibraryPanelKindSpecBuilder

## Constructor

```go
func NewLibraryPanelKindSpecBuilder() *LibraryPanelKindSpecBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LibraryPanelKindSpecBuilder) Build() (LibraryPanelKindSpec, error)
```

### <span class="badge object-method"></span> Id

Panel ID for the library panel in the dashboard

```go
func (builder *LibraryPanelKindSpecBuilder) Id(id float64) *LibraryPanelKindSpecBuilder
```

### <span class="badge object-method"></span> LibraryPanel

```go
func (builder *LibraryPanelKindSpecBuilder) LibraryPanel(libraryPanel cog.Builder[dashboardv2beta1.LibraryPanelRef]) *LibraryPanelKindSpecBuilder
```

### <span class="badge object-method"></span> Title

Title for the library panel in the dashboard

```go
func (builder *LibraryPanelKindSpecBuilder) Title(title string) *LibraryPanelKindSpecBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LibraryPanelKindSpec](./object-LibraryPanelKindSpec.md)
