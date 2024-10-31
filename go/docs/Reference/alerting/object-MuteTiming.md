---
title: <span class="badge object-type-struct"></span> MuteTiming
---
# <span class="badge object-type-struct"></span> MuteTiming

## Definition

```go
type MuteTiming struct {
    Name *string `json:"name,omitempty"`
    TimeIntervals []alerting.TimeInterval `json:"time_intervals,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MuteTiming` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (muteTiming *MuteTiming) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MuteTiming` objects.

```go
func (muteTiming *MuteTiming) Equals(other MuteTiming) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MuteTiming` fields for violations and returns them.

```go
func (muteTiming *MuteTiming) Validate() error
```

## See also

 * <span class="badge builder"></span> [MuteTimingBuilder](./builder-MuteTimingBuilder.md)
