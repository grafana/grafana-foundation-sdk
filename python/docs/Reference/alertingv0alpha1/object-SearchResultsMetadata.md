---
title: <span class="badge object-type-class"></span> SearchResultsMetadata
---
# <span class="badge object-type-class"></span> SearchResultsMetadata

#SearchResultsMetadata carries the paging token and total authorised match

count.

## Definition

```python
class SearchResultsMetadata:
    """
    #SearchResultsMetadata carries the paging token and total authorised match
    count.
    """

    continue_val: typing.Optional[str]
    total_hits: typing.Optional[int]
    # Always read totalHits together with totalHitsRelation.
    total_hits_relation: typing.Optional[alertingv0alpha1.TotalHitsRelation]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [SearchResultsMetadata](./builder-SearchResultsMetadata.md)
