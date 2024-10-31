---
title: <span class="badge builder"></span> MovingAverageEWMAModelSettingsBuilder
---
# <span class="badge builder"></span> MovingAverageEWMAModelSettingsBuilder

## Constructor

```go
func NewMovingAverageEWMAModelSettingsBuilder() *MovingAverageEWMAModelSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MovingAverageEWMAModelSettingsBuilder) Build() (MovingAverageEWMAModelSettings, error)
```

### <span class="badge object-method"></span> Minimize

```go
func (builder *MovingAverageEWMAModelSettingsBuilder) Minimize(minimize bool) *MovingAverageEWMAModelSettingsBuilder
```

### <span class="badge object-method"></span> Predict

```go
func (builder *MovingAverageEWMAModelSettingsBuilder) Predict(predict string) *MovingAverageEWMAModelSettingsBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MovingAverageEWMAModelSettingsBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchMovingAverageEWMAModelSettingsSettings]) *MovingAverageEWMAModelSettingsBuilder
```

### <span class="badge object-method"></span> Window

```go
func (builder *MovingAverageEWMAModelSettingsBuilder) Window(window string) *MovingAverageEWMAModelSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MovingAverageEWMAModelSettings](./object-MovingAverageEWMAModelSettings.md)
