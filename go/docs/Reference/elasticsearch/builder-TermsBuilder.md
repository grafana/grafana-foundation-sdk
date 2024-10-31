---
title: <span class="badge builder"></span> TermsBuilder
---
# <span class="badge builder"></span> TermsBuilder

## Constructor

```go
func NewTermsBuilder() *TermsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TermsBuilder) Build() (Terms, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *TermsBuilder) Field(field string) *TermsBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *TermsBuilder) Id(id string) *TermsBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *TermsBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchTermsSettings]) *TermsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Terms](./object-Terms.md)
