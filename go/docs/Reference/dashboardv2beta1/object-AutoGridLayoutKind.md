---
title: <span class="badge object-type-struct"></span> AutoGridLayoutKind
---
# <span class="badge object-type-struct"></span> AutoGridLayoutKind

## Definition

```go
type AutoGridLayoutKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.AutoGridLayoutSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (autoGridLayoutKind *AutoGridLayoutKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AutoGridLayoutKind` objects.

```go
func (autoGridLayoutKind *AutoGridLayoutKind) Equals(other AutoGridLayoutKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AutoGridLayoutKind` fields for violations and returns them.

```go
func (autoGridLayoutKind *AutoGridLayoutKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [AutoGridBuilder](./builder-AutoGridBuilder.md)
