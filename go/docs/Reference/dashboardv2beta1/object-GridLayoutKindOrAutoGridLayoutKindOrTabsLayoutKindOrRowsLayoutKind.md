---
title: <span class="badge object-type-struct"></span> GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind
---
# <span class="badge object-type-struct"></span> GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind

## Definition

```go
type GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind struct {
    GridLayoutKind *dashboardv2beta1.GridLayoutKind `json:"GridLayoutKind,omitempty"`
    AutoGridLayoutKind *dashboardv2beta1.AutoGridLayoutKind `json:"AutoGridLayoutKind,omitempty"`
    TabsLayoutKind *dashboardv2beta1.TabsLayoutKind `json:"TabsLayoutKind,omitempty"`
    RowsLayoutKind *dashboardv2beta1.RowsLayoutKind `json:"RowsLayoutKind,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` as JSON.

```go
func (gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind *GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` from JSON.

```go
func (gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind *GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind *GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` objects.

```go
func (gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind *GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) Equals(other GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind` fields for violations and returns them.

```go
func (gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind *GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind) Validate() error
```

