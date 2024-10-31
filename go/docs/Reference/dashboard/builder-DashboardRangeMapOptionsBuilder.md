---
title: <span class="badge builder"></span> DashboardRangeMapOptionsBuilder
---
# <span class="badge builder"></span> DashboardRangeMapOptionsBuilder

## Constructor

```go
func NewDashboardRangeMapOptionsBuilder() *DashboardRangeMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DashboardRangeMapOptionsBuilder) Build() (DashboardRangeMapOptions, error)
```

### <span class="badge object-method"></span> From

Min value of the range. It can be null which means -Infinity

```go
func (builder *DashboardRangeMapOptionsBuilder) From(from float64) *DashboardRangeMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value is within the range

```go
func (builder *DashboardRangeMapOptionsBuilder) Result(result cog.Builder[dashboard.ValueMappingResult]) *DashboardRangeMapOptionsBuilder
```

### <span class="badge object-method"></span> To

Max value of the range. It can be null which means +Infinity

```go
func (builder *DashboardRangeMapOptionsBuilder) To(to float64) *DashboardRangeMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DashboardRangeMapOptions](./object-DashboardRangeMapOptions.md)
