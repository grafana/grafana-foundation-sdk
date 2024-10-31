---
title: <span class="badge builder"></span> TimePickerBuilder
---
# <span class="badge builder"></span> TimePickerBuilder

## Constructor

```go
func NewTimePickerBuilder() *TimePickerBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TimePickerBuilder) Build() (TimePickerConfig, error)
```

### <span class="badge object-method"></span> Hidden

Whether timepicker is visible or not.

```go
func (builder *TimePickerBuilder) Hidden(hidden bool) *TimePickerBuilder
```

### <span class="badge object-method"></span> RefreshIntervals

Interval options available in the refresh picker dropdown.

```go
func (builder *TimePickerBuilder) RefreshIntervals(refreshIntervals []string) *TimePickerBuilder
```

### <span class="badge object-method"></span> TimeOptions

Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.

```go
func (builder *TimePickerBuilder) TimeOptions(timeOptions []string) *TimePickerBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimePickerConfig](./object-TimePickerConfig.md)
