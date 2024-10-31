---
title: <span class="badge object-type-struct"></span> TimePickerConfig
---
# <span class="badge object-type-struct"></span> TimePickerConfig

Time picker configuration

It defines the default config for the time picker and the refresh picker for the specific dashboard.

## Definition

```go
type TimePickerConfig struct {
    // Whether timepicker is visible or not.
    Hidden bool `json:"hidden"`
    // Interval options available in the refresh picker dropdown.
    RefreshIntervals []string `json:"refresh_intervals"`
    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    TimeOptions []string `json:"time_options"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimePickerConfig` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (timePickerConfig *TimePickerConfig) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TimePickerConfig` objects.

```go
func (timePickerConfig *TimePickerConfig) Equals(other TimePickerConfig) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TimePickerConfig` fields for violations and returns them.

```go
func (timePickerConfig *TimePickerConfig) Validate() error
```

## See also

 * <span class="badge builder"></span> [TimePickerBuilder](./builder-TimePickerBuilder.md)
