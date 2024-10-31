---
title: <span class="badge object-type-struct"></span> ElasticsearchHistogramSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchHistogramSettings

## Definition

```go
type ElasticsearchHistogramSettings struct {
    Interval *string `json:"interval,omitempty"`
    MinDocCount *string `json:"min_doc_count,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchHistogramSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchHistogramSettings *ElasticsearchHistogramSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchHistogramSettings` objects.

```go
func (elasticsearchHistogramSettings *ElasticsearchHistogramSettings) Equals(other ElasticsearchHistogramSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchHistogramSettings` fields for violations and returns them.

```go
func (elasticsearchHistogramSettings *ElasticsearchHistogramSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchHistogramSettingsBuilder](./builder-ElasticsearchHistogramSettingsBuilder.md)
