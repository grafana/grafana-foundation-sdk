---
title: <span class="badge object-type-class"></span> AzureLogsQuery
---
# <span class="badge object-type-class"></span> AzureLogsQuery

Azure Monitor Logs sub-query properties

## Definition

```python
class AzureLogsQuery:
    """
    Azure Monitor Logs sub-query properties
    """

    # KQL query to be executed.
    query: typing.Optional[str]
    # Specifies the format results should be returned as.
    result_format: typing.Optional[azuremonitor.ResultFormat]
    # Array of resource URIs to be queried.
    resources: typing.Optional[list[str]]
    # If set to true the intersection of time ranges specified in the query and Grafana will be used. Otherwise the query time ranges will be used. Defaults to false
    intersect_time: typing.Optional[bool]
    # Workspace ID. This was removed in Grafana 8, but remains for backwards compat
    workspace: typing.Optional[str]
    # @deprecated Use resources instead
    resource: typing.Optional[str]
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

 * <span class="badge builder"></span> [AzureLogsQuery](./builder-AzureLogsQuery.md)
