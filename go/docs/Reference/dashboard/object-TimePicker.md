---
title: <span class="badge object-type-struct"></span> TimePicker
---
# <span class="badge object-type-struct"></span> TimePicker

## Definition

```go
type TimePicker struct {
    // Whether timepicker is visible or not.
    Hidden bool `json:"hidden"`
    // Interval options available in the refresh picker dropdown.
    RefreshIntervals []string `json:"refresh_intervals"`
    // Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
    Collapse bool `json:"collapse"`
    // Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
    Enable bool `json:"enable"`
    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    TimeOptions []string `json:"time_options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimePicker` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timePicker *TimePicker) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimePicker` objects.

```go
func (timePicker *TimePicker) Equals(other TimePicker) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimePicker` fields for violations and returns them.

```go
func (timePicker *TimePicker) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimePickerBuilder](./builder-TimePickerBuilder.md)
