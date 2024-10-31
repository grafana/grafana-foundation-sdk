---
title: <span class="badge builder"></span> ElasticsearchDateHistogramSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchDateHistogramSettingsBuilder

## Constructor

```go
func NewElasticsearchDateHistogramSettingsBuilder() *ElasticsearchDateHistogramSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchDateHistogramSettingsBuilder) Build() (ElasticsearchDateHistogramSettings, error)
```

### <span class="badge object-method"></span> Interval

```go
func (builder *ElasticsearchDateHistogramSettingsBuilder) Interval(interval string) *ElasticsearchDateHistogramSettingsBuilder
```

### <span class="badge object-method"></span> MinDocCount

```go
func (builder *ElasticsearchDateHistogramSettingsBuilder) MinDocCount(minDocCount string) *ElasticsearchDateHistogramSettingsBuilder
```

### <span class="badge object-method"></span> Offset

```go
func (builder *ElasticsearchDateHistogramSettingsBuilder) Offset(offset string) *ElasticsearchDateHistogramSettingsBuilder
```

### <span class="badge object-method"></span> TimeZone

```go
func (builder *ElasticsearchDateHistogramSettingsBuilder) TimeZone(timeZone string) *ElasticsearchDateHistogramSettingsBuilder
```

### <span class="badge object-method"></span> TrimEdges

```go
func (builder *ElasticsearchDateHistogramSettingsBuilder) TrimEdges(trimEdges string) *ElasticsearchDateHistogramSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchDateHistogramSettings](./object-ElasticsearchDateHistogramSettings.md)
