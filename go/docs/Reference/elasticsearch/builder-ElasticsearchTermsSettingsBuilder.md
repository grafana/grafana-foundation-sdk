---
title: <span class="badge builder"></span> ElasticsearchTermsSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchTermsSettingsBuilder

## Constructor

```go
func NewElasticsearchTermsSettingsBuilder() *ElasticsearchTermsSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchTermsSettingsBuilder) Build() (ElasticsearchTermsSettings, error)
```

### <span class="badge object-method"></span> MinDocCount

```go
func (builder *ElasticsearchTermsSettingsBuilder) MinDocCount(minDocCount string) *ElasticsearchTermsSettingsBuilder
```

### <span class="badge object-method"></span> Missing

```go
func (builder *ElasticsearchTermsSettingsBuilder) Missing(missing string) *ElasticsearchTermsSettingsBuilder
```

### <span class="badge object-method"></span> Order

```go
func (builder *ElasticsearchTermsSettingsBuilder) Order(order elasticsearch.TermsOrder) *ElasticsearchTermsSettingsBuilder
```

### <span class="badge object-method"></span> OrderBy

```go
func (builder *ElasticsearchTermsSettingsBuilder) OrderBy(orderBy string) *ElasticsearchTermsSettingsBuilder
```

### <span class="badge object-method"></span> Size

```go
func (builder *ElasticsearchTermsSettingsBuilder) Size(size string) *ElasticsearchTermsSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchTermsSettings](./object-ElasticsearchTermsSettings.md)
