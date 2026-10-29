---
title: <span class="badge object-type-class"></span> SearchResultHit
---
# <span class="badge object-type-class"></span> SearchResultHit

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Definition

```python
class SearchResultHit:
    """
    #SearchResultHit is a single match: its identity, an optional relevance
    score (present only when the query included free text), and the requested
    fields.
    """

    resource: alertingv0alpha1.SearchResultResource
    score: typing.Optional[float]
    fields: alertingv0alpha1.RuleSearchHitFields
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

 * <span class="badge builder"></span> [SearchResultHit](./builder-SearchResultHit.md)
