---
title: <span class="badge object-type-class"></span> TimeSeriesQuery
---
# <span class="badge object-type-class"></span> TimeSeriesQuery

Time Series sub-query properties.

## Definition

```python
class TimeSeriesQuery:
    """
    Time Series sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # MQL query to be executed.
    query: str
    # To disable the graphPeriod, it should explictly be set to 'disabled'.
    graph_period: typing.Optional[str]
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

 * <span class="badge builder"></span> [TimeSeriesQuery](./builder-TimeSeriesQuery.md)
