---
title: <span class="badge builder"></span> MovingAverageHoltWintersModelSettingsBuilder
---
# <span class="badge builder"></span> MovingAverageHoltWintersModelSettingsBuilder

## Constructor

```go
func NewMovingAverageHoltWintersModelSettingsBuilder() *MovingAverageHoltWintersModelSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MovingAverageHoltWintersModelSettingsBuilder) Build() (MovingAverageHoltWintersModelSettings, error)
```

### <span class="badge object-method"></span> Minimize

```go
func (builder *MovingAverageHoltWintersModelSettingsBuilder) Minimize(minimize bool) *MovingAverageHoltWintersModelSettingsBuilder
```

### <span class="badge object-method"></span> Predict

```go
func (builder *MovingAverageHoltWintersModelSettingsBuilder) Predict(predict string) *MovingAverageHoltWintersModelSettingsBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MovingAverageHoltWintersModelSettingsBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchMovingAverageHoltWintersModelSettingsSettings]) *MovingAverageHoltWintersModelSettingsBuilder
```

### <span class="badge object-method"></span> Window

```go
func (builder *MovingAverageHoltWintersModelSettingsBuilder) Window(window string) *MovingAverageHoltWintersModelSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MovingAverageHoltWintersModelSettings](./object-MovingAverageHoltWintersModelSettings.md)
