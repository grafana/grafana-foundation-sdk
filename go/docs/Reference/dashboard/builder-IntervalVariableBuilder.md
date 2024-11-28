---
title: <span class="badge builder"></span> IntervalVariableBuilder
---
# <span class="badge builder"></span> IntervalVariableBuilder

## Constructor

```go
func NewIntervalVariableBuilder(name string) *IntervalVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *IntervalVariableBuilder) Build() (VariableModel, error)
```

### <span class="badge object-method"></span> AllowCustomValue

Allow custom values to be entered in the variable

```go
func (builder *IntervalVariableBuilder) AllowCustomValue(allowCustomValue bool) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Auto

Dynamically calculates interval by dividing time range by the count specified.

```go
func (builder *IntervalVariableBuilder) Auto(auto bool) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Current

Shows current selected variable text/value on the dashboard

```go
func (builder *IntervalVariableBuilder) Current(current dashboard.VariableOption) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Description

Description of variable. It can be defined but `null`.

```go
func (builder *IntervalVariableBuilder) Description(description string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Hide

Visibility configuration for the variable

```go
func (builder *IntervalVariableBuilder) Hide(hide dashboard.VariableHide) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Label

Optional display name

```go
func (builder *IntervalVariableBuilder) Label(label string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> MinInterval

The minimum threshold below which the step count intervals will not divide the time.

```go
func (builder *IntervalVariableBuilder) MinInterval(autoMin string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Name

Name of variable

```go
func (builder *IntervalVariableBuilder) Name(name string) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Options

Options that can be selected for a variable.

```go
func (builder *IntervalVariableBuilder) Options(options []dashboard.VariableOption) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> StepCount

How many times the current time range should be divided to calculate the value, similar to the Max data points query option.

For example, if the current visible time range is 30 minutes, then the auto interval groups the data into 30 one-minute increments.

```go
func (builder *IntervalVariableBuilder) StepCount(autoCount int32) *IntervalVariableBuilder
```

### <span class="badge object-method"></span> Values

Query used to fetch values for a variable

```go
func (builder *IntervalVariableBuilder) Values(query dashboard.StringOrMap) *IntervalVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [VariableModel](./object-VariableModel.md)
