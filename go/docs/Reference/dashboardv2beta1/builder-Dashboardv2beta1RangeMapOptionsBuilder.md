---
title: <span class="badge builder"></span> Dashboardv2beta1RangeMapOptionsBuilder
---
# <span class="badge builder"></span> Dashboardv2beta1RangeMapOptionsBuilder

## Constructor

```go
func NewDashboardv2beta1RangeMapOptionsBuilder() *Dashboardv2beta1RangeMapOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *Dashboardv2beta1RangeMapOptionsBuilder) Build() (Dashboardv2beta1RangeMapOptions, error)
```

### <span class="badge object-method"></span> From

Min value of the range. It can be null which means -Infinity

```go
func (builder *Dashboardv2beta1RangeMapOptionsBuilder) From(from float64) *Dashboardv2beta1RangeMapOptionsBuilder
```

### <span class="badge object-method"></span> Result

Config to apply when the value is within the range

```go
func (builder *Dashboardv2beta1RangeMapOptionsBuilder) Result(result cog.Builder[dashboardv2beta1.ValueMappingResult]) *Dashboardv2beta1RangeMapOptionsBuilder
```

### <span class="badge object-method"></span> To

Max value of the range. It can be null which means +Infinity

```go
func (builder *Dashboardv2beta1RangeMapOptionsBuilder) To(to float64) *Dashboardv2beta1RangeMapOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dashboardv2beta1RangeMapOptions](./object-Dashboardv2beta1RangeMapOptions.md)
