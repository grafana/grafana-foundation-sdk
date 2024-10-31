---
title: <span class="badge object-type-struct"></span> ElasticsearchMovingAverageEWMAModelSettingsSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchMovingAverageEWMAModelSettingsSettings

## Definition

```go
type ElasticsearchMovingAverageEWMAModelSettingsSettings struct {
    Alpha *string `json:"alpha,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMovingAverageEWMAModelSettingsSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchMovingAverageEWMAModelSettingsSettings *ElasticsearchMovingAverageEWMAModelSettingsSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchMovingAverageEWMAModelSettingsSettings` objects.

```go
func (elasticsearchMovingAverageEWMAModelSettingsSettings *ElasticsearchMovingAverageEWMAModelSettingsSettings) Equals(other ElasticsearchMovingAverageEWMAModelSettingsSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchMovingAverageEWMAModelSettingsSettings` fields for violations and returns them.

```go
func (elasticsearchMovingAverageEWMAModelSettingsSettings *ElasticsearchMovingAverageEWMAModelSettingsSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder](./builder-ElasticsearchMovingAverageEWMAModelSettingsSettingsBuilder.md)
