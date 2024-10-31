---
title: <span class="badge builder"></span> MuteTimingBuilder
---
# <span class="badge builder"></span> MuteTimingBuilder

## Constructor

```go
func NewMuteTimingBuilder() *MuteTimingBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MuteTimingBuilder) Build() (MuteTiming, error)
```

### <span class="badge object-method"></span> Name

```go
func (builder *MuteTimingBuilder) Name(name string) *MuteTimingBuilder
```

### <span class="badge object-method"></span> TimeIntervals

```go
func (builder *MuteTimingBuilder) TimeIntervals(timeIntervals []cog.Builder[alerting.TimeInterval]) *MuteTimingBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MuteTiming](./object-MuteTiming.md)
