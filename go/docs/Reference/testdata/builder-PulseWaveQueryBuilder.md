---
title: <span class="badge builder"></span> PulseWaveQueryBuilder
---
# <span class="badge builder"></span> PulseWaveQueryBuilder

## Constructor

```go
func NewPulseWaveQueryBuilder() *PulseWaveQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PulseWaveQueryBuilder) Build() (PulseWaveQuery, error)
```

### <span class="badge object-method"></span> OffCount

```go
func (builder *PulseWaveQueryBuilder) OffCount(offCount int64) *PulseWaveQueryBuilder
```

### <span class="badge object-method"></span> OffValue

```go
func (builder *PulseWaveQueryBuilder) OffValue(offValue float64) *PulseWaveQueryBuilder
```

### <span class="badge object-method"></span> OnCount

```go
func (builder *PulseWaveQueryBuilder) OnCount(onCount int64) *PulseWaveQueryBuilder
```

### <span class="badge object-method"></span> OnValue

```go
func (builder *PulseWaveQueryBuilder) OnValue(onValue float64) *PulseWaveQueryBuilder
```

### <span class="badge object-method"></span> TimeStep

```go
func (builder *PulseWaveQueryBuilder) TimeStep(timeStep int64) *PulseWaveQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PulseWaveQuery](./object-PulseWaveQuery.md)
