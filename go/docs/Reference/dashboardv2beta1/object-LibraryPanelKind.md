---
title: <span class="badge object-type-struct"></span> LibraryPanelKind
---
# <span class="badge object-type-struct"></span> LibraryPanelKind

## Definition

```go
type LibraryPanelKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.LibraryPanelKindSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryPanelKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (libraryPanelKind *LibraryPanelKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LibraryPanelKind` objects.

```go
func (libraryPanelKind *LibraryPanelKind) Equals(other LibraryPanelKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LibraryPanelKind` fields for violations and returns them.

```go
func (libraryPanelKind *LibraryPanelKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [LibraryPanelBuilder](./builder-LibraryPanelBuilder.md)
