---
title: <span class="badge object-type-class"></span> MetricDefinitionsQuery
---
# <span class="badge object-type-class"></span> MetricDefinitionsQuery

@deprecated Use MetricNamespaceQuery instead

## Definition

```python
class MetricDefinitionsQuery:
    """
    @deprecated Use MetricNamespaceQuery instead
    """

    raw_query: typing.Optional[str]
    kind: typing.Literal["MetricDefinitionsQuery"]
    subscription: str
    resource_group: str
    metric_namespace: typing.Optional[str]
    resource_name: typing.Optional[str]
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

 * <span class="badge builder"></span> [MetricDefinitionsQuery](./builder-MetricDefinitionsQuery.md)
