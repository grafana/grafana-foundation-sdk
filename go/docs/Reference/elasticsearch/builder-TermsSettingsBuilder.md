---
title: <span class="badge builder"></span> TermsSettingsBuilder
---
# <span class="badge builder"></span> TermsSettingsBuilder

## Constructor

```go
func NewTermsSettingsBuilder() *TermsSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TermsSettingsBuilder) Build() (TermsSettings, error)
```

### <span class="badge object-method"></span> MinDocCount

```go
func (builder *TermsSettingsBuilder) MinDocCount(minDocCount string) *TermsSettingsBuilder
```

### <span class="badge object-method"></span> Missing

```go
func (builder *TermsSettingsBuilder) Missing(missing string) *TermsSettingsBuilder
```

### <span class="badge object-method"></span> Order

```go
func (builder *TermsSettingsBuilder) Order(order elasticsearch.TermsOrder) *TermsSettingsBuilder
```

### <span class="badge object-method"></span> OrderBy

```go
func (builder *TermsSettingsBuilder) OrderBy(orderBy string) *TermsSettingsBuilder
```

### <span class="badge object-method"></span> Size

```go
func (builder *TermsSettingsBuilder) Size(size string) *TermsSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TermsSettings](./object-TermsSettings.md)
