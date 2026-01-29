---
title: <span class="badge object-type-struct"></span> PanelKindOrLibraryPanelKind
---
# <span class="badge object-type-struct"></span> PanelKindOrLibraryPanelKind

## Definition

```go
type PanelKindOrLibraryPanelKind struct {
    PanelKind *dashboardv2beta1.PanelKind `json:"PanelKind,omitempty"`
    LibraryPanelKind *dashboardv2beta1.LibraryPanelKind `json:"LibraryPanelKind,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `PanelKindOrLibraryPanelKind` as JSON.

```go
func (panelKindOrLibraryPanelKind *PanelKindOrLibraryPanelKind) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `PanelKindOrLibraryPanelKind` from JSON.

```go
func (panelKindOrLibraryPanelKind *PanelKindOrLibraryPanelKind) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelKindOrLibraryPanelKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (panelKindOrLibraryPanelKind *PanelKindOrLibraryPanelKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PanelKindOrLibraryPanelKind` objects.

```go
func (panelKindOrLibraryPanelKind *PanelKindOrLibraryPanelKind) Equals(other PanelKindOrLibraryPanelKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PanelKindOrLibraryPanelKind` fields for violations and returns them.

```go
func (panelKindOrLibraryPanelKind *PanelKindOrLibraryPanelKind) Validate() error
```

