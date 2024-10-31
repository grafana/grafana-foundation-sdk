---
title: <span class="badge builder"></span> BaseMovingAverageModelSettingsBuilder
---
# <span class="badge builder"></span> BaseMovingAverageModelSettingsBuilder

## Constructor

```go
func NewBaseMovingAverageModelSettingsBuilder() *BaseMovingAverageModelSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BaseMovingAverageModelSettingsBuilder) Build() (BaseMovingAverageModelSettings, error)
```

### <span class="badge object-method"></span> Model

```go
func (builder *BaseMovingAverageModelSettingsBuilder) Model(model elasticsearch.MovingAverageModel) *BaseMovingAverageModelSettingsBuilder
```

### <span class="badge object-method"></span> Predict

```go
func (builder *BaseMovingAverageModelSettingsBuilder) Predict(predict string) *BaseMovingAverageModelSettingsBuilder
```

### <span class="badge object-method"></span> Window

```go
func (builder *BaseMovingAverageModelSettingsBuilder) Window(window string) *BaseMovingAverageModelSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BaseMovingAverageModelSettings](./object-BaseMovingAverageModelSettings.md)
