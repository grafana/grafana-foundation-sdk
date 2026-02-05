---
title: <span class="badge object-type-struct"></span> IntervalVariableKind
---
# <span class="badge object-type-struct"></span> IntervalVariableKind

Interval variable kind

## Definition

```go
type IntervalVariableKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.IntervalVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `IntervalVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (intervalVariableKind *IntervalVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `IntervalVariableKind` objects.

```go
func (intervalVariableKind *IntervalVariableKind) Equals(other IntervalVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `IntervalVariableKind` fields for violations and returns them.

```go
func (intervalVariableKind *IntervalVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [IntervalVariableBuilder](./builder-IntervalVariableBuilder.md)
