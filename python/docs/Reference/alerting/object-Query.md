---
title: <span class="badge object-type-class"></span> Query
---
# <span class="badge object-type-class"></span> Query

## Definition

```python
class Query:
    # Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
    datasource_uid: typing.Optional[str]
    # JSON is the raw JSON query and includes the above properties as well as custom properties.
    model: typing.Optional[cogvariants.Dataquery]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: typing.Optional[str]
    # RelativeTimeRange is the per query start and end time
    # for requests.
    relative_time_range: typing.Optional[alerting.RelativeTimeRange]
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

 * <span class="badge builder"></span> [Query](./builder-Query.md)
