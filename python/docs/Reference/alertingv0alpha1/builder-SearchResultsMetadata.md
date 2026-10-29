---
title: <span class="badge builder"></span> SearchResultsMetadata
---
# <span class="badge builder"></span> SearchResultsMetadata

#SearchResultsMetadata carries the paging token and total authorised match

count.

## Constructor

```python
SearchResultsMetadata()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> alertingv0alpha1.SearchResultsMetadata
```

### <span class="badge object-method"></span> continue_val

```python
def continue_val(continue_val: str) -> typing.Self
```

### <span class="badge object-method"></span> total_hits

```python
def total_hits(total_hits: int) -> typing.Self
```

### <span class="badge object-method"></span> total_hits_relation

Always read totalHits together with totalHitsRelation.

```python
def total_hits_relation(total_hits_relation: alertingv0alpha1.TotalHitsRelation) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [SearchResultsMetadata](./object-SearchResultsMetadata.md)
