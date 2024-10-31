---
title: <span class="badge object-type-struct"></span> MovingAverageLinearModelSettings
---
# <span class="badge object-type-struct"></span> MovingAverageLinearModelSettings

## Definition

```go
type MovingAverageLinearModelSettings struct {
    Model string `json:"model"`
    Window string `json:"window"`
    Predict string `json:"predict"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageLinearModelSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingAverageLinearModelSettings *MovingAverageLinearModelSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingAverageLinearModelSettings` objects.

```go
func (movingAverageLinearModelSettings *MovingAverageLinearModelSettings) Equals(other MovingAverageLinearModelSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingAverageLinearModelSettings` fields for violations and returns them.

```go
func (movingAverageLinearModelSettings *MovingAverageLinearModelSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [MovingAverageLinearModelSettingsBuilder](./builder-MovingAverageLinearModelSettingsBuilder.md)
