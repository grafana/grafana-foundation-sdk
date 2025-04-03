---
title: <span class="badge object-type-class"></span> AzureTracesQuery
---
# <span class="badge object-type-class"></span> AzureTracesQuery

Application Insights Traces sub-query properties

## Definition

```python
class AzureTracesQuery:
    """
    Application Insights Traces sub-query properties
    """

    # Specifies the format results should be returned as.
    result_format: typing.Optional[azuremonitor.ResultFormat]
    # Array of resource URIs to be queried.
    resources: typing.Optional[list[str]]
    # Operation ID. Used only for Traces queries.
    operation_id: typing.Optional[str]
    # Types of events to filter by.
    trace_types: typing.Optional[list[str]]
    # Filters for property values.
    filters: typing.Optional[list[azuremonitor.AzureTracesFilter]]
    # KQL query to be executed.
    query: typing.Optional[str]
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

 * <span class="badge builder"></span> [AzureTracesQuery](./builder-AzureTracesQuery.md)
