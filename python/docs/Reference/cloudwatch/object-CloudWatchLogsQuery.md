---
title: <span class="badge object-type-class"></span> CloudWatchLogsQuery
---
# <span class="badge object-type-class"></span> CloudWatchLogsQuery

Shape of a CloudWatch Logs query

## Definition

```python
class CloudWatchLogsQuery(cogvariants.Dataquery):
    """
    Shape of a CloudWatch Logs query
    """

    # Whether a query is a Metrics, Logs, or Annotations query
    query_mode: cloudwatch.CloudWatchQueryMode
    id_val: str
    # AWS region to query for the logs
    region: str
    # The CloudWatch Logs Insights query to execute
    expression: typing.Optional[str]
    # Fields to group the results by, this field is automatically populated whenever the query is updated
    stats_groups: typing.Optional[list[str]]
    # Log groups to query
    log_groups: typing.Optional[list[cloudwatch.LogGroup]]
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
    # @deprecated use logGroups
    log_group_names: typing.Optional[list[str]]
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

 * <span class="badge builder"></span> [CloudWatchLogsQuery](./builder-CloudWatchLogsQuery.md)
