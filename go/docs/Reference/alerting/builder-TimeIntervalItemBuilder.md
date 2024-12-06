---
title: <span class="badge builder"></span> TimeIntervalItemBuilder
---
# <span class="badge builder"></span> TimeIntervalItemBuilder

## Constructor

```go
func NewTimeIntervalItemBuilder() *TimeIntervalItemBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TimeIntervalItemBuilder) Build() (TimeIntervalItem, error)
```

### <span class="badge object-method"></span> DaysOfMonth

```go
func (builder *TimeIntervalItemBuilder) DaysOfMonth(daysOfMonth []string) *TimeIntervalItemBuilder
```

### <span class="badge object-method"></span> Location

```go
func (builder *TimeIntervalItemBuilder) Location(location string) *TimeIntervalItemBuilder
```

### <span class="badge object-method"></span> Months

```go
func (builder *TimeIntervalItemBuilder) Months(months []string) *TimeIntervalItemBuilder
```

### <span class="badge object-method"></span> Times

```go
func (builder *TimeIntervalItemBuilder) Times(times []cog.Builder[alerting.TimeIntervalTimeRange]) *TimeIntervalItemBuilder
```

### <span class="badge object-method"></span> Weekdays

```go
func (builder *TimeIntervalItemBuilder) Weekdays(weekdays []string) *TimeIntervalItemBuilder
```

### <span class="badge object-method"></span> Years

```go
func (builder *TimeIntervalItemBuilder) Years(years []string) *TimeIntervalItemBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimeIntervalItem](./object-TimeIntervalItem.md)
