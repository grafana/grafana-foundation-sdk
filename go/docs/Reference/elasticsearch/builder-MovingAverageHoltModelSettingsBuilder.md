---
title: <span class="badge builder"></span> MovingAverageHoltModelSettingsBuilder
---
# <span class="badge builder"></span> MovingAverageHoltModelSettingsBuilder

## Constructor

```go
func NewMovingAverageHoltModelSettingsBuilder() *MovingAverageHoltModelSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MovingAverageHoltModelSettingsBuilder) Build() (MovingAverageHoltModelSettings, error)
```

### <span class="badge object-method"></span> Minimize

```go
func (builder *MovingAverageHoltModelSettingsBuilder) Minimize(minimize bool) *MovingAverageHoltModelSettingsBuilder
```

### <span class="badge object-method"></span> Predict

```go
func (builder *MovingAverageHoltModelSettingsBuilder) Predict(predict string) *MovingAverageHoltModelSettingsBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MovingAverageHoltModelSettingsBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchMovingAverageHoltModelSettingsSettings]) *MovingAverageHoltModelSettingsBuilder
```

### <span class="badge object-method"></span> Window

```go
func (builder *MovingAverageHoltModelSettingsBuilder) Window(window string) *MovingAverageHoltModelSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MovingAverageHoltModelSettings](./object-MovingAverageHoltModelSettings.md)
