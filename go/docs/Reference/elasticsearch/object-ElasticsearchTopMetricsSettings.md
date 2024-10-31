---
title: <span class="badge object-type-struct"></span> ElasticsearchTopMetricsSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchTopMetricsSettings

## Definition

```go
type ElasticsearchTopMetricsSettings struct {
    Order *string `json:"order,omitempty"`
    OrderBy *string `json:"orderBy,omitempty"`
    Metrics []string `json:"metrics,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchTopMetricsSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchTopMetricsSettings *ElasticsearchTopMetricsSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchTopMetricsSettings` objects.

```go
func (elasticsearchTopMetricsSettings *ElasticsearchTopMetricsSettings) Equals(other ElasticsearchTopMetricsSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchTopMetricsSettings` fields for violations and returns them.

```go
func (elasticsearchTopMetricsSettings *ElasticsearchTopMetricsSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchTopMetricsSettingsBuilder](./builder-ElasticsearchTopMetricsSettingsBuilder.md)
