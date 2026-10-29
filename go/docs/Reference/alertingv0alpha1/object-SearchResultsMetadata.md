---
title: <span class="badge object-type-struct"></span> SearchResultsMetadata
---
# <span class="badge object-type-struct"></span> SearchResultsMetadata

#SearchResultsMetadata carries the paging token and total authorised match

count.

## Definition

```go
type SearchResultsMetadata struct {
    Continue *string `json:"continue,omitempty"`
    TotalHits *int64 `json:"totalHits,omitempty"`
    // Always read totalHits together with totalHitsRelation.
    TotalHitsRelation *alertingv0alpha1.TotalHitsRelation `json:"totalHitsRelation,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SearchResultsMetadata` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (searchResultsMetadata *SearchResultsMetadata) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SearchResultsMetadata` objects.

```go
func (searchResultsMetadata *SearchResultsMetadata) Equals(other SearchResultsMetadata) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SearchResultsMetadata` fields for violations and returns them.

```go
func (searchResultsMetadata *SearchResultsMetadata) Validate() error
```

## See also

 * <span class="badge builder"></span> [SearchResultsMetadataBuilder](./builder-SearchResultsMetadataBuilder.md)
