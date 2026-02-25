---
title: <span class="badge object-type-struct"></span> LibraryPanelKindSpec
---
# <span class="badge object-type-struct"></span> LibraryPanelKindSpec

## Definition

```go
type LibraryPanelKindSpec struct {
    // Panel ID for the library panel in the dashboard
    Id float64 `json:"id"`
    // Title for the library panel in the dashboard
    Title string `json:"title"`
    LibraryPanel dashboardv2beta1.LibraryPanelRef `json:"libraryPanel"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryPanelKindSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (libraryPanelKindSpec *LibraryPanelKindSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LibraryPanelKindSpec` objects.

```go
func (libraryPanelKindSpec *LibraryPanelKindSpec) Equals(other LibraryPanelKindSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LibraryPanelKindSpec` fields for violations and returns them.

```go
func (libraryPanelKindSpec *LibraryPanelKindSpec) Validate() error
```

