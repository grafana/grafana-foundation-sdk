---
title: <span class="badge object-type-struct"></span> MovingAverageEWMAModelSettings
---
# <span class="badge object-type-struct"></span> MovingAverageEWMAModelSettings

## Definition

```go
type MovingAverageEWMAModelSettings struct {
    Model string `json:"model"`
    Settings *elasticsearch.ElasticsearchMovingAverageEWMAModelSettingsSettings `json:"settings,omitempty"`
    Window string `json:"window"`
    Minimize bool `json:"minimize"`
    Predict string `json:"predict"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageEWMAModelSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingAverageEWMAModelSettings *MovingAverageEWMAModelSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingAverageEWMAModelSettings` objects.

```go
func (movingAverageEWMAModelSettings *MovingAverageEWMAModelSettings) Equals(other MovingAverageEWMAModelSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingAverageEWMAModelSettings` fields for violations and returns them.

```go
func (movingAverageEWMAModelSettings *MovingAverageEWMAModelSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [MovingAverageEWMAModelSettingsBuilder](./builder-MovingAverageEWMAModelSettingsBuilder.md)
