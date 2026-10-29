---
title: <span class="badge object-type-struct"></span> IntervalTrigger
---
# <span class="badge object-type-struct"></span> IntervalTrigger

## Definition

```go
type IntervalTrigger struct {
    Interval alertingv0alpha1.PromDuration `json:"interval"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `IntervalTrigger` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (intervalTrigger *IntervalTrigger) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `IntervalTrigger` objects.

```go
func (intervalTrigger *IntervalTrigger) Equals(other IntervalTrigger) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `IntervalTrigger` fields for violations and returns them.

```go
func (intervalTrigger *IntervalTrigger) Validate() error
```

## See also

 * <span class="badge builder"></span> [IntervalTriggerBuilder](./builder-IntervalTriggerBuilder.md)
