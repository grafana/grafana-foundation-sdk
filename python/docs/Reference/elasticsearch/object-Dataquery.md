---
title: <span class="badge object-type-class"></span> Dataquery
---
# <span class="badge object-type-class"></span> Dataquery

## Definition

```python
class Dataquery(cogvariants.Dataquery):
    # Alias pattern
    alias: typing.Optional[str]
    # Lucene query
    query: typing.Optional[str]
    # Name of time field
    time_field: typing.Optional[str]
    # List of bucket aggregations
    bucket_aggs: typing.Optional[list[elasticsearch.BucketAggregation]]
    # List of metric aggregations
    metrics: typing.Optional[list[elasticsearch.MetricAggregation]]
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]
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

 * <span class="badge builder"></span> [Dataquery](./builder-Dataquery.md)
