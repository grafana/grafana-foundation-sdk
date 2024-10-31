---
title: <span class="badge object-type-struct"></span> MovingAverageHoltModelSettings
---
# <span class="badge object-type-struct"></span> MovingAverageHoltModelSettings

## Definition

```go
type MovingAverageHoltModelSettings struct {
    Model string `json:"model"`
    Settings elasticsearch.ElasticsearchMovingAverageHoltModelSettingsSettings `json:"settings"`
    Window string `json:"window"`
    Minimize bool `json:"minimize"`
    Predict string `json:"predict"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageHoltModelSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingAverageHoltModelSettings *MovingAverageHoltModelSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingAverageHoltModelSettings` objects.

```go
func (movingAverageHoltModelSettings *MovingAverageHoltModelSettings) Equals(other MovingAverageHoltModelSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingAverageHoltModelSettings` fields for violations and returns them.

```go
func (movingAverageHoltModelSettings *MovingAverageHoltModelSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [MovingAverageHoltModelSettingsBuilder](./builder-MovingAverageHoltModelSettingsBuilder.md)
