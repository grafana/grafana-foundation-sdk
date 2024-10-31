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

### <span class="badge object-method"></span> Name

```go
func (builder *TimeIntervalBuilder) Name(name string) *TimeIntervalBuilder
```

### <span class="badge object-method"></span> TimeIntervals

```go
func (builder *TimeIntervalBuilder) TimeIntervals(timeIntervals []cog.Builder[alerting.TimeIntervalItem]) *TimeIntervalBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TimeInterval](./object-TimeInterval.md)
