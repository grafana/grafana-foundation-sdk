---
title: <span class="badge object-type-class"></span> SearchResults
---
# <span class="badge object-type-class"></span> SearchResults

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Definition

```python
class SearchResults:
    """
    #SearchResults is the response envelope, mirroring
    search.grafana.app SearchResults.
    """

    metadata: alertingv0alpha1.SearchResultsMetadata
    items: list[alertingv0alpha1.SearchResultHit]
    facets: typing.Optional[dict[str, list[alertingv0alpha1.FacetValue]]]
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

 * <span class="badge builder"></span> [SearchResults](./builder-SearchResults.md)
