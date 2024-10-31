---
title: <span class="badge object-type-class"></span> TempoQuery
---
# <span class="badge object-type-class"></span> TempoQuery

## Definition

```python
class TempoQuery(cogvariants.Dataquery):
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # Note this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # TraceQL query or trace ID
    query: typing.Optional[str]
    # @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
    search: typing.Optional[str]
    # @deprecated Query traces by service name
    service_name: typing.Optional[str]
    # @deprecated Query traces by span name
    span_name: typing.Optional[str]
    # @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
    min_duration: typing.Optional[str]
    # @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
    max_duration: typing.Optional[str]
    # Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}. Providing multiple values will produce union of results for each filter, using PromQL OR operator internally.
    service_map_query: typing.Optional[typing.Union[str, list[str]]]
    # Use service.namespace in addition to service.name to uniquely identify a service.
    service_map_include_namespace: typing.Optional[bool]
    # Defines the maximum number of traces that are returned from Tempo
    limit: typing.Optional[int]
    # Defines the maximum number of spans per spanset that are returned from Tempo
    spss: typing.Optional[int]
    filters: list[tempo.TraceqlFilter]
    # Filters that are used to query the metrics summary
    group_by: typing.Optional[list[tempo.TraceqlFilter]]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]
    # The type of the table that is used to display the search results
    table_type: typing.Optional[tempo.SearchTableType]
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

 * <span class="badge builder"></span> [TempoQuery](./builder-TempoQuery.md)
