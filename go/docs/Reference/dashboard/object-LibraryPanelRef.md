---
title: <span class="badge object-type-struct"></span> LibraryPanelRef
---
# <span class="badge object-type-struct"></span> LibraryPanelRef

A library panel is a reusable panel that you can use in any dashboard.

When you make a change to a library panel, that change propagates to all instances of where the panel is used.

Library panels streamline reuse of panels across multiple dashboards.

## Definition

```go
type LibraryPanelRef struct {
    // Library panel name
    Name string `json:"name"`
    // Library panel uid
    Uid string `json:"uid"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LibraryPanelRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (libraryPanelRef *LibraryPanelRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `LibraryPanelRef` objects.

```go
func (libraryPanelRef *LibraryPanelRef) Equals(other LibraryPanelRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `LibraryPanelRef` fields for violations and returns them.

```go
func (libraryPanelRef *LibraryPanelRef) Validate() error
```

