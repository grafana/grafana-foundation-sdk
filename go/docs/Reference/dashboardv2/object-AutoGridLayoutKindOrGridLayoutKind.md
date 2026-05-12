---
title: <span class="badge object-type-struct"></span> AutoGridLayoutKindOrGridLayoutKind
---
# <span class="badge object-type-struct"></span> AutoGridLayoutKindOrGridLayoutKind

## Definition

```go
type AutoGridLayoutKindOrGridLayoutKind struct {
    AutoGridLayoutKind *dashboardv2.AutoGridLayoutKind `json:"AutoGridLayoutKind,omitempty"`
    GridLayoutKind *dashboardv2.GridLayoutKind `json:"GridLayoutKind,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `AutoGridLayoutKindOrGridLayoutKind` as JSON.

```go
func (autoGridLayoutKindOrGridLayoutKind *AutoGridLayoutKindOrGridLayoutKind) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `AutoGridLayoutKindOrGridLayoutKind` from JSON.

```go
func (autoGridLayoutKindOrGridLayoutKind *AutoGridLayoutKindOrGridLayoutKind) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutKindOrGridLayoutKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (autoGridLayoutKindOrGridLayoutKind *AutoGridLayoutKindOrGridLayoutKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AutoGridLayoutKindOrGridLayoutKind` objects.

```go
func (autoGridLayoutKindOrGridLayoutKind *AutoGridLayoutKindOrGridLayoutKind) Equals(other AutoGridLayoutKindOrGridLayoutKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AutoGridLayoutKindOrGridLayoutKind` fields for violations and returns them.

```go
func (autoGridLayoutKindOrGridLayoutKind *AutoGridLayoutKindOrGridLayoutKind) Validate() error
```

