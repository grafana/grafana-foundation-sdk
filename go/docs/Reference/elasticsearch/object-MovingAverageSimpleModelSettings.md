---
title: <span class="badge object-type-struct"></span> MovingAverageSimpleModelSettings
---
# <span class="badge object-type-struct"></span> MovingAverageSimpleModelSettings

## Definition

```go
type MovingAverageSimpleModelSettings struct {
    Model string `json:"model"`
    Window string `json:"window"`
    Predict string `json:"predict"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageSimpleModelSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingAverageSimpleModelSettings *MovingAverageSimpleModelSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingAverageSimpleModelSettings` objects.

```go
func (movingAverageSimpleModelSettings *MovingAverageSimpleModelSettings) Equals(other MovingAverageSimpleModelSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingAverageSimpleModelSettings` fields for violations and returns them.

```go
func (movingAverageSimpleModelSettings *MovingAverageSimpleModelSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [MovingAverageSimpleModelSettingsBuilder](./builder-MovingAverageSimpleModelSettingsBuilder.md)
