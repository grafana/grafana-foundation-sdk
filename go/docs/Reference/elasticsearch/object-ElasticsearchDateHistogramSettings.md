---
title: <span class="badge object-type-struct"></span> ElasticsearchDateHistogramSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchDateHistogramSettings

## Definition

```go
type ElasticsearchDateHistogramSettings struct {
    Interval *string `json:"interval,omitempty"`
    MinDocCount *string `json:"min_doc_count,omitempty"`
    TrimEdges *string `json:"trimEdges,omitempty"`
    Offset *string `json:"offset,omitempty"`
    TimeZone *string `json:"timeZone,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchDateHistogramSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchDateHistogramSettings *ElasticsearchDateHistogramSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchDateHistogramSettings` objects.

```go
func (elasticsearchDateHistogramSettings *ElasticsearchDateHistogramSettings) Equals(other ElasticsearchDateHistogramSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchDateHistogramSettings` fields for violations and returns them.

```go
func (elasticsearchDateHistogramSettings *ElasticsearchDateHistogramSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchDateHistogramSettingsBuilder](./builder-ElasticsearchDateHistogramSettingsBuilder.md)
