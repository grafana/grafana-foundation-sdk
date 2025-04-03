---
title: <span class="badge builder"></span> ElasticsearchHistogramSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchHistogramSettingsBuilder

## Constructor

```go
func NewElasticsearchHistogramSettingsBuilder() *ElasticsearchHistogramSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchHistogramSettingsBuilder) Build() (ElasticsearchHistogramSettings, error)
```

### <span class="badge object-method"></span> Interval

```go
func (builder *ElasticsearchHistogramSettingsBuilder) Interval(interval string) *ElasticsearchHistogramSettingsBuilder
```

### <span class="badge object-method"></span> MinDocCount

```go
func (builder *ElasticsearchHistogramSettingsBuilder) MinDocCount(minDocCount string) *ElasticsearchHistogramSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchHistogramSettings](./object-ElasticsearchHistogramSettings.md)
