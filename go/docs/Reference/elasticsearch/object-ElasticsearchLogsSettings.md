---
title: <span class="badge object-type-struct"></span> ElasticsearchLogsSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchLogsSettings

## Definition

```go
type ElasticsearchLogsSettings struct {
    Limit *string `json:"limit,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchLogsSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchLogsSettings *ElasticsearchLogsSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchLogsSettings` objects.

```go
func (elasticsearchLogsSettings *ElasticsearchLogsSettings) Equals(other ElasticsearchLogsSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchLogsSettings` fields for violations and returns them.

```go
func (elasticsearchLogsSettings *ElasticsearchLogsSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchLogsSettingsBuilder](./builder-ElasticsearchLogsSettingsBuilder.md)
