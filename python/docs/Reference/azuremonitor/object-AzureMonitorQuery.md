---
title: <span class="badge object-type-class"></span> AzureMonitorQuery
---
# <span class="badge object-type-class"></span> AzureMonitorQuery

## Definition

```python
class AzureMonitorQuery(cogvariants.Dataquery):
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # Azure subscription containing the resource(s) to be queried.
    subscription: typing.Optional[str]
    # Subscriptions to be queried via Azure Resource Graph.
    subscriptions: typing.Optional[list[str]]
    # Azure Monitor Metrics sub-query properties.
    azure_monitor: typing.Optional[azuremonitor.AzureMetricQuery]
    # Azure Monitor Logs sub-query properties.
    azure_log_analytics: typing.Optional[azuremonitor.AzureLogsQuery]
    # Azure Resource Graph sub-query properties.
    azure_resource_graph: typing.Optional[azuremonitor.AzureResourceGraphQuery]
    # Application Insights Traces sub-query properties.
    azure_traces: typing.Optional[azuremonitor.AzureTracesQuery]
    # @deprecated Legacy template variable support.
    grafana_template_variable_fn: typing.Optional[azuremonitor.GrafanaTemplateVariableQuery]
    # Template variables params. These exist for backwards compatiblity with legacy template variables.
    resource_group: typing.Optional[str]
    namespace: typing.Optional[str]
    resource: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]
    # Azure Monitor query type.
    # queryType: #AzureQueryType
    region: typing.Optional[str]
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

 * <span class="badge builder"></span> [AzureMonitorQuery](./builder-AzureMonitorQuery.md)
