---
title: <span class="badge object-type-struct"></span> MovingAverageHoltWintersModelSettings
---
# <span class="badge object-type-struct"></span> MovingAverageHoltWintersModelSettings

## Definition

```go
type MovingAverageHoltWintersModelSettings struct {
    Model string `json:"model"`
    Settings elasticsearch.ElasticsearchMovingAverageHoltWintersModelSettingsSettings `json:"settings"`
    Window string `json:"window"`
    Minimize bool `json:"minimize"`
    Predict string `json:"predict"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageHoltWintersModelSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingAverageHoltWintersModelSettings *MovingAverageHoltWintersModelSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingAverageHoltWintersModelSettings` objects.

```go
func (movingAverageHoltWintersModelSettings *MovingAverageHoltWintersModelSettings) Equals(other MovingAverageHoltWintersModelSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingAverageHoltWintersModelSettings` fields for violations and returns them.

```go
func (movingAverageHoltWintersModelSettings *MovingAverageHoltWintersModelSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [MovingAverageHoltWintersModelSettingsBuilder](./builder-MovingAverageHoltWintersModelSettingsBuilder.md)
