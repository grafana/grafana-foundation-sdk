---
title: <span class="badge object-type-struct"></span> BaseMovingAverageModelSettings
---
# <span class="badge object-type-struct"></span> BaseMovingAverageModelSettings

## Definition

```go
type BaseMovingAverageModelSettings struct {
    Model elasticsearch.MovingAverageModel `json:"model"`
    Window string `json:"window"`
    Predict string `json:"predict"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `BaseMovingAverageModelSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (baseMovingAverageModelSettings *BaseMovingAverageModelSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `BaseMovingAverageModelSettings` objects.

```go
func (baseMovingAverageModelSettings *BaseMovingAverageModelSettings) Equals(other BaseMovingAverageModelSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `BaseMovingAverageModelSettings` fields for violations and returns them.

```go
func (baseMovingAverageModelSettings *BaseMovingAverageModelSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [BaseMovingAverageModelSettingsBuilder](./builder-BaseMovingAverageModelSettingsBuilder.md)
