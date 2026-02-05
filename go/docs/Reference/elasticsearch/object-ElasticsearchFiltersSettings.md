---
title: <span class="badge object-type-struct"></span> ElasticsearchFiltersSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchFiltersSettings

## Definition

```go
type ElasticsearchFiltersSettings struct {
    Filters []elasticsearch.Filter `json:"filters,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchFiltersSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchFiltersSettings *ElasticsearchFiltersSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchFiltersSettings` objects.

```go
func (elasticsearchFiltersSettings *ElasticsearchFiltersSettings) Equals(other ElasticsearchFiltersSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchFiltersSettings` fields for violations and returns them.

```go
func (elasticsearchFiltersSettings *ElasticsearchFiltersSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchFiltersSettingsBuilder](./builder-ElasticsearchFiltersSettingsBuilder.md)
