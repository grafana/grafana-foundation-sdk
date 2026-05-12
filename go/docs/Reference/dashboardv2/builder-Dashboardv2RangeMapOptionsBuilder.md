---
title: <span class="badge builder"></span> Dashboardv2RangeMapOptionsBuilder
---
# <span class="badge builder"></span> Dashboardv2RangeMapOptionsBuilder

## Constructor

```go
func NewDashboardv2RangeMapOptionsBuilder() *Dashboardv2RangeMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *Dashboardv2RangeMapOptionsBuilder) Build() (Dashboardv2RangeMapOptions, error)
```

### <span class="badge object-method"></span> From

Min value of the range. It can be null which means -Infinity

```go
func (builder *Dashboardv2RangeMapOptionsBuilder) From(from float64) *Dashboardv2RangeMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value is within the range

```go
func (builder *Dashboardv2RangeMapOptionsBuilder) Result(result cog.Builder[dashboardv2.ValueMappingResult]) *Dashboardv2RangeMapOptionsBuilder
```

### <span class="badge object-method"></span> To

Max value of the range. It can be null which means +Infinity

```go
func (builder *Dashboardv2RangeMapOptionsBuilder) To(to float64) *Dashboardv2RangeMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dashboardv2RangeMapOptions](./object-Dashboardv2RangeMapOptions.md)
