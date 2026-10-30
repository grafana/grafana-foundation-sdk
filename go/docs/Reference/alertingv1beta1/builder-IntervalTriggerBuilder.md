---
title: <span class="badge builder"></span> IntervalTriggerBuilder
---
# <span class="badge builder"></span> IntervalTriggerBuilder

## Constructor

```go
func NewIntervalTriggerBuilder() *IntervalTriggerBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *IntervalTriggerBuilder) Build() (IntervalTrigger, error)
```

### <span class="badge object-method"></span> Interval

```go
func (builder *IntervalTriggerBuilder) Interval(interval alertingv1beta1.PromDuration) *IntervalTriggerBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [IntervalTrigger](./object-IntervalTrigger.md)
