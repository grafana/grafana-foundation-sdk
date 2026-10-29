---
title: <span class="badge builder"></span> SimplifiedRoutingBuilder
---
# <span class="badge builder"></span> SimplifiedRoutingBuilder

## Constructor

```go
func NewSimplifiedRoutingBuilder() *SimplifiedRoutingBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SimplifiedRoutingBuilder) Build() (SimplifiedRouting, error)
```

### <span class="badge object-method"></span> ActiveTimeIntervals

```go
func (builder *SimplifiedRoutingBuilder) ActiveTimeIntervals(activeTimeIntervals []alertingv0alpha1.TimeIntervalRef) *SimplifiedRoutingBuilder
```

### <span class="badge object-method"></span> GroupBy

```go
func (builder *SimplifiedRoutingBuilder) GroupBy(groupBy []string) *SimplifiedRoutingBuilder
```

### <span class="badge object-method"></span> GroupInterval

```go
func (builder *SimplifiedRoutingBuilder) GroupInterval(groupInterval alertingv0alpha1.PromDuration) *SimplifiedRoutingBuilder
```

### <span class="badge object-method"></span> GroupWait

```go
func (builder *SimplifiedRoutingBuilder) GroupWait(groupWait alertingv0alpha1.PromDuration) *SimplifiedRoutingBuilder
```

### <span class="badge object-method"></span> MuteTimeIntervals

```go
func (builder *SimplifiedRoutingBuilder) MuteTimeIntervals(muteTimeIntervals []alertingv0alpha1.TimeIntervalRef) *SimplifiedRoutingBuilder
```

### <span class="badge object-method"></span> Receiver

```go
func (builder *SimplifiedRoutingBuilder) Receiver(receiver string) *SimplifiedRoutingBuilder
```

### <span class="badge object-method"></span> RepeatInterval

```go
func (builder *SimplifiedRoutingBuilder) RepeatInterval(repeatInterval alertingv0alpha1.PromDuration) *SimplifiedRoutingBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SimplifiedRouting](./object-SimplifiedRouting.md)
