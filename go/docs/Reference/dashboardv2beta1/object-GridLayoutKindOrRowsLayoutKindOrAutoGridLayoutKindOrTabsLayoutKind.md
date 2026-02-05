---
title: <span class="badge object-type-struct"></span> GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind
---
# <span class="badge object-type-struct"></span> GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind

## Definition

```go
type GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind struct {
    GridLayoutKind *dashboardv2beta1.GridLayoutKind `json:"GridLayoutKind,omitempty"`
    RowsLayoutKind *dashboardv2beta1.RowsLayoutKind `json:"RowsLayoutKind,omitempty"`
    AutoGridLayoutKind *dashboardv2beta1.AutoGridLayoutKind `json:"AutoGridLayoutKind,omitempty"`
    TabsLayoutKind *dashboardv2beta1.TabsLayoutKind `json:"TabsLayoutKind,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` as JSON.

```go
func (gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind *GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` from JSON.

```go
func (gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind *GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind *GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` objects.

```go
func (gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind *GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) Equals(other GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind` fields for violations and returns them.

```go
func (gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind *GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind) Validate() error
```

