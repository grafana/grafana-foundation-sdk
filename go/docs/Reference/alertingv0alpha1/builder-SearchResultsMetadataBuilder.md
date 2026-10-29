---
title: <span class="badge builder"></span> SearchResultsMetadataBuilder
---
# <span class="badge builder"></span> SearchResultsMetadataBuilder

#SearchResultsMetadata carries the paging token and total authorised match

count.

## Constructor

```go
func NewSearchResultsMetadataBuilder() *SearchResultsMetadataBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SearchResultsMetadataBuilder) Build() (SearchResultsMetadata, error)
```

### <span class="badge object-method"></span> Continue

```go
func (builder *SearchResultsMetadataBuilder) Continue(continueArg string) *SearchResultsMetadataBuilder
```

### <span class="badge object-method"></span> TotalHits

```go
func (builder *SearchResultsMetadataBuilder) TotalHits(totalHits int64) *SearchResultsMetadataBuilder
```

### <span class="badge object-method"></span> TotalHitsRelation

Always read totalHits together with totalHitsRelation.

```go
func (builder *SearchResultsMetadataBuilder) TotalHitsRelation(totalHitsRelation alertingv0alpha1.TotalHitsRelation) *SearchResultsMetadataBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SearchResultsMetadata](./object-SearchResultsMetadata.md)
