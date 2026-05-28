---
title: <span class="badge object-type-struct"></span> TimeRangeOption
---
# <span class="badge object-type-struct"></span> TimeRangeOption

## Definition

```go
type TimeRangeOption struct {
    Display string `json:"display"`
    From string `json:"from"`
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeRangeOption` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timeRangeOption *TimeRangeOption) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimeRangeOption` objects.

```go
func (timeRangeOption *TimeRangeOption) Equals(other TimeRangeOption) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimeRangeOption` fields for violations and returns them.

```go
func (timeRangeOption *TimeRangeOption) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimeRangeOptionBuilder](./builder-TimeRangeOptionBuilder.md)
