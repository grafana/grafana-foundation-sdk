---
title: <span class="badge object-type-struct"></span> PulseWaveQuery
---
# <span class="badge object-type-struct"></span> PulseWaveQuery

## Definition

```go
type PulseWaveQuery struct {
    TimeStep *int64 `json:"timeStep,omitempty"`
    OnCount *int64 `json:"onCount,omitempty"`
    OffCount *int64 `json:"offCount,omitempty"`
    OnValue *float64 `json:"onValue,omitempty"`
    OffValue *float64 `json:"offValue,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PulseWaveQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (pulseWaveQuery *PulseWaveQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PulseWaveQuery` objects.

```go
func (pulseWaveQuery *PulseWaveQuery) Equals(other PulseWaveQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PulseWaveQuery` fields for violations and returns them.

```go
func (pulseWaveQuery *PulseWaveQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [PulseWaveQueryBuilder](./builder-PulseWaveQueryBuilder.md)
