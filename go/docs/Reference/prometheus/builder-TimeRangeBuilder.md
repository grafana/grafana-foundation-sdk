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

From is the start time of the query.

```go
func (builder *TimeRangeBuilder) From(from string) *TimeRangeBuilder
```

### <span class="badge object-method"></span> To

To is the end time of the query.

```go
func (builder *TimeRangeBuilder) To(to string) *TimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimeRange](./object-TimeRange.md)
