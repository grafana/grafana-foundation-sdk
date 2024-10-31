---
title: <span class="badge object-type-class"></span> PromQLQuery
---
# <span class="badge object-type-class"></span> PromQLQuery

PromQL sub-query properties.

## Definition

```python
class PromQLQuery:
    """
    PromQL sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # PromQL expression/query to be executed.
    expr: str
    # PromQL min step
    step: str
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

 * <span class="badge builder"></span> [PromQLQuery](./builder-PromQLQuery.md)
