---
title: <span class="badge builder"></span> AzureMonitorQuery
---
# <span class="badge builder"></span> AzureMonitorQuery

## Constructor

```python
AzureMonitorQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> azuremonitor.AzureMonitorQuery
```

### <span class="badge object-method"></span> azure_log_analytics

Azure Monitor Logs sub-query properties.

```python
def azure_log_analytics(azure_log_analytics: cogbuilder.Builder[azuremonitor.AzureLogsQuery]) -> typing.Self
```

### <span class="badge object-method"></span> azure_monitor

Azure Monitor Metrics sub-query properties.

```python
def azure_monitor(azure_monitor: cogbuilder.Builder[azuremonitor.AzureMetricQuery]) -> typing.Self
```

### <span class="badge object-method"></span> azure_resource_graph

Azure Resource Graph sub-query properties.

```python
def azure_resource_graph(azure_resource_graph: cogbuilder.Builder[azuremonitor.AzureResourceGraphQuery]) -> typing.Self
```

### <span class="badge object-method"></span> azure_traces

Application Insights Traces sub-query properties.

```python
def azure_traces(azure_traces: cogbuilder.Builder[azuremonitor.AzureTracesQuery]) -> typing.Self
```

### <span class="badge object-method"></span> datasource

For mixed data sources the selected datasource is on the query level.

For non mixed scenarios this is undefined.

TODO find a better way to do this ^ that's friendly to schema

TODO this shouldn't be unknown but DataSourceRef | null

```python
def datasource(datasource: dashboard.DataSourceRef) -> typing.Self
```

### <span class="badge object-method"></span> grafana_template_variable_fn

@deprecated Legacy template variable support.

```python
def grafana_template_variable_fn(grafana_template_variable_fn: azuremonitor.GrafanaTemplateVariableQuery) -> typing.Self
```

### <span class="badge object-method"></span> hide

true if query is disabled (ie should not be returned to the dashboard)

Note this does not always imply that the query should not be executed since

the results from a hidden query may be used as the input to other queries (SSE etc)

```python
def hide(hide: bool) -> typing.Self
```

### <span class="badge object-method"></span> namespace

```python
def namespace(namespace: str) -> typing.Self
```

### <span class="badge object-method"></span> query_type

Specify the query flavor

TODO make this required and give it a default

```python
def query_type(query_type: str) -> typing.Self
```

### <span class="badge object-method"></span> ref_id

A unique identifier for the query within the list of targets.

In server side expressions, the refId is used as a variable name to identify results.

By default, the UI will assign A->Z; however setting meaningful names may be useful.

```python
def ref_id(ref_id: str) -> typing.Self
```

### <span class="badge object-method"></span> region

Azure Monitor query type.

queryType: #AzureQueryType

```python
def region(region: str) -> typing.Self
```

### <span class="badge object-method"></span> resource

```python
def resource(resource: str) -> typing.Self
```

### <span class="badge object-method"></span> resource_group

Template variables params. These exist for backwards compatiblity with legacy template variables.

```python
def resource_group(resource_group: str) -> typing.Self
```

### <span class="badge object-method"></span> subscription

Azure subscription containing the resource(s) to be queried.

```python
def subscription(subscription: str) -> typing.Self
```

### <span class="badge object-method"></span> subscriptions

Subscriptions to be queried via Azure Resource Graph.

```python
def subscriptions(subscriptions: list[str]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AzureMonitorQuery](./object-AzureMonitorQuery.md)
