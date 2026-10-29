---
title: <span class="badge builder"></span> RelativeTimeRangeBuilder
---
# <span class="badge builder"></span> RelativeTimeRangeBuilder

## Constructor

```go
func NewRelativeTimeRangeBuilder() *RelativeTimeRangeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *RelativeTimeRangeBuilder) Build() (RelativeTimeRange, error)
```

### <span class="badge object-method"></span> From

```go
func (builder *RelativeTimeRangeBuilder) From(from alertingv0alpha1.PromDurationWMillis) *RelativeTimeRangeBuilder
```

### <span class="badge object-method"></span> To

```go
func (builder *RelativeTimeRangeBuilder) To(to alertingv0alpha1.PromDurationWMillis) *RelativeTimeRangeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [RelativeTimeRange](./object-RelativeTimeRange.md)
