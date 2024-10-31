---
title: <span class="badge object-type-struct"></span> ElasticsearchMovingAverageHoltWintersModelSettingsSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchMovingAverageHoltWintersModelSettingsSettings

## Definition

```go
type ElasticsearchMovingAverageHoltWintersModelSettingsSettings struct {
    Alpha *string `json:"alpha,omitempty"`
    Beta *string `json:"beta,omitempty"`
    Gamma *string `json:"gamma,omitempty"`
    Period *string `json:"period,omitempty"`
    Pad *bool `json:"pad,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMovingAverageHoltWintersModelSettingsSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchMovingAverageHoltWintersModelSettingsSettings *ElasticsearchMovingAverageHoltWintersModelSettingsSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchMovingAverageHoltWintersModelSettingsSettings` objects.

```go
func (elasticsearchMovingAverageHoltWintersModelSettingsSettings *ElasticsearchMovingAverageHoltWintersModelSettingsSettings) Equals(other ElasticsearchMovingAverageHoltWintersModelSettingsSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchMovingAverageHoltWintersModelSettingsSettings` fields for violations and returns them.

```go
func (elasticsearchMovingAverageHoltWintersModelSettingsSettings *ElasticsearchMovingAverageHoltWintersModelSettingsSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder](./builder-ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder.md)
