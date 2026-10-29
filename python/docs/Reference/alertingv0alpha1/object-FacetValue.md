---
title: <span class="badge object-type-class"></span> FacetValue
---
# <span class="badge object-type-class"></span> FacetValue

#FacetValue is a single value/count pair in a facet breakdown.

## Definition

```python
class FacetValue:
    """
    #FacetValue is a single value/count pair in a facet breakdown.
    """

    value: str
    count: int
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

 * <span class="badge builder"></span> [FacetValue](./builder-FacetValue.md)
