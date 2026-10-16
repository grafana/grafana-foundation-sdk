---
title: <span class="badge object-type-struct"></span> CellKindOrV2PanelKindOrLibraryPanelKind
---
# <span class="badge object-type-struct"></span> CellKindOrV2PanelKindOrLibraryPanelKind

## Definition

```go
type CellKindOrV2PanelKindOrLibraryPanelKind struct {
    CellKind *dashboardv2beta1.CellKind `json:"CellKind,omitempty"`
    V2PanelKind *dashboardv2beta1.V2PanelKind `json:"V2PanelKind,omitempty"`
    LibraryPanelKind *dashboardv2beta1.LibraryPanelKind `json:"LibraryPanelKind,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `CellKindOrV2PanelKindOrLibraryPanelKind` as JSON.

```go
func (cellKindOrV2PanelKindOrLibraryPanelKind *CellKindOrV2PanelKindOrLibraryPanelKind) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `CellKindOrV2PanelKindOrLibraryPanelKind` from JSON.

```go
func (cellKindOrV2PanelKindOrLibraryPanelKind *CellKindOrV2PanelKindOrLibraryPanelKind) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CellKindOrV2PanelKindOrLibraryPanelKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cellKindOrV2PanelKindOrLibraryPanelKind *CellKindOrV2PanelKindOrLibraryPanelKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CellKindOrV2PanelKindOrLibraryPanelKind` objects.

```go
func (cellKindOrV2PanelKindOrLibraryPanelKind *CellKindOrV2PanelKindOrLibraryPanelKind) Equals(other CellKindOrV2PanelKindOrLibraryPanelKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CellKindOrV2PanelKindOrLibraryPanelKind` fields for violations and returns them.

```go
func (cellKindOrV2PanelKindOrLibraryPanelKind *CellKindOrV2PanelKindOrLibraryPanelKind) Validate() error
```

