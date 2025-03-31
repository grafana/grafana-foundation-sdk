---
title: <span class="badge builder"></span> TimeRangeBuilder
---
# <span class="badge builder"></span> TimeRangeBuilder

## Constructor

```go
func NewTimeRangeBuilder() *TimeRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TimeRangeBuilder) Build() (TimeRange, error)
```

### <span class="badge object-method"></span> From

```go
func (builder *TimeRangeBuilder) From(from time.Time) *TimeRangeBuilder
```

### <span class="badge object-method"></span> To

```go
func (builder *TimeRangeBuilder) To(to time.Time) *TimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimeRange](./object-TimeRange.md)
