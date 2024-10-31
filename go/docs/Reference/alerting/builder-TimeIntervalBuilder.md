---
title: <span class="badge builder"></span> TimeIntervalBuilder
---
# <span class="badge builder"></span> TimeIntervalBuilder

## Constructor

```go
func NewTimeIntervalBuilder() *TimeIntervalBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TimeIntervalBuilder) Build() (TimeInterval, error)
```

### <span class="badge object-method"></span> DaysOfMonth

```go
func (builder *TimeIntervalBuilder) DaysOfMonth(daysOfMonth []string) *TimeIntervalBuilder
```

### <span class="badge object-method"></span> Location

```go
func (builder *TimeIntervalBuilder) Location(location string) *TimeIntervalBuilder
```

### <span class="badge object-method"></span> Months

```go
func (builder *TimeIntervalBuilder) Months(months []string) *TimeIntervalBuilder
```

### <span class="badge object-method"></span> Times

```go
func (builder *TimeIntervalBuilder) Times(times []cog.Builder[alerting.TimeRange]) *TimeIntervalBuilder
```

### <span class="badge object-method"></span> Weekdays

```go
func (builder *TimeIntervalBuilder) Weekdays(weekdays []string) *TimeIntervalBuilder
```

### <span class="badge object-method"></span> Years

```go
func (builder *TimeIntervalBuilder) Years(years []string) *TimeIntervalBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimeInterval](./object-TimeInterval.md)
