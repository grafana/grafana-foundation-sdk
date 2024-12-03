---
title: <span class="badge object-type-struct"></span> ElasticsearchMovingAverageHoltModelSettingsSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchMovingAverageHoltModelSettingsSettings

## Definition

```go
type ElasticsearchMovingAverageHoltModelSettingsSettings struct {
    Alpha *string `json:"alpha,omitempty"`
    Beta *string `json:"beta,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMovingAverageHoltModelSettingsSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchMovingAverageHoltModelSettingsSettings *ElasticsearchMovingAverageHoltModelSettingsSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchMovingAverageHoltModelSettingsSettings` objects.

```go
func (elasticsearchMovingAverageHoltModelSettingsSettings *ElasticsearchMovingAverageHoltModelSettingsSettings) Equals(other ElasticsearchMovingAverageHoltModelSettingsSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchMovingAverageHoltModelSettingsSettings` fields for violations and returns them.

```go
func (elasticsearchMovingAverageHoltModelSettingsSettings *ElasticsearchMovingAverageHoltModelSettingsSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder](./builder-ElasticsearchMovingAverageHoltModelSettingsSettingsBuilder.md)